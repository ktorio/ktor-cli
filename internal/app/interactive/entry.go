package interactive

import (
	"context"
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/ktorio/ktor-cli/internal/app/i18n"
	"github.com/ktorio/ktor-cli/internal/app/interactive/draw"
	"github.com/ktorio/ktor-cli/internal/app/interactive/model"
	"github.com/ktorio/ktor-cli/internal/app/network"
	"net/http"
	"slices"
	"strings"
)

func Run(client *http.Client, ctx context.Context) (result model.Result, err error) {
	settings, err := network.FetchSettings(client)

	if err != nil {
		return
	}

	result = model.Result{}
	drawState := draw.NewState()
	mdl := model.NewState()

	mdl.Plugins = []string{}
	mdl.ProjectName = settings.ProjectName.DefaultVal
	model.InitProjectDir(mdl)

	if !model.IsDirEmptyOrAbsent(mdl.GetProjectPath()) {
		n, err := model.FindVacantProjectName(mdl)

		if err != nil {
			return result, err
		}

		mdl.ProjectName = n
	}

	scr, err := tcell.NewScreen()

	if err != nil {
		return
	}

	if err = scr.Init(); err != nil {
		return
	}

	scr.EnableMouse()
	scr.Clear()

	quit := func() {
		maybePanic := recover()
		scr.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}
	defer quit()

	scr.SetCursorStyle(tcell.CursorStyleBlinkingBar)
	model.CheckProjectSettings(mdl)

	for mdl.Running {
		event := scr.PollEvent()

		switch drawState.ActiveElement {
		case draw.ProjectNameInput:
			processEvent(event, drawState, mdl, &result, &mdl.ProjectName)
		case draw.LocationInput:
			processEvent(event, drawState, mdl, &result, &mdl.ProjectDir)
		case draw.SearchInput:
			processEvent(event, drawState, mdl, &result, &mdl.Search)
		case draw.Tabs:
			processEvent(event, drawState, mdl, &result, nil)
		case draw.CreateButton:
			processEvent(event, drawState, mdl, &result, nil)
		default:
			panic("unhandled default case")
		}

		draw.HideCursorIfNeeded(drawState, scr)

		if mdl.ShouldFetchPlugins && !mdl.PluginsFetched {
			var plugins []network.Plugin
			plugins, err = network.FetchPlugins(client, settings.KtorVersion.DefaultId, ctx)

			if err != nil {
				mdl.SetError(model.UnableFetchPluginsError, i18n.Get(i18n.UnableFetchPluginsError))
				mdl.PluginsFetched = true
				continue
			}

			mdl.PluginsFetched = true

			mdl.PluginDeps = make(map[string][]string, len(plugins))
			mdl.AllPluginsByGroup = make(map[string][]network.Plugin, len(plugins))
			for _, p := range plugins {
				mdl.PluginDeps[p.Id] = p.RequiredPlugins

				if !slices.Contains(mdl.AllSortedGroups, p.Group) {
					mdl.AllSortedGroups = append(mdl.AllSortedGroups, p.Group)
				}
				mdl.AllPluginsByGroup[p.Group] = append(mdl.AllPluginsByGroup[p.Group], p)
			}
			mdl.PluginsByGroup = mdl.AllPluginsByGroup

			slices.Sort(mdl.AllSortedGroups)
			mdl.Groups = mdl.AllSortedGroups
		}

		scr.Clear()
		scr.Fill(' ', draw.DefaultStyle)

		draw.Tui(scr, drawState, mdl)
		scr.Show()
	}

	return
}

var firstResize = false

func processEvent(ev tcell.Event, drawState *draw.State, mdl *model.State, result *model.Result, input *string) {
	inputOff := draw.ActiveInputOffset(drawState)

	switch ev := ev.(type) {
	case *tcell.EventResize:
		if !firstResize {
			firstResize = true
		}
	case *tcell.EventKey:
		mod, key := ev.Modifiers(), ev.Key()

		switch {
		case (mod == tcell.ModCtrl && key == tcell.KeyCtrlC) || (key == tcell.KeyEscape):
			mdl.Running = false
			result.Quit = true
		case mod == tcell.ModCtrl && key == tcell.KeyBEL:
			if drawState.PluginsShown {
				generateProject(result, mdl)
			}
		case key == tcell.KeyCtrlA:
			if input != nil {
				drawState.CursorOffs[drawState.ActiveElement] = 0
				drawState.VisibleOffs[drawState.ActiveElement] = 0
			}
		case key == tcell.KeyCtrlE:
			if input != nil {
				runes := []rune(*input)
				drawState.CursorOffs[drawState.ActiveElement] = len(runes)
				drawState.VisibleOffs[drawState.ActiveElement] = max(0, len(runes)-drawState.InputLen()+1)
			}
		case key == tcell.KeyRune:
			if !draw.IsElementActive(drawState, draw.ProjectNameInput) &&
				!draw.IsElementActive(drawState, draw.LocationInput) &&
				!draw.IsElementActive(drawState, draw.SearchInput) && ev.Rune() == '/' {

				drawState.ActiveElement = draw.SearchInput
				return
			}

			if draw.IsElementActive(drawState, draw.Tabs) && ev.Rune() == ' ' {
				toggleSelectedPlugin(drawState, mdl)
				mdl.StatusLine = fmt.Sprintf(i18n.Get(i18n.SelectedPluginsCount, len(mdl.AddedPlugins)))
				return
			}

			if input != nil {
				*input = model.InsertRune(*input, inputOff, ev.Rune())

				runes := []rune(*input)
				draw.MoveCursor(drawState, len(runes), 1)
				onInputChanged(drawState, mdl, *input)

				if drawState.CursorPos()-drawState.VisOff() >= drawState.InputLen() {
					drawState.VisibleOffs[drawState.ActiveElement]++
				}
			}
		case key == tcell.KeyLeft:
			if draw.IsElementActive(drawState, draw.Tabs) {
				draw.SwitchTab(drawState, len(mdl.Groups), -1)
				return
			}

			if input != nil {
				runes := []rune(*input)
				draw.MoveCursor(drawState, len(runes), -1)

				if drawState.CursorPos() < drawState.VisOff() {
					drawState.VisibleOffs[drawState.ActiveElement]--
				}
			}
		case key == tcell.KeyRight:
			if draw.IsElementActive(drawState, draw.Tabs) {
				draw.SwitchTab(drawState, len(mdl.Groups), 1)
				return
			}

			if input != nil {
				runes := []rune(*input)
				draw.MoveCursor(drawState, len(runes), 1)

				if drawState.CursorPos()-drawState.VisOff() >= drawState.InputLen() {
					drawState.VisibleOffs[drawState.ActiveElement]++
				}
			}
		case key == tcell.KeyUp:
			if draw.IsElementActive(drawState, draw.Tabs) {
				if drawState.ActivePlugin == 0 {
					draw.SwitchElement(drawState, -1)
					return
				}

				if drawState.ActivePlugin-1 >= 0 {
					drawState.ActivePlugin--
				}
				return
			}

			draw.SwitchElement(drawState, -1)
		case key == tcell.KeyDown:
			if drawState.ActiveElement == draw.Tabs {
				if drawState.ActivePlugin+1 > len(pluginsOnCurrentTab(drawState, mdl))-1 {
					draw.SwitchElement(drawState, 1)
					return
				}

				if drawState.ActivePlugin+1 < len(pluginsOnCurrentTab(drawState, mdl)) {
					drawState.ActivePlugin++
				}
				return
			}

			switch {
			case drawState.ActiveElement == draw.ProjectNameInput:
				draw.SwitchElement(drawState, 1)
			case drawState.ActiveElement == draw.LocationInput && drawState.PluginsShown:
				draw.SwitchElement(drawState, 1)
			case drawState.PluginsShown:
				draw.SwitchElement(drawState, 1)
			default:
				// do nothing yet
			}
		case key == tcell.KeyDEL: // Backspace
			if input != nil {
				*input = model.DeleteChar(*input, inputOff-1)
				runes := []rune(*input)
				draw.MoveCursor(drawState, len(runes), -1)
				if drawState.VisOff() > 0 {
					drawState.VisibleOffs[drawState.ActiveElement]--
				}
				onInputChanged(drawState, mdl, *input)
			}
		case key == tcell.KeyDelete:
			if input != nil {
				*input = model.DeleteChar(*input, inputOff)
				if drawState.VisOff() > 0 {
					drawState.VisibleOffs[drawState.ActiveElement]--
				}
				onInputChanged(drawState, mdl, *input)
			}
		case key == tcell.KeyEnter && mod == tcell.ModNone:
			if drawState.ActiveElement == draw.Tabs {
				toggleSelectedPlugin(drawState, mdl)
				mdl.StatusLine = fmt.Sprintf(i18n.Get(i18n.SelectedPluginsCount, len(mdl.AddedPlugins)))
				return
			} else if drawState.ActiveElement == draw.CreateButton {
				if generateProject(result, mdl) {
					return
				}
			}

			switch drawState.ActiveElement {
			case draw.ProjectNameInput:
				model.InitProjectDir(mdl)
				model.CheckProjectSettings(mdl)
			case draw.LocationInput:
				drawState.PluginsShown = true
			default:
				// do nothing yet
			}

			draw.SwitchElement(drawState, 1)
		case key == tcell.KeyTab:
			switch drawState.ActiveElement {
			case draw.ProjectNameInput:
				model.InitProjectDir(mdl)
				model.CheckProjectSettings(mdl)
			case draw.LocationInput:
				drawState.PluginsShown = true
			default:
				// do nothing yet
			}

			draw.SwitchElement(drawState, 1)
		case key == tcell.KeyBacktab: // Shift + Tab
			draw.SwitchElement(drawState, -1)
		}
	}
}

func generateProject(result *model.Result, mdl *model.State) bool {
	result.ProjectName = mdl.ProjectName
	result.ProjectDir = mdl.GetProjectPath()

	result.Plugins = []string{}
	for id := range mdl.AddedPlugins {
		result.Plugins = append(result.Plugins, id)
	}

	if hasError := model.CheckProjectSettings(mdl); !hasError {
		mdl.Running = false
		return true
	}

	return false
}

func searchPlugins(mdl *model.State, drawState *draw.State) map[string][]network.Plugin {
	m := make(map[string][]network.Plugin)
	mdl.Groups = mdl.Groups[:0]
	for gr, ps := range mdl.AllPluginsByGroup {
		var plugins []network.Plugin

		for _, p := range ps {
			if strings.Contains(strings.ToLower(p.Name), strings.ToLower(mdl.Search)) || strings.Contains(strings.ToLower(p.Description), strings.ToLower(mdl.Search)) {
				plugins = append(plugins, p)
			}
		}

		if len(plugins) > 0 {
			slices.SortFunc(plugins, func(a, b network.Plugin) int {
				return strings.Compare(a.Name, b.Name)
			})
			m[gr] = plugins
			mdl.Groups = append(mdl.Groups, gr)
		}
	}

	if drawState.ActiveTab >= len(mdl.Groups) {
		drawState.ActiveTab = len(mdl.Groups) - 1
	}

	if drawState.ActiveTab < 0 {
		drawState.ActiveTab = 0
	}

	slices.Sort(mdl.Groups)

	return m
}

func onInputChanged(drawState *draw.State, mdl *model.State, input string) {
	if drawState.ActiveElement == draw.SearchInput {
		mdl.Search = input
		mdl.PluginsByGroup = searchPlugins(mdl, drawState)
		return
	}

	model.CheckProjectSettings(mdl)
}

func toggleSelectedPlugin(drawState *draw.State, mdl *model.State) {
	p := pluginsOnCurrentTab(drawState, mdl)[drawState.ActivePlugin]

	if _, ok := mdl.AddedPlugins[p.Id]; ok { // Plugin is selected
		for id := range mdl.AddedPlugins {
			if isPluginRequiredFor(mdl, p.Id, id) {
				delete(mdl.AddedPlugins, id)
			}
		}

		delete(mdl.AddedPlugins, p.Id)

		if ips, ok := mdl.IndirectPlugins[p.Id]; ok {
			for id := range ips {
				delete(mdl.AddedPlugins, id)
			}
		}
		delete(mdl.IndirectPlugins, p.Id)
		return
	}

	mdl.AddedPlugins[p.Id] = struct{}{}
	deps := make(model.IdSet)
	getDepPlugins(mdl, p.Id, deps)

	for _, s := range mdl.IndirectPlugins {
		if _, ok := s[p.Id]; ok {
			delete(s, p.Id)
		}
	}

	for id := range deps {
		if _, ok := mdl.AddedPlugins[id]; ok {
			delete(deps, id)
		}
	}

	for id := range deps {
		mdl.AddedPlugins[id] = struct{}{}
	}

	mdl.IndirectPlugins[p.Id] = deps
}

func isPluginRequiredFor(mdl *model.State, parentId string, childId string) bool {
	for _, id := range mdl.PluginDeps[childId] {
		if id == parentId || isPluginRequiredFor(mdl, id, childId) {
			return true
		}
	}
	return false
}

func getDepPlugins(mdl *model.State, id string, m model.IdSet) {
	for _, dp := range mdl.PluginDeps[id] {
		m[dp] = struct{}{}
		getDepPlugins(mdl, dp, m)
	}
}

func pluginsOnCurrentTab(drawState *draw.State, mdl *model.State) []network.Plugin {
	if drawState.ActiveTab >= len(mdl.Groups) {
		return []network.Plugin{}
	}

	return mdl.PluginsByGroup[mdl.Groups[drawState.ActiveTab]]
}
