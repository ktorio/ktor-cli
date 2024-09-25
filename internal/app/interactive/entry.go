package interactive

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/ktorio/ktor-cli/internal/app/network"
	"net/http"
)

func Run(client *http.Client) error {
	scr, err := tcell.NewScreen()

	if err != nil {
		return err
	}

	if err := scr.Init(); err != nil {
		return err
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

	width, height := scr.Size()
	surface := Surface{Width: width, Height: height, Scr: scr}

	settings, err := network.FetchSettings(client)

	if err != nil {
		// TODO: handle error
		return err
	}

	ktorVersion := settings.KtorVersion.DefaultId

	plugins, err := network.FetchPlugins(client, ktorVersion)

	if err != nil {
		// TODO: handle error
		return err
	}

	fmt.Println(plugins)

	for {
		scr.Show()

		ev := scr.PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventResize:
			scr.Sync()
		case *tcell.EventKey:
			mod, key, ch := ev.Modifiers(), ev.Key(), ev.Rune()
			if mod == tcell.ModCtrl && key == tcell.KeyCtrlC {
				return nil
			}

			log(scr, fmt.Sprintf("EventKey Modifiers: %d Key: %d Rune: %d", mod, key, ch), 0, 1)

			if ev.Key() == tcell.KeyEscape {
				return nil
			}
		case *tcell.EventMouse:
			mod := ev.Modifiers()
			btns := ev.Buttons()
			x, y := ev.Position()

			if btns == tcell.Button1 {
				scr.SetContent(x, y, ' ', nil, tcell.StyleDefault.Background(tcell.ColorGreen))
			}

			surface.DrawStatus(fmt.Sprintf("EventMouse Modifiers: %d Buttons: %d Position: %d,%d", mod, btns, x, y))
		}

	}
}

type Surface struct {
	Width  int
	Height int
	Scr    tcell.Screen
}

func (s *Surface) DrawStatus(msg string) {
	style := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorWhiteSmoke).Bold(true)
	maxLen := s.Width / 3
	if s.Width <= 80 {
		maxLen = s.Width - 20
	}

	s.Scr.Clear()
	drawText(s.Scr, s.Width-maxLen, s.Height-1, s.Width, s.Height-1, style, msg)
}

func log(s tcell.Screen, msg string, x, y int) {
	drawText(s, x, y, len(msg), y+2, tcell.StyleDefault, msg)
}

func drawText(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string) {
	row := y1
	col := x1
	for _, r := range []rune(text) {
		s.SetContent(col, row, r, nil, style)
		col++
		if col >= x2 {
			row++
			col = x1
		}
		if row > y2 {
			break
		}
	}
}
