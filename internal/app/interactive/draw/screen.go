package draw

import (
	"github.com/gdamore/tcell/v2"
)

//type Screen struct {
//	Width  int
//	Height int
//	Scr    tcell.Screen
//}
//
//func NewScreen(scr tcell.Screen) *Screen {
//	return &Screen{Scr: scr}
//}

//func (s *Screen) DrawStatus(msg string) {
//	//style := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorWhiteSmoke).Bold(true)
//	//maxLen := s.Width / 3
//	//if s.Width <= 80 {
//	//	maxLen = s.Width - 20
//	//}
//	//
//	//s.Text(msg, s.Width-maxLen, s.Height-1, style)
//}

//func Element(scr tcell.Screen, el interface{}) {
//	switch e := el.(type) {
//	case element.Vertical:
//		for _, ch := range e.Children {
//
//		}
//	}
//}

func Text(scr tcell.Screen, text string, x, y int, style tcell.Style) (int, int) {
	for _, r := range []rune(text) {
		scr.SetContent(x, y, r, nil, style)
		x++
	}

	return x, y
}
