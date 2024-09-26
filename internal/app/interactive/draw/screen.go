package draw

import (
	"github.com/gdamore/tcell/v2"
	"github.com/ktorio/ktor-cli/internal/app/interactive/state"
)

type Screen struct {
	Width  int
	Height int
	Scr    tcell.Screen
}

func NewScreen(scr tcell.Screen) *Screen {
	return &Screen{Scr: scr}
}

func (s *Screen) DrawStatus(msg string) {
	//style := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorWhiteSmoke).Bold(true)
	//maxLen := s.Width / 3
	//if s.Width <= 80 {
	//	maxLen = s.Width - 20
	//}
	//
	//s.DrawText(msg, s.Width-maxLen, s.Height-1, style)
}

func (s *Screen) DrawText(view state.View, text string, x, y int, style tcell.Style) {
	cx, cy := x-view.X, y-view.Y

	for _, r := range []rune(text) {
		s.Scr.SetContent(cx, cy, r, nil, style)
		cx++
		if cx >= view.X+view.Width {
			cy++
			cx = x
		}
	}
}
