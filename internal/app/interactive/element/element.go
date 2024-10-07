package element

import "github.com/gdamore/tcell/v2"

type Type int

const (
	BoxType Type = iota
	LabelType
	InputType
)

//type BlockElement interface {
//	Children() []Element
//}

type Box struct {
	Margin
	Padding
	Children    []interface{}
	Subscribers []func(ev tcell.Event, el interface{})
}

type Label struct {
	Margin
	Padding
	Text        string
	TextStyle   tcell.Style
	Subscribers []func(ev tcell.Event, el interface{})
}

type Input struct {
	Margin
	Padding
	Value       string
	TextStyle   tcell.Style
	CursorOff   int
	Subscribers []func(ev tcell.Event, el interface{})
}

func (i Input) DefaultBehavior() Input {
	i.Subscribers = append(i.Subscribers, func(ev tcell.Event, el interface{}) {
		switch event := ev.(type) {
		case *tcell.EventKey:
			if event.Key() == tcell.KeyRune {
				i.Value += string(event.Rune())
				i.CursorOff++
			}
		}
	})
	return i
}

func (Input) Type() Type {
	return InputType
}

func (Label) Type() Type {
	return LabelType
}

func (Box) Type() Type {
	return BoxType
}

type Padding struct {
	Top, Right, Bottom, Left int
}
type Margin struct {
	Top, Right, Bottom, Left int
}

func PaddingAll(v int) Padding {
	return Padding{Top: v, Right: v, Bottom: v, Left: v}
}

func MarginRight(v int) Margin {
	return Margin{Right: v}
}
