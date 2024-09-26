package state

import "github.com/ktorio/ktor-cli/internal/app/network"

type View struct {
	X, Y, Width, Height int
}

func New(tWidth, tHeight int, view View) *State {
	return &State{Running: true, View: view, TWidth: tWidth, THeight: tHeight}
}

type State struct {
	Running         bool
	Status          string
	View            View
	TWidth, THeight int
	Plugins         []network.Plugin
}
