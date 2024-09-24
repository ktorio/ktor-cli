package interactive

import (
	"github.com/gdamore/tcell/v2"
)

func Run() error {
	scr, err := tcell.NewScreen()

	if err != nil {
		return err
	}

	if err := scr.Init(); err != nil {
		return err
	}

	scr.EnableMouse()
	scr.Clear()

	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	scr.SetContent(0, 0, 't', nil, defStyle)
	scr.SetContent(1, 0, 'e', nil, defStyle)
	scr.SetContent(2, 0, 's', nil, defStyle)
	scr.SetContent(3, 0, 't', nil, defStyle)

	quit := func() {
		maybePanic := recover()
		scr.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}
	defer quit()

	for {
		scr.Show()

		ev := scr.PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventResize:
			scr.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape {
				return nil
			}
		}
	}
}
