package interactive

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/ktorio/ktor-cli/internal/app/network"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"time"
)

var defaultStyle = tcell.StyleDefault.Background(tcell.Color233)
var inputStyle = defaultStyle.Background(tcell.Color117).Foreground(tcell.Color233)
var cursorStyle = inputStyle.Background(tcell.ColorWhite)
var buttonStyle = defaultStyle.Background(tcell.Color163).Foreground(tcell.ColorWhite)
var activeTabStyle = defaultStyle.Foreground(tcell.Color163).Background(tcell.ColorWhite)
var textStyle = defaultStyle.Foreground(tcell.ColorWhite)
var weakTextStyle = defaultStyle.Foreground(tcell.Color139)

type Element int

const (
	ProjectNameInput Element = iota
	LocationInput
	SearchInput
	Tabs
	Last
)
const projectInputLen = 64
const locationInputLen = 128
const searchInputLen = 48

var running bool
var cursorOffs map[Element]int
var locationShown bool
var pluginsShown bool
var activeElement Element
var cursorAnimTimer float64 = 0
var pluginsByGroup map[string][]network.Plugin
var pluginsFetched bool
var sortedGroups []string
var activeTab int
var tabsCount int
var activePlugin int
var addedPlugins []string

func init() {
	running = true
	cursorOffs = map[Element]int{
		ProjectNameInput: 0,
		LocationInput:    0,
		SearchInput:      0,
	}
	locationShown = true
	pluginsShown = true
	activeElement = ProjectNameInput
	pluginsFetched = false
	activeTab = 0
	activePlugin = 0
}

func Run(client *http.Client) error {
	settings, err := network.FetchSettings(client)

	if err != nil {
		// TODO: handle error
		return err
	}

	projectName := settings.ProjectName.DefaultVal

	wd, err := os.Getwd()

	if err != nil {
		return err
	}

	location := filepath.Join(wd, projectName)
	searchStr := ""

	scr, err := tcell.NewScreen()

	if err != nil {
		// TODO: handle error
		return err
	}

	if err := scr.Init(); err != nil {
		return err
	}

	scr.EnableMouse()
	scr.Clear()
	//scrWidth, scrHeight := scr.Size()

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
				processEvent(event, &projectName)
			case LocationInput:
				processEvent(event, &location)
			case SearchInput:
				processEvent(event, &searchStr)
			case Tabs:
				processEvent(event, &searchStr)
			default:
				panic("unhandled default case")
			}
		default:
			// do nothing
		}

		if pluginsShown && !pluginsFetched {
			plugins, err := network.FetchPlugins(client, settings.KtorVersion.DefaultId)

			if err != nil {
				// TODO: handle error
				return err
			}

			pluginsFetched = true

			pluginsByGroup = make(map[string][]network.Plugin, len(plugins))
			for _, p := range plugins {
				if !slices.Contains(sortedGroups, p.Group) {
					sortedGroups = append(sortedGroups, p.Group)
				}
				pluginsByGroup[p.Group] = append(pluginsByGroup[p.Group], p)
			}

			tabsCount = len(sortedGroups)
			slices.Sort(sortedGroups)
		}

		scr.Clear()
		scr.Fill(' ', defaultStyle)
		drawTui(scr, delta, projectName, location, searchStr)
		scr.Show()

		if frameMs-delta > 0 {
			time.Sleep(time.Duration(frameMs-delta) * time.Millisecond)
		}
	}

	return nil
}

func drawTui(scr tcell.Screen, deltaTime float64, projectName string, location string, searchStr string) {
	cursorAnimTimer += deltaTime

	defer func() {
		if cursorAnimTimer >= 1500 {
			cursorAnimTimer = 0
		}
	}()

	strongStyle := defaultStyle.Foreground(tcell.Color141)
	cursorPos := cursorOffs[activeElement]
	padding := 1
	posX := padding
	posY := padding
	posX, posY = drawInlineText(scr, posX, posY, strongStyle, "Project name:")
	posX++

	posX, posY = drawInput(scr, posX, posY, projectInputLen, projectName, cursorPos, activeElement == ProjectNameInput)

	if !locationShown {
		return
	}

	posY += 2
	posX = padding
	posX, posY = drawInlineText(scr, posX, posY, strongStyle, "Location:")
	posX++

	drawInput(scr, posX, posY, locationInputLen, location, cursorPos, activeElement == LocationInput)

	if !pluginsShown {
		return
	}

	posY += 2
	posX = padding
	posX, posY = drawInlineText(scr, posX, posY, strongStyle, "Search for plugins:")
	posX++
	drawInput(scr, posX, posY, searchInputLen, searchStr, cursorPos, activeElement == SearchInput)

	if !pluginsFetched {
		return
	}

	posX = padding
	posY += 2
	for i, gr := range sortedGroups {
		ps := pluginsByGroup[gr]

		style := buttonStyle
		if activeElement == Tabs && i == activeTab {
			style = activeTabStyle
		}

		posX, posY = drawInlineText(scr, posX, posY, style, fmt.Sprintf("%s (%d)", gr, len(ps)))

		if i != len(sortedGroups)-1 {
			scr.SetContent(posX, posY, tcell.RuneHLine, nil, textStyle.Bold(true))
		}

		posX += 1
	}

	posX = padding
	posY++
	scr.SetContent(posX, posY, tcell.RuneVLine, nil, textStyle.Bold(true))
	posY++

	posX += 2
	pluginsXStart := posX
	activeGroup := sortedGroups[activeTab]

	for i, p := range pluginsByGroup[activeGroup] {
		checkboxStyle := buttonStyle
		if activeElement == Tabs && i == activePlugin {
			checkboxStyle = activeTabStyle
		}

		checkboxVal := ' '
		if slices.Contains(addedPlugins, p.Id) {
			checkboxVal = 'x'
		}
		scr.SetContent(padding, posY, checkboxVal, nil, checkboxStyle)
		scr.SetContent(padding, posY+1, tcell.RuneVLine, nil, textStyle.Bold(true))
		scr.SetContent(padding, posY+2, tcell.RuneVLine, nil, textStyle.Bold(true))

		nameStyle := textStyle
		if activeElement == Tabs && i == activePlugin {
			nameStyle = activeTabStyle
		}
		drawInlineText(scr, posX, posY, nameStyle, p.Name)
		posY++

		descrStyle := weakTextStyle
		if activeElement == Tabs && i == activePlugin {
			descrStyle = weakTextStyle.Background(tcell.ColorWhite)
		}

		drawInlineText(scr, pluginsXStart, posY, descrStyle, p.Description)
		posY += 2
	}
}

func drawInlineText(scr tcell.Screen, x, y int, style tcell.Style, text string) (int, int) {
	for _, r := range []rune(text) {
		scr.SetContent(x, y, r, nil, style)
		x++
	}

	return x, y
}

func processEvent(ev tcell.Event, input *string) {
	inputOff := cursorOffs[activeElement]

	switch ev := ev.(type) {
	case *tcell.EventKey:
		mod, key := ev.Modifiers(), ev.Key()

		cursorAnimTimer = 0

		switch {
		case (mod == tcell.ModCtrl && key == tcell.KeyCtrlC) || (key == tcell.KeyEscape):
			running = false
		case key == tcell.KeyRune:
			if activeElement == Tabs && ev.Rune() == ' ' {
				toggleSelectedPlugin()
				return
			}

			*input = insertRune(*input, inputOff, ev.Rune())
			cursorOffs[activeElement] = moveCursor(inputOff, len(*input), 1)
		case key == tcell.KeyLeft:
			if activeElement == Tabs {
				if activeTab-1 >= 0 {
					activeTab--
					activePlugin = 0
				}

				return
			}

			cursorOffs[activeElement] = moveCursor(inputOff, len(*input), -1)
		case key == tcell.KeyRight:
			if activeElement == Tabs {
				if activeTab+1 < tabsCount {
					activeTab++
					activePlugin = 0
				}

				return
			}

			cursorOffs[activeElement] = moveCursor(inputOff, len(*input), 1)
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
				if activePlugin+1 < len(pluginsByGroup[sortedGroups[activeTab]]) {
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
			*input = deleteChar(*input, inputOff-1)
			cursorOffs[activeElement] = moveCursor(inputOff, len(*input), -1)
		case key == tcell.KeyDelete:
			*input = deleteChar(*input, inputOff)
			cursorOffs[activeElement] = moveCursor(inputOff, len(*input), -1)
		case key == tcell.KeyEnter:
			if activeElement == Tabs {
				toggleSelectedPlugin()
				return
			}

			switch activeElement {
			case ProjectNameInput:
				locationShown = true
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

func toggleSelectedPlugin() {
	p := pluginsByGroup[sortedGroups[activeTab]][activePlugin]
	if pIndex := slices.Index(addedPlugins, p.Id); pIndex >= 0 {
		addedPlugins = slices.Delete(addedPlugins, pIndex, pIndex+1)
	} else {
		addedPlugins = append(addedPlugins, p.Id)
	}
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
