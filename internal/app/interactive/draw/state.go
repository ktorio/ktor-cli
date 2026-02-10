package draw

import "github.com/gdamore/tcell/v2"

type Range struct {
	start, end int
}

type CursorGlobalInfo struct {
	R    rune
	X, Y int
}

type State struct {
	CursorOffs      map[Element]int
	ActiveElement   Element
	PluginsShown    bool
	tabVisRanges    []Range
	pluginVisRanges []Range
	ActiveTab       int
	ActivePlugin    int
	VisibleOffs     map[Element]int
	InputLens       map[Element]int
}

type Element int

const (
	ProjectNameInput Element = iota
	LocationInput
	VersionCatalogCheckbox
	SearchInput
	Tabs
	CreateButton
	LastElement
)

func NewState() *State {
	return &State{
		CursorOffs: map[Element]int{
			ProjectNameInput: 0,
			LocationInput:    0,
			SearchInput:      0,
		},
		VisibleOffs: map[Element]int{
			ProjectNameInput: 0,
			LocationInput:    0,
			SearchInput:      0,
		},
		InputLens: map[Element]int{
			ProjectNameInput: 0,
			LocationInput:    0,
			SearchInput:      0,
		},
	}
}

func ActiveInputOffset(st *State) int {
	return st.CursorOffs[st.ActiveElement]
}

func IsElementActive(st *State, element Element) bool {
	return st.ActiveElement == element
}

func MoveCursor(st *State, inputLen, delta int) {
	off := ActiveInputOffset(st)

	off += delta
	if off > inputLen {
		off = inputLen
	}

	if off < 0 {
		off = 0
	}

	st.CursorOffs[st.ActiveElement] = off
}

func SwitchTab(st *State, numTabs int, delta int) {
	newTab := st.ActiveTab + delta

	if newTab < 0 {
		return
	}

	if newTab >= numTabs {
		return
	}

	st.ActiveTab = newTab
	st.ActivePlugin = 0
}

func SwitchElement(st *State, delta int) {
	newElement := int(st.ActiveElement) + delta

	if newElement < 0 {
		return
	}

	if newElement >= int(LastElement) {
		return
	}

	st.ActiveElement = Element(newElement)
}

func (st *State) VisOff() int {
	return st.VisibleOffs[st.ActiveElement]
}

func (st *State) CursorPos() int {
	return st.CursorOffs[st.ActiveElement]
}

func (st *State) InputLen() int {
	return st.InputLens[st.ActiveElement]
}

func HideCursorIfNeeded(st *State, scr tcell.Screen) {
	if st.ActiveElement != ProjectNameInput && st.ActiveElement != LocationInput && st.ActiveElement != SearchInput {
		scr.HideCursor()
	}
}
