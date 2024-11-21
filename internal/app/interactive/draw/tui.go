package draw

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/ktorio/ktor-cli/internal/app/i18n"
	"github.com/ktorio/ktor-cli/internal/app/interactive/model"
	"github.com/ktorio/ktor-cli/internal/app/network"
	"strings"
)

var scrWidth = 0
var scrHeight = 0

func Tui(scr tcell.Screen, st *State, mdl *model.State) {
	width, height := scr.Size()

	if scrWidth != width || scrHeight != height {
		for _, inp := range []Element{ProjectNameInput, LocationInput, SearchInput} {
			st.CursorOffs[inp] = 0
			st.VisibleOffs[inp] = 0
		}

		scrWidth = width
		scrHeight = height
	}

	strongStyle := DefaultStyle.Foreground(strongTextColor)
	padding := 1
	posX := padding
	posY := padding

	drawInlineText(scr, width-len(mdl.StatusLine)-1, height-2, DefaultStyle.Foreground(statusColor), mdl.StatusLine)

	posX, posY = drawInlineText(scr, posX, posY, strongStyle, i18n.Get(i18n.ProjectNameCaption))
	posX++
	projectInputStartX := posX

	st.InputLens[ProjectNameInput] = width - posX - padding
	posX, posY = drawInput(scr, st, projectInputStartX, posY, mdl.ProjectName, ProjectNameInput)

	if errs := mdl.GetErrors(model.ProjectNameEmptyError, model.ProjectNameAllowedCharsError); len(errs) > 0 {
		drawInlineTextEllipsis(scr, projectInputStartX, posY+1, width, padding, errorStyle, strings.Join(errs, ". "))
		posY++
	}

	posY += 2
	posX = padding
	posX, posY = drawInlineText(scr, posX, posY, strongStyle, i18n.Get(i18n.LocationCaption))
	posX++

	st.InputLens[LocationInput] = width - posX - padding
	drawInput(scr, st, posX, posY, mdl.ProjectDir, LocationInput)

	posY++
	drawInlineTextEllipsis(scr, posX, posY, width, padding, subTextStyle, fmt.Sprintf(i18n.Get(i18n.ProjectCreatedIn, mdl.GetProjectPath())))

	if errs := mdl.GetErrors(model.ProjectDirNotEmptyError, model.DirNotExistError, model.ProjectDirTooLongError); len(errs) > 0 {
		drawInlineTextEllipsis(scr, posX, posY+1, width, padding, errorStyle, strings.Join(errs, ". "))
		posY++
	}

	if !st.PluginsShown {
		return
	}

	posY += 2
	posX = padding

	if errs := mdl.GetErrors(model.UnableFetchPluginsError); len(errs) > 0 {
		drawInlineTextEllipsis(scr, posX, posY, width, padding, errorStyle, strings.Join(errs, ". "))
		return
	}

	posX, posY = drawInlineText(scr, posX, posY, strongStyle, i18n.Get(i18n.SearchPluginsCaption))
	posX++
	st.InputLens[SearchInput] = width - posX - padding
	drawInput(scr, st, posX, posY, mdl.Search, SearchInput)

	if !mdl.PluginsFetched {
		return
	}

	posX = padding
	posY += 2

	st.tabVisRanges = st.tabVisRanges[:0]
	for off := 0; off < len(mdl.Groups); {
		r := Range{start: off, end: off + getVisibleTabsCount(padding, width, mdl.PluginsByGroup, mdl.Groups[off:])}
		st.tabVisRanges = append(st.tabVisRanges, r)
		off = r.end
	}

	tabRange := findRange(st.tabVisRanges, st.ActiveTab)
	visibleTabs := mdl.Groups[tabRange.start:tabRange.end]

	if tabRange.start > 0 {
		drawInlineText(scr, posX, posY, buttonStyle.Bold(true), "...")
		posX += 3
		scr.SetContent(posX, posY, tcell.RuneHLine, nil, textStyle.Bold(true))
		posX++
	}

	for i, gr := range visibleTabs {
		ps := mdl.PluginsByGroup[gr]

		style := buttonStyle
		if st.ActiveElement == Tabs && i == st.ActiveTab-tabRange.start {
			style = activeTabStyle
		}

		groupText := fmt.Sprintf("%s (%d)", gr, len(ps))
		posX, posY = drawInlineText(scr, posX, posY, style, groupText)

		if i != len(visibleTabs)-1 {
			scr.SetContent(posX, posY, tcell.RuneHLine, nil, textStyle.Bold(true))
		}

		posX += 1
	}

	if tabRange.end < len(mdl.Groups) {
		posX--
		scr.SetContent(posX, posY, tcell.RuneHLine, nil, textStyle.Bold(true))
		drawInlineText(scr, posX+1, posY, buttonStyle.Bold(true), "...")
	}

	createStyle := buttonStyle
	if st.ActiveElement == CreateButton {
		createStyle = activeTabStyle
	}
	drawInlineText(scr, padding, height-2, createStyle, i18n.Get(i18n.CreateProjectButton))

	if len(mdl.Groups) == 0 {
		drawInlineText(scr, posX, posY, textStyle, i18n.Get(i18n.NoPluginsFound))
		return
	}

	posX = padding + 2
	posY += 2
	pluginsXStart := posX

	activeGroup := mdl.Groups[st.ActiveTab]
	plugins := mdl.PluginsByGroup[activeGroup]

	st.pluginVisRanges = st.pluginVisRanges[:0]
	for off := 0; off < len(plugins); {
		count := getVisiblePluginsCount(posY, height, plugins[off:], off)

		if count == 0 {
			msg := fmt.Sprintf(i18n.Get(i18n.TermHeightSmall, height))
			mdl.SetError(model.TerminalHeightError, fmt.Sprintf(i18n.Get(i18n.TermHeightSmall, height)))
			drawInlineTextEllipsis(scr, padding, posY, width, padding, errorStyle, msg)
			return
		} else {
			mdl.RemoveErrors(model.TerminalHeightError)
		}

		r := Range{start: off, end: off + count}
		st.pluginVisRanges = append(st.pluginVisRanges, r)
		off = r.end
	}

	plugsRange := findRange(st.pluginVisRanges, st.ActivePlugin)
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
		if st.ActiveElement == Tabs && i == st.ActivePlugin {
			checkboxStyle = activeTabStyle
		}

		checkboxVal := ' '
		if _, ok := mdl.AddedPlugins[p.Id]; ok {
			checkboxVal = 'x'
		}

		scr.SetContent(padding, posY, checkboxVal, nil, checkboxStyle)
		if i != len(visPlugins)-1 {
			scr.SetContent(padding, posY+1, tcell.RuneVLine, nil, textStyle.Bold(true))
			scr.SetContent(padding, posY+2, tcell.RuneVLine, nil, textStyle.Bold(true))
		}

		nameStyle := textStyle
		if st.ActiveElement == Tabs && i+plugsRange.start == st.ActivePlugin {
			nameStyle = activeStyle
		}

		x := posX
		var searchIndices []int

		if len(mdl.Search) > 0 {
			searchIndices = searchAll(p.Name, mdl.Search)
		}

		for i, r := range []rune(p.Name) {
			style := nameStyle

			for _, ind := range searchIndices {
				if i >= ind && i < ind+len(mdl.Search) {
					style = nameStyle.Foreground(searchHighlightColor)
					break
				}
			}

			scr.SetContent(x, posY, r, nil, style)
			x++
		}

		posY++

		descrStyle := weakTextStyle
		if st.ActiveElement == Tabs && i+plugsRange.start == st.ActivePlugin {
			descrStyle = weakTextStyle.Background(activeColor).Foreground(bgColor)
		}

		searchIndices = searchIndices[:0]
		if len(mdl.Search) > 0 {
			searchIndices = searchAll(p.Description, mdl.Search)
		}

		x = pluginsXStart
		for i, r := range []rune(p.Description) {
			style := descrStyle

			for _, ind := range searchIndices {
				if i >= ind && i < ind+len(mdl.Search) {
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

func drawInlineText(scr tcell.Screen, x, y int, style tcell.Style, text string) (int, int) {
	for _, r := range []rune(text) {
		scr.SetContent(x, y, r, nil, style)
		x++
	}

	return x, y
}

func drawInlineTextEllipsis(scr tcell.Screen, x, y, width, padding int, style tcell.Style, text string) {
	c := 0
	complete := true
	for i, r := range []rune(text) {
		c = i

		ellipsisLen := len("...")

		if i+x+len(text[i+1:])+padding < width {
			ellipsisLen = 0
		}

		if i+x+padding+ellipsisLen >= width {
			complete = false
			break
		}

		scr.SetContent(i+x, y, r, nil, style)
	}

	if !complete {
		drawInlineText(scr, c+x, y, style, "...")
	}
}

func drawInput(scr tcell.Screen, st *State, posX, posY int, input string, el Element) (int, int) {
	for i := posX; i < posX+st.InputLens[el]; i++ {
		scr.SetContent(i, posY, ' ', nil, inputStyle)
	}

	focused := st.ActiveElement == el
	cursorPos := st.CursorOffs[st.ActiveElement]
	visOff := st.VisibleOffs[el]

	runes := []rune(input)
	start := visOff
	end := min(len(runes), visOff+st.InputLens[el]-1)

	if end-start >= st.InputLens[el] {
		panic(fmt.Sprintf("%d:%d = %d out of len %d", start, end, end-start, st.InputLens[el]))
	}

	for i, r := range append(runes[start:end], ' ') {
		style := inputStyle

		if focused && i == cursorPos-visOff {
			scr.ShowCursor(posX, posY)
		}
		scr.SetContent(posX, posY, r, nil, style)

		posX++
	}

	return posX, posY
}

func getVisibleTabsCount(padding, width int, pluginsByGroup map[string][]network.Plugin, groups []string) int {
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

func inlinePos(x, y int, text string) (int, int) {
	for range []rune(text) {
		x++
	}

	return x, y
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
