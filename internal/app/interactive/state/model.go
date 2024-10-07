package state

import (
	"errors"
	"fmt"
)

type View struct {
	X, Y, Width, Height int
}

func New() *Model {
	return &Model{Running: true}
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

type Model struct {
	Running        bool
	WorkDir        string
	Input          string
	ProjectName    string
	Location       string
	Status         string
	Plugins        Plugins
	SelectedPlugin string
}

type pluginIndex func(int) int

func FindRelatedToSelected(st *Model, getPluginIndex pluginIndex) (string, error) {
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
