package interactive

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/ktorio/ktor-cli/internal/app/network"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"
	"unicode"
)

var searchHighlightColor = tcell.Color172
var mainColor = tcell.Color126
var bgColor = tcell.Color233
var textColor = tcell.ColorWhite
var inputColor = tcell.Color235
var strongTextColor = tcell.Color141
var weakTextColor = tcell.Color139
var errorColor = tcell.Color160
var statusColor = tcell.Color106

var defaultStyle = tcell.StyleDefault.Background(bgColor)
var inputStyle = defaultStyle.Background(inputColor).Foreground(textColor)
var cursorStyle = defaultStyle.Background(mainColor)
var buttonStyle = defaultStyle.Background(mainColor).Foreground(textColor)
var activeTabStyle = defaultStyle.Foreground(mainColor).Background(textColor)
var textStyle = defaultStyle.Foreground(textColor)
var weakTextStyle = defaultStyle.Foreground(weakTextColor)

type Element int

const (
	ProjectNameInput Element = iota
	LocationInput
	SearchInput
	Tabs
	CreateButton
	Last
)

type Range struct {
	start, end int
}

type idSet map[string]struct{}

var running bool
var cursorOffs map[Element]int
var locationShown bool
var pluginsShown bool
var activeElement Element
var cursorAnimTimer float64 = 0
var pluginsFetched bool
var allPluginsByGroup map[string][]network.Plugin
var allSortedGroups []string
var pluginsByGroup map[string][]network.Plugin
var groups []string
var activeTab int
var activePlugin int
var genResult Result
var search string
var pluginDeps map[string][]string
var indirectPlugins map[string]idSet
var addedPlugins idSet
var statusLine string
var errorLine string
var tabVisRanges []Range
var pluginVisRanges []Range

func init() {
	running = true
	cursorOffs = map[Element]int{
		ProjectNameInput: 0,
		LocationInput:    0,
		SearchInput:      0,
	}
	locationShown = false
	pluginsShown = false
	activeElement = ProjectNameInput
	pluginsFetched = false
	activeTab = 0
	activePlugin = 0
	genResult = Result{}
	search = ""
	indirectPlugins = make(map[string]idSet)
	addedPlugins = make(idSet)
	statusLine = ""
	errorLine = ""
}

type Result struct {
	ProjectName string
	ProjectDir  string
	Plugins     []string
	Quit        bool
}

// TODO: Improve color scheme
// TODO: Separate draw state from model state ???
// TODO: Add input key combinations: CTRL+A, CTRL+E, ALT+BACKSPACE, CTRL+Right, CTRL+Left

func Run(client *http.Client) (result Result, err error) {
	settings, err := network.FetchSettings(client)

	if err != nil {
		return
	}

	genResult.Plugins = []string{}
	genResult.ProjectName = settings.ProjectName.DefaultVal
	initProjectDir()

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

	eventChan := make(chan tcell.Event)
	startTime := time.Now().UnixMicro()
	frameStart := startTime
	frameMs := 1000.0 / 30

	go func() {
		for running {
			eventChan <- scr.PollEvent()
		}
	}()

	for running {
		delta := float64(time.Now().UnixMicro()-startTime) / 1000.0
		startTime = time.Now().UnixMicro()

		if startTime-frameStart > 1e6 {
			frameStart = startTime
		}

		select {
		case event := <-eventChan:
			switch activeElement {
			case ProjectNameInput:
				processEvent(event, &genResult.ProjectName)
			case LocationInput:
				processEvent(event, &genResult.ProjectDir)
			case SearchInput:
				processEvent(event, &search)
			case Tabs:
				processEvent(event, nil)
			case CreateButton:
				processEvent(event, nil)
			default:
				panic("unhandled default case")
			}
		default:
			// do nothing
		}

		if pluginsShown && !pluginsFetched {
			var plugins []network.Plugin
			plugins, err = network.FetchPlugins(client, settings.KtorVersion.DefaultId)

			if err != nil {
				errorLine = "Unable to fetch plugins from the generation server. Please restart the app."
				pluginsFetched = true
				continue
			}

			pluginsFetched = true

			pluginDeps = make(map[string][]string, len(plugins))
			allPluginsByGroup = make(map[string][]network.Plugin, len(plugins))
			for _, p := range plugins {
				pluginDeps[p.Id] = p.RequiredPlugins

				if !slices.Contains(allSortedGroups, p.Group) {
					allSortedGroups = append(allSortedGroups, p.Group)
				}
				allPluginsByGroup[p.Group] = append(allPluginsByGroup[p.Group], p)
			}
			pluginsByGroup = allPluginsByGroup

			slices.Sort(allSortedGroups)
			groups = allSortedGroups
		}

		scr.Clear()
		scr.Fill(' ', defaultStyle)
		drawTui(scr, delta)
		scr.Show()

		if frameMs-delta > 0 {
			time.Sleep(time.Duration(frameMs-delta) * time.Millisecond)
		}
	}

	result = genResult
	return
}

func drawTui(scr tcell.Screen, deltaTime float64) {
	cursorAnimTimer += deltaTime

	defer func() {
		if cursorAnimTimer >= 1500 {
			cursorAnimTimer = 0
		}
	}()

	width, height := scr.Size()

	strongStyle := defaultStyle.Foreground(strongTextColor)
	cursorPos := cursorOffs[activeElement]
	padding := 1
	posX := padding
	posY := padding

	errY := height - 4
	if _, y := multilinePos(width/2, width, padding, errorLine); y > 0 {
		errY -= y
	}

	drawMultilineText(scr, width/2, errY, width, padding, defaultStyle.Foreground(errorColor), errorLine)
	drawInlineText(scr, width-len(statusLine)-1, height-2, defaultStyle.Foreground(statusColor), statusLine)

	posX, posY = drawInlineText(scr, posX, posY, strongStyle, "Project name:")
	posX++

	posX, posY = drawInput(scr, posX, posY, width-posX-padding, genResult.ProjectName, cursorPos, activeElement == ProjectNameInput)

	if !locationShown {
		return
	}

	posY += 2
	posX = padding
	posX, posY = drawInlineText(scr, posX, posY, strongStyle, "Location:")
	posX++

	drawInput(scr, posX, posY, width-posX-padding, genResult.ProjectDir, cursorPos, activeElement == LocationInput)

	if !pluginsShown {
		return
	}

	posY += 2
	posX = padding
	posX, posY = drawInlineText(scr, posX, posY, strongStyle, "Search for plugins:")
	posX++
	drawInput(scr, posX, posY, width-posX-padding, search, cursorPos, activeElement == SearchInput)

	if !pluginsFetched {
		return
	}

	posX = padding
	posY += 2

	tabVisRanges = tabVisRanges[:0]
	for off := 0; off < len(groups); {
		r := Range{start: off, end: off + getVisibleTabsCount(padding, width, groups[off:])}
		tabVisRanges = append(tabVisRanges, r)
		off = r.end
	}

	tabRange := findRange(tabVisRanges, activeTab)
	visibleTabs := groups[tabRange.start:tabRange.end]

	if tabRange.start > 0 {
		drawInlineText(scr, posX, posY, buttonStyle.Bold(true), "...")
		posX += 3
		scr.SetContent(posX, posY, tcell.RuneHLine, nil, textStyle.Bold(true))
		posX++
	}

	for i, gr := range visibleTabs {
		ps := pluginsByGroup[gr]

		style := buttonStyle
		if activeElement == Tabs && i == activeTab-tabRange.start {
			style = activeTabStyle
		}

		groupText := fmt.Sprintf("%s (%d)", gr, len(ps))
		posX, posY = drawInlineText(scr, posX, posY, style, groupText)

		if i != len(visibleTabs)-1 {
			scr.SetContent(posX, posY, tcell.RuneHLine, nil, textStyle.Bold(true))
		}

		posX += 1
	}

	if tabRange.end < len(groups) {
		posX--
		scr.SetContent(posX, posY, tcell.RuneHLine, nil, textStyle.Bold(true))
		drawInlineText(scr, posX+1, posY, buttonStyle.Bold(true), "...")
	}

	createStyle := buttonStyle
	if activeElement == CreateButton {
		createStyle = activeTabStyle
	}
	drawInlineText(scr, padding, height-2, createStyle, "CREATE PROJECT (ALT+ENTER)")

	if len(groups) == 0 {
		drawInlineText(scr, posX, posY, textStyle, "No plugins found by the search query")
		return
	}

	posX = padding + 2
	posY += 2
	pluginsXStart := posX

	activeGroup := groups[activeTab]
	plugins := pluginsByGroup[activeGroup]

	pluginVisRanges = pluginVisRanges[:0]
	for off := 0; off < len(plugins); {
		count := getVisiblePluginsCount(posY, height, plugins[off:], off)

		if count == 0 {
			errorLine = fmt.Sprintf("Terminal height %d is too small to display plugins", height)
			return
		}

		r := Range{start: off, end: off + count}
		pluginVisRanges = append(pluginVisRanges, r)
		off = r.end
	}

	plugsRange := findRange(pluginVisRanges, activePlugin)
	visPlugins := plugins[plugsRange.start:plugsRange.end]

	if plugsRange.start > 0 {
		scr.SetContent(padding, posY, tcell.RuneVLine, nil, textStyle.Bold(true))
		scr.SetContent(padding, posY, ' ', nil, buttonStyle)
		drawInlineText(scr, posX, posY, textStyle, "...")
		posY++
		scr.SetContent(padding, posY, tcell.RuneVLine, nil, textStyle.Bold(true))
		posY++
	}

	for i, p := range visPlugins {
		checkboxStyle := buttonStyle
		if activeElement == Tabs && i == activePlugin {
			checkboxStyle = activeTabStyle
		}

		checkboxVal := ' '
		if _, ok := addedPlugins[p.Id]; ok {
			checkboxVal = 'x'
		}

		scr.SetContent(padding, posY, checkboxVal, nil, checkboxStyle)
		if i != len(visPlugins)-1 {
			scr.SetContent(padding, posY+1, tcell.RuneVLine, nil, textStyle.Bold(true))
			scr.SetContent(padding, posY+2, tcell.RuneVLine, nil, textStyle.Bold(true))
		}

		nameStyle := textStyle
		if activeElement == Tabs && i+plugsRange.start == activePlugin {
			nameStyle = activeTabStyle
		}

		x := posX
		var searchIndices []int

		if len(search) > 0 {
			searchIndices = searchAll(p.Name, search)
		}

		for i, r := range []rune(p.Name) {
			style := nameStyle

			for _, ind := range searchIndices {
				if i >= ind && i < ind+len(search) {
					style = nameStyle.Foreground(searchHighlightColor)
					break
				}
			}

			scr.SetContent(x, posY, r, nil, style)
			x++
		}

		posY++

		descrStyle := weakTextStyle
		if activeElement == Tabs && i+plugsRange.start == activePlugin {
			descrStyle = weakTextStyle.Background(textColor)
		}

		searchIndices = searchIndices[:0]
		if len(search) > 0 {
			searchIndices = searchAll(p.Description, search)
		}

		x = pluginsXStart
		for i, r := range []rune(p.Description) {
			style := descrStyle

			for _, ind := range searchIndices {
				if i >= ind && i < ind+len(search) {
					style = descrStyle.Foreground(searchHighlightColor)
					break
				}
			}

			scr.SetContent(x, posY, r, nil, style)
			x++
		}
		posY += 2
	}

	if plugsRange.end < len(plugins) {
		posY -= 2
		scr.SetContent(padding, posY, tcell.RuneVLine, nil, textStyle.Bold(true))
		posY++
		scr.SetContent(padding, posY, tcell.RuneVLine, nil, textStyle.Bold(true))
		posY++
		drawInlineText(scr, padding+2, posY, textStyle, "...")
		scr.SetContent(padding, posY, ' ', nil, buttonStyle)
	}
}

func findRange(ranges []Range, index int) Range {
	for _, r := range ranges {
		if index >= r.start && index < r.end {
			return r
		}
	}

	return Range{0, 0}
}

func getVisiblePluginsCount(startY, height int, plugs []network.Plugin, off int) int {
	cy := startY
	count := 0
	createButtonSpace := 3
	dotsHeight := 3

	for i := range plugs {
		isLast := i == len(plugs)-1

		extraSpace := 0
		if (isLast && off > 0) || (!isLast && off == 0) {
			extraSpace = dotsHeight + 1
		} else if !isLast && off > 0 {
			extraSpace = dotsHeight * 2
		}

		if isLast {
			extraSpace += 2
		}

		if cy+createButtonSpace+extraSpace > height {
			break
		}
		count++

		cy += 2 // name + description
		cy += 1
	}

	return count
}

func getVisibleTabsCount(padding, width int, groups []string) int {
	cx := padding
	count := 0
	for i, gr := range groups {
		ps := pluginsByGroup[gr]

		groupText := fmt.Sprintf("%s (%d)", gr, len(ps))

		dotsSpace := 2 + 3
		if i == len(groups)-1 {
			dotsSpace = 0
		}

		if x, _ := inlinePos(cx, 0, groupText); x+padding+dotsSpace > width {
			break
		}
		count++

		cx, _ = inlinePos(cx, 0, groupText)
		cx += 1
	}

	return count
}

func drawInlineText(scr tcell.Screen, x, y int, style tcell.Style, text string) (int, int) {
	for _, r := range []rune(text) {
		scr.SetContent(x, y, r, nil, style)
		x++
	}

	return x, y
}

func multilinePos(x, width, padding int, text string) (int, int) {
	startX := x
	y := 0
	for range []rune(text) {
		x++

		if x >= width-padding {
			x = startX
			y++
		}
	}

	return x, y
}

func drawMultilineText(scr tcell.Screen, x, y, width, padding int, style tcell.Style, text string) {
	startX := x
	spaceThresh := 10
	for _, r := range []rune(text) {
		if unicode.IsSpace(r) && (x+spaceThresh >= width-padding) {
			x = startX
			y++
			continue
		}

		if x >= width-padding {
			x = startX
			y++
		}

		scr.SetContent(x, y, r, nil, style)
		x++
	}
}

func inlinePos(x, y int, text string) (int, int) {
	for range []rune(text) {
		x++
	}

	return x, y
}

func searchAll(s string, substr string) []int {
	sLow := strings.ToLower(s)
	substrLow := strings.ToLower(substr)
	var indices []int
	off := 0
	for i := strings.Index(sLow, substrLow); i < len(sLow) && i >= 0; i = strings.Index(sLow[off:], substrLow) {
		indices = append(indices, i+off)
		i += len(substrLow)
		off += i
	}

	return indices
}

var firstResize = false

func processEvent(ev tcell.Event, input *string) {
	inputOff := cursorOffs[activeElement]

	switch ev := ev.(type) {
	case *tcell.EventResize:
		if !firstResize {
			firstResize = true
			break
		}

		errorLine = ""
	case *tcell.EventKey:
		mod, key := ev.Modifiers(), ev.Key()

		cursorAnimTimer = 0

		switch {
		case (mod == tcell.ModCtrl && key == tcell.KeyCtrlC) || (key == tcell.KeyEscape):
			running = false
			genResult.Quit = true
		case key == tcell.KeyRune:
			if mod == tcell.ModAlt && ev.Rune() == 'm' {
				inputColor++
				statusLine = fmt.Sprintf("Color%d", inputColor-tcell.ColorValid)
				return
			}

			if mod == tcell.ModAlt && ev.Rune() == 'n' {
				inputColor--
				statusLine = fmt.Sprintf("Color%d", inputColor-tcell.ColorValid)
				return
			}

			if activeElement == Tabs && ev.Rune() == ' ' {
				toggleSelectedPlugin()
				statusLine = fmt.Sprintf("%d plugins selected", len(addedPlugins))
				return
			}

			if input != nil {
				*input = insertRune(*input, inputOff, ev.Rune())
				cursorOffs[activeElement] = moveCursor(inputOff, len(*input), 1)
				onInputChanged(activeElement, *input)
			}
		case key == tcell.KeyLeft:
			if activeElement == Tabs {
				if activeTab-1 >= 0 {
					activeTab--
					activePlugin = 0
				}

				return
			}

			if input != nil {
				cursorOffs[activeElement] = moveCursor(inputOff, len(*input), -1)
			}
		case key == tcell.KeyRight:
			if activeElement == Tabs {
				if activeTab+1 < len(groups) {
					activeTab++
					activePlugin = 0
				}

				return
			}

			if input != nil {
				cursorOffs[activeElement] = moveCursor(inputOff, len(*input), 1)
			}
		case key == tcell.KeyUp:
			if activeElement == Tabs {
				if activePlugin == 0 {
					activeElement = prevElement()
					return
				}

				if activePlugin-1 >= 0 {
					activePlugin--
				}
				return
			}

			activeElement = prevElement()
		case key == tcell.KeyDown:
			if activeElement == Tabs {
				if activePlugin+1 > len(pluginsOnCurrentTab())-1 {
					activeElement = nextElement()
					return
				}

				if activePlugin+1 < len(pluginsOnCurrentTab()) {
					activePlugin++
				}
				return
			}

			switch {
			case activeElement == ProjectNameInput && locationShown:
				activeElement = nextElement()
			case activeElement == LocationInput && pluginsShown:
				activeElement = nextElement()
			case locationShown && pluginsShown:
				activeElement = nextElement()
			default:
				// do nothing yet
			}
		case key == tcell.KeyDEL: // Backspace
			if input != nil {
				*input = deleteChar(*input, inputOff-1)
				cursorOffs[activeElement] = moveCursor(inputOff, len(*input), -1)
				onInputChanged(activeElement, *input)
			}
		case key == tcell.KeyDelete:
			if input != nil {
				*input = deleteChar(*input, inputOff)
				cursorOffs[activeElement] = moveCursor(inputOff, len(*input), -1)
				onInputChanged(activeElement, *input)
			}
		case key == tcell.KeyEnter && mod == tcell.ModAlt && pluginsShown: // Generate project
			for id := range addedPlugins {
				genResult.Plugins = append(genResult.Plugins, id)
			}

			if isDirEmptyOrAbsent(genResult.ProjectDir) {
				running = false
				return
			}
		case key == tcell.KeyEnter && mod == tcell.ModNone:
			if activeElement == Tabs {
				toggleSelectedPlugin()
				statusLine = fmt.Sprintf("%d plugins selected", len(addedPlugins))
				return
			} else if activeElement == CreateButton { // Generate project
				for id := range addedPlugins {
					genResult.Plugins = append(genResult.Plugins, id)
				}

				if isDirEmptyOrAbsent(genResult.ProjectDir) {
					running = false
					return
				}
			}

			switch activeElement {
			case ProjectNameInput:
				locationShown = true
				initProjectDir()
				checkProjectDir(genResult.ProjectDir)
			case LocationInput:
				pluginsShown = true
			default:
				// do nothing yet
			}

			activeElement = nextElement()
		case key == tcell.KeyTab:
			switch activeElement {
			case ProjectNameInput:
				locationShown = true
				initProjectDir()
				checkProjectDir(genResult.ProjectDir)
			case LocationInput:
				pluginsShown = true
			default:
				// do nothing yet
			}

			activeElement = nextElement()
		case key == tcell.KeyBacktab: // Shift + Tab
			activeElement = prevElement()
		}
	}
}

func searchPlugins() map[string][]network.Plugin {
	m := make(map[string][]network.Plugin)
	groups = groups[:0]
	for gr, ps := range allPluginsByGroup {
		var plugins []network.Plugin

		for _, p := range ps {
			if strings.Contains(strings.ToLower(p.Name), strings.ToLower(search)) || strings.Contains(strings.ToLower(p.Description), strings.ToLower(search)) {
				plugins = append(plugins, p)
			}
		}

		if len(plugins) > 0 {
			slices.SortFunc(plugins, func(a, b network.Plugin) int {
				return strings.Compare(a.Name, b.Name)
			})
			m[gr] = plugins
			groups = append(groups, gr)
		}
	}

	if activeTab >= len(groups) {
		activeTab = len(groups) - 1
	}

	if activeTab < 0 {
		activeTab = 0
	}

	slices.Sort(groups)

	return m
}

func onInputChanged(element Element, input string) {
	if element == SearchInput {
		search = input
		pluginsByGroup = searchPlugins()
		return
	}

	if element == LocationInput {
		checkProjectDir(input)
	}
}

func checkProjectDir(dir string) {
	if !isDirEmptyOrAbsent(dir) {
		errorLine = fmt.Sprintf("Directory %s isn't empty", dir)
	} else {
		errorLine = ""
	}
}

func isDirEmptyOrAbsent(dir string) bool {
	f, err := os.Open(dir)
	if err != nil {
		return true
	}
	defer f.Close()

	_, err = f.Readdirnames(1)
	if err == io.EOF {
		return true
	}
	return false
}

func initProjectDir() {
	wd, err := os.Getwd()

	if err != nil {
		genResult.ProjectDir = filepath.Join(".", genResult.ProjectName)
		return
	}

	genResult.ProjectDir = filepath.Join(wd, genResult.ProjectName)
}

func toggleSelectedPlugin() {
	p := pluginsOnCurrentTab()[activePlugin]

	if _, ok := addedPlugins[p.Id]; ok { // Plugin is selected
		for id := range addedPlugins {
			if isPluginRequiredFor(p.Id, id) {
				delete(addedPlugins, id)
			}
		}

		delete(addedPlugins, p.Id)

		if ips, ok := indirectPlugins[p.Id]; ok {
			for id := range ips {
				delete(addedPlugins, id)
			}
		}
		delete(indirectPlugins, p.Id)
		return
	}

	addedPlugins[p.Id] = struct{}{}
	deps := make(idSet)
	getDepPlugins(p.Id, deps)

	for _, s := range indirectPlugins {
		if _, ok := s[p.Id]; ok {
			delete(s, p.Id)
		}
	}

	for id := range deps {
		if _, ok := addedPlugins[id]; ok {
			delete(deps, id)
		}
	}

	for id := range deps {
		addedPlugins[id] = struct{}{}
	}

	indirectPlugins[p.Id] = deps
}

func isPluginRequiredFor(parentId string, childId string) bool {
	for _, id := range pluginDeps[childId] {
		if id == parentId || isPluginRequiredFor(id, childId) {
			return true
		}
	}
	return false
}

func getDepPlugins(id string, m idSet) {
	for _, dp := range pluginDeps[id] {
		m[dp] = struct{}{}
		getDepPlugins(dp, m)
	}
}

func pluginsOnCurrentTab() []network.Plugin {
	return pluginsByGroup[groups[activeTab]]
}

func nextElement() Element {
	el := Element(int(activeElement) + 1)
	if el != Last {
		return el
	}

	return activeElement
}

func prevElement() Element {
	elIndex := int(activeElement) - 1
	if elIndex >= 0 {
		return Element(elIndex)
	}

	return activeElement
}

func deleteChar(input string, pos int) string {
	if pos >= len(input) || pos < 0 {
		return input
	}

	return fmt.Sprintf("%s%s", input[0:pos], input[pos+1:])
}

func insertRune(input string, pos int, r rune) string {
	if pos < 0 {
		return input
	}

	if input == "" {
		return fmt.Sprintf("%c", r)
	}

	if pos >= len(input) {
		return fmt.Sprintf("%s%c", input, r)
	}

	return fmt.Sprintf("%s%c%s", input[0:pos], r, input[pos:])
}

func moveCursor(off, inputLen, delta int) int {
	off += delta
	if off > inputLen {
		off = inputLen
	}

	if off < 0 {
		off = 0
	}

	return off
}

func drawInput(scr tcell.Screen, posX, posY int, inputLen int, input string, cursorPos int, focused bool) (int, int) {
	inputStart := posX
	for i := posX; i < posX+inputLen; i++ {
		scr.SetContent(i, posY, ' ', nil, inputStyle)
	}

	for i, r := range []rune(input) {
		style := inputStyle
		if focused && i == cursorPos && cursorAnimTimer < 700 {
			style = cursorStyle
		}

		scr.SetContent(posX, posY, r, nil, style)

		posX++
	}

	if focused && cursorPos >= len(input) && cursorAnimTimer < 700 {
		scr.SetContent(inputStart+cursorPos, posY, ' ', nil, cursorStyle)
	}

	return posX, posY
}
