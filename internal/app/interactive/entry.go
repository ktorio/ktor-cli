package interactive

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/ktorio/ktor-cli/internal/app/interactive/draw"
	"github.com/ktorio/ktor-cli/internal/app/interactive/element"
	"github.com/ktorio/ktor-cli/internal/app/interactive/state"
	"github.com/ktorio/ktor-cli/internal/app/network"
	"net/http"
	"os"
	"path/filepath"
	"time"
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

	//width, height := scr.Size()
	//surface := draw.NewScreen(scr)
	model := state.New()

	settings, err := network.FetchSettings(client)

	if err != nil {
		// TODO: handle error
		return err
	}

	model.ProjectName = settings.ProjectName.DefaultVal
	model.Input = model.ProjectName

	wd, err := os.Getwd()

	if err != nil {
		// TODO: handle error
		return err
	}

	model.WorkDir = wd
	model.Location = filepath.Join(model.WorkDir, model.ProjectName)

	//ktorVersion := settings.KtorVersion.DefaultId
	//
	//plugins, err := network.FetchPlugins(client, ktorVersion)
	//
	//if err != nil {
	//	// TODO: handle error
	//	return err
	//}
	//
	//plugs := make(map[string]state.Plugin, len(plugins))
	//sort := make([]string, 0, len(plugins))
	//for _, p := range plugins {
	//	sort = append(sort, p.Id)
	//	plugs[p.Id] = state.Plugin{
	//		Name:        p.Name,
	//		Description: p.Description,
	//		Group:       p.Group,
	//	}
	//}
	//model.Plugins = state.Plugins{Map: plugs, Sort: sort}
	//
	//if len(plugins) > 0 {
	//	model.SelectedPlugin = plugins[0].Id
	//}

	eventChan := make(chan tcell.Event)
	startTime := time.Now().UnixMicro()
	frameStart := startTime
	frameMs := 1000.0 / 60
	frame := 0

	go func() {
		for model.Running {
			eventChan <- scr.PollEvent()
		}
	}()

	captionStyle := tcell.StyleDefault.Foreground(tcell.ColorGreen)
	var root interface{} = element.Box{Padding: element.PaddingAll(2), Children: []interface{}{
		element.Label{Text: "Enter project name:", Margin: element.MarginRight(1), TextStyle: captionStyle},
		element.Input{Value: ""}.DefaultBehavior(),
	}}

	drawState := draw.NewState(scr)
	drawState.Padding = 2

	for model.Running {
		delta := float64(time.Now().UnixMicro()-startTime) / 1000.0
		startTime = time.Now().UnixMicro()

		if startTime-frameStart > 1e6 {
			frame = 0
			frameStart = startTime
		}

		select {
		case event := <-eventChan:
			processInput(event, model, root)
		default:
			// do nothing
		}

		updateModel(model)

		scr.Clear()
		drawTui(root, scr, model, frame)
		scr.Show()

		frame++

		if frameMs-delta > 0 {
			time.Sleep(time.Duration(frameMs-delta) * time.Millisecond)
		}
	}

	return nil
}

func processInput(ev tcell.Event, model *state.Model, root interface{}) {
	//scrollSpeed := 3

	dispatchEvent(ev, root)

	switch ev := ev.(type) {
	//case *tcell.EventResize:
	//	model.TWidth, model.THeight = ev.Size()
	//	model.View.Width, model.View.Height = model.TWidth, model.THeight
	case *tcell.EventKey:
		mod, key := ev.Modifiers(), ev.Key()

		switch {
		case (mod == tcell.ModCtrl && key == tcell.KeyCtrlC) || (key == tcell.KeyEscape):
			model.Running = false
		case key == tcell.KeyRune:
			model.Input += string(ev.Rune())
		case key == tcell.KeyDEL:
			if len(model.Input) > 0 {
				model.Input = model.Input[:len(model.Input)-1]
			}
		case key == tcell.KeyDown:
			if id, err := state.FindRelatedToSelected(model, func(i int) int { return i + 1 }); err == nil {
				model.SelectedPlugin = id
			}
		case key == tcell.KeyUp:
			if id, err := state.FindRelatedToSelected(model, func(i int) int { return i - 1 }); err == nil {
				model.SelectedPlugin = id
			}
			//case key == tcell.KeyPgDn:
			//	model.View.Y += scrollSpeed
			//case key == tcell.KeyPgUp:
			//	model.View.Y -= scrollSpeed
			//	if model.View.Y <= 0 {
			//		model.View.Y = 0
			//	}
		}
	case *tcell.EventMouse:
		mod := ev.Modifiers()
		btns := ev.Buttons()
		x, y := ev.Position()

		model.Status = fmt.Sprintf("EventMouse Modifiers: %d Buttons: %d Position: %d,%d", mod, btns, x, y)

		//if btns == tcell.WheelDown {
		//	model.View.Y += scrollSpeed
		//}
		//
		//if btns == tcell.WheelUp {
		//	model.View.Y -= scrollSpeed
		//	if model.View.Y <= 0 {
		//		model.View.Y = 0
		//	}
		//}
	}
}

func dispatchEvent(ev tcell.Event, root interface{}) {
	switch el := root.(type) {
	case element.Label:
		for _, sub := range el.Subscribers {
			sub(ev, el)
		}
	case element.Input:
		for _, sub := range el.Subscribers {
			sub(ev, el)
		}
	case element.Box:
		for _, ch := range el.Children {
			dispatchEvent(ev, ch)
		}

		for _, sub := range el.Subscribers {
			sub(ev, el)
		}
	}

	//if el, ok := root.(interface{ Children() []interface{} }); ok {
	//	for _, e := range el.Children() {
	//		if fff, ok := e.(interface{ Subscribers() []interface{} }); ok {
	//
	//		}
	//	}
	//} else {
	//	root.TriggerEvent(ev)
	//}

	//switch e := ev.(type) {
	//case *tcell.EventKey:
	//	if e.Key() == tcell.KeyRune {
	//
	//	}
	//}
}

func drawTui(el interface{}, scr tcell.Screen, model *state.Model, frame int) {
	DrawElement(el, scr, 0, 0, frame)

	//textColor := tcell.ColorWhite
	//secTextColor := tcell.Color243
	//activeColor := tcell.Color63
	////activeSecColor := tcell.Color56

	//padding := 2
	//dr.Move(2, 2)
	//captionStyle := tcell.StyleDefault.Foreground(tcell.ColorGreen)
	//
	//dr.Text("Enter project name: ", captionStyle)
	//
	////if frame <= 30 {
	////	draw.Text(scr, "_", curPosX+len(model.Input), curPosY, tcell.StyleDefault)
	////} else if frame <= 60 {
	////	draw.Text(scr, " ", curPosX+len(model.Input), curPosY, tcell.StyleDefault)
	////}
	//
	//dr.Text(model.Input, tcell.StyleDefault)
	//dr.ResetPos()
	//dr.MoveY(2)
	//
	//dr.Text("Location: ", captionStyle)
	//
	//curPosX, curPosY = draw.Text(scr, "Location: ", 2, curPosY+2, tcell.StyleDefault.Foreground(tcell.ColorGreen))
	//draw.Text(scr, model.Location, curPosX, curPosY, tcell.StyleDefault)

	//go func() {
	//	draw.Text(scr, " ", curPosX, curPosY, tcell.StyleDefault)
	//	time
	//	draw.Text(scr, "|", curPosX, curPosY, tcell.StyleDefault)
	//}()

	//
	//padding := 3
	//startX, startY := 0+padding, 0+padding
	//x, y := startX, startY
	//
	//for _, id := range st.Plugins.Sort {
	//	color := textColor
	//	if id == st.SelectedPlugin {
	//		color = activeColor
	//		surface.Text(st, `⇒`, x-2, y, tcell.StyleDefault.Foreground(color))
	//	}
	//
	//	p := st.Plugins.Map[id]
	//	surface.Text(st, p.Name, x, y, tcell.StyleDefault.Foreground(color))
	//
	//	y++
	//	surface.Text(st, p.Description, x+2, y, tcell.StyleDefault.Foreground(secTextColor))
	//	y += 2
	//}

	//surface.DrawStatus(st.Status)
}

func DrawElement(el interface{}, scr tcell.Screen, x, y int, frame int) (int, int) {
	switch e := el.(type) {
	case element.Label:
		pad := e.Padding
		marg := e.Margin
		x, y = draw.Text(scr, e.Text, x+pad.Left+marg.Left, y+pad.Top+marg.Top, e.TextStyle)
		x += marg.Right + pad.Right
		return x, y
	case element.Input:
		activeCurStyle := tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorWhite)
		for i, r := range []rune(e.Value) {
			style := e.TextStyle
			if i == e.CursorOff && frame < 60 {
				style = activeCurStyle
			}

			scr.SetContent(x, y, r, nil, style)
			x++
		}

		if e.Value == "" && frame < 60 {
			scr.SetContent(x, y, ' ', nil, activeCurStyle)
		}

		return x, y
	case element.Box:
		pad := e.Padding
		marg := e.Margin
		cx, cy := x+pad.Left+marg.Left, y+pad.Top+marg.Top

		for _, ch := range e.Children {
			cx, cy = DrawElement(ch, scr, cx, cy, frame)
		}
	}

	//pad := el.Padding()
	//cx, cy := x+pad.Left, y+pad.Top
	//
	//switch e := el.(type) {
	//case element.BlockElement:
	//	for _, ch := range e.Children() {
	//		sw
	//	}
	//}

	return x, y
}

func updateModel(m *state.Model) {
	m.ProjectName = m.Input
	m.Location = filepath.Join(m.WorkDir, m.ProjectName)
}
