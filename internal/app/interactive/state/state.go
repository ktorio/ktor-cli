package state

import (
	"errors"
	"fmt"
)

type View struct {
	X, Y, Width, Height int
}

func New(tWidth, tHeight int, view View) *State {
	return &State{Running: true, View: view, TWidth: tWidth, THeight: tHeight}
}

type Plugin struct {
	Name        string
	Description string
	Group       string
}

type Plugins struct {
	Map  map[string]Plugin
	Sort []string
}

type State struct {
	Running         bool
	Status          string
	View            View
	TWidth, THeight int
	Plugins         Plugins
	SelectedPlugin  string
}

type pluginIndex func(int) int

func FindPlugin(st *State, getPluginIndex pluginIndex) (string, error) {
	index := -1
	for i, id := range st.Plugins.Sort {
		if id == st.SelectedPlugin {
			index = getPluginIndex(i)
			if index >= len(st.Plugins.Sort) {
				index = i
			}

			if index <= 0 {
				index = 0
			}
		}
	}

	if index != -1 {
		return st.Plugins.Sort[index], nil
	}

	return "", errors.New(fmt.Sprintf("selected plugin %s not found", st.SelectedPlugin))
}
