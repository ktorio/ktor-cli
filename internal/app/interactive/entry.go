package interactive

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/ktorio/ktor-cli/internal/app/interactive/draw"
	"github.com/ktorio/ktor-cli/internal/app/interactive/state"
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
	surface := draw.NewScreen(scr)
	st := state.New(width, height, state.View{X: 0, Y: 0, Width: width, Height: height})

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

	st.Plugins = plugins

	for st.Running {
		processInput(scr.PollEvent(), st)
		updateState(st)

		scr.Clear()
		drawTui(surface, st)
		scr.Show()
	}

	return nil
}

func processInput(ev tcell.Event, st *state.State) {
	scrollSpeed := 3

	switch ev := ev.(type) {
	case *tcell.EventResize:
		st.TWidth, st.THeight = ev.Size()
		st.View.Width, st.View.Height = st.TWidth, st.THeight
	case *tcell.EventKey:
		mod, key := ev.Modifiers(), ev.Key()
		if (mod == tcell.ModCtrl && key == tcell.KeyCtrlC) || (key == tcell.KeyEscape) {
			st.Running = false
		}

		if key == tcell.KeyPgDn {
			st.View.Y += scrollSpeed
		}

		if key == tcell.KeyPgUp {
			st.View.Y -= scrollSpeed
			if st.View.Y <= 0 {
				st.View.Y = 0
			}
		}

	case *tcell.EventMouse:
		mod := ev.Modifiers()
		btns := ev.Buttons()
		x, y := ev.Position()

		st.Status = fmt.Sprintf("EventMouse Modifiers: %d Buttons: %d Position: %d,%d", mod, btns, x, y)

		if btns == tcell.WheelDown {
			st.View.Y += scrollSpeed
		}

		if btns == tcell.WheelUp {
			st.View.Y -= scrollSpeed
			if st.View.Y <= 0 {
				st.View.Y = 0
			}
		}
	}
}

func drawTui(surface *draw.Screen, st *state.State) {
	textColor := tcell.ColorWhite
	secColor := tcell.Color243
	//mainColor := tcell.Color27

	padding := 1
	startX, startY := 0+padding, 1+padding
	x, y := startX, startY

	for _, p := range st.Plugins {
		surface.DrawText(st.View, p.Name, x, y, tcell.StyleDefault.Foreground(textColor))
		y++
		surface.DrawText(st.View, p.Description, x+2, y, tcell.StyleDefault.Foreground(secColor))
		y += 2
	}

	surface.DrawStatus(st.Status)
}

func updateState(st *state.State) {

}
