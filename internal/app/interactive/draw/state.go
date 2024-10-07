package draw

import "github.com/gdamore/tcell/v2"

func NewState(scr tcell.Screen) *State {
	return &State{scr: scr}
}

type State struct {
	scr         tcell.Screen
	Padding     int
	CurrentPosX int
	CurrentPosY int
}

func (st *State) Move(x, y int) {
	st.CurrentPosX = x
	st.CurrentPosY = y
}

func (st *State) Text(text string, style tcell.Style) {
	cx, cy := st.CurrentPosX, st.CurrentPosY

	if cx == 0 {
		cx += st.Padding
	}

	if cy == 0 {
		cy += st.Padding
	}

	x, y := Text(st.scr, text, cx, cy, style)
	st.CurrentPosX = x
	st.CurrentPosY = y
}

func (st *State) MoveY(offY int) {
	st.CurrentPosY += offY + st.Padding
}

func (st *State) ResetPos() {
	st.CurrentPosX = 0
	st.CurrentPosY = 0
}
