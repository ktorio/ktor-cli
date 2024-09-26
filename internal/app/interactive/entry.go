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

	plugs := make(map[string]state.Plugin, len(plugins))
	sort := make([]string, 0, len(plugins))
	for _, p := range plugins {
		sort = append(sort, p.Id)
		plugs[p.Id] = state.Plugin{
			Name:        p.Name,
			Description: p.Description,
			Group:       p.Group,
		}
	}
	st.Plugins = state.Plugins{Map: plugs, Sort: sort}

	if len(plugins) > 0 {
		st.SelectedPlugin = plugins[0].Id
	}

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

		if key == tcell.KeyDown {
			if id, err := state.FindPlugin(st, func(i int) int { return i + 1 }); err == nil {
				st.SelectedPlugin = id
			}
		}

		if key == tcell.KeyUp {
			if id, err := state.FindPlugin(st, func(i int) int { return i - 1 }); err == nil {
				st.SelectedPlugin = id
			}
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
	secTextColor := tcell.Color243
	activeColor := tcell.Color63
	//activeSecColor := tcell.Color56

	padding := 3
	startX, startY := 0+padding, 0+padding
	x, y := startX, startY

	for _, id := range st.Plugins.Sort {
		color := textColor
		if id == st.SelectedPlugin {
			color = activeColor
			surface.DrawText(st.View, `⇒`, x-2, y, tcell.StyleDefault.Foreground(color))
		}

		p := st.Plugins.Map[id]
		surface.DrawText(st.View, p.Name, x, y, tcell.StyleDefault.Foreground(color))

		y++
		surface.DrawText(st.View, p.Description, x+2, y, tcell.StyleDefault.Foreground(secTextColor))
		y += 2
	}

	surface.DrawStatus(st.Status)
}

func updateState(st *state.State) {

}
