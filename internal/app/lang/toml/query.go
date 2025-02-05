package toml

import (
	"github.com/ktorio/ktor-cli/internal/app/ktor"
	"strings"
)

func FindLib(doc *Document, mc ktor.MavenCoords) (*TableEntry, bool) {
	libTable, ok := FindTable(doc, "libraries")

	if !ok {
		return nil, false
	}

	for _, e := range libTable.Entries {
		if e.Kind != ValueMap {
			continue
		}

		path, ok := e.KeyValue["module"]

		if !ok {
			continue
		}

		if coords, ok := ktor.ParseMavenCoords(path); ok && mc.RoughlySame(coords) {
			return &e, true
		}
	}

	return nil, false
}

func FindPlugin(doc *Document, pluginId string) (*TableEntry, bool) {
	pluginsTable, ok := FindTable(doc, "plugins")

	if !ok {
		return nil, false
	}

	for _, e := range pluginsTable.Entries {
		if id, ok := e.Get("id"); ok && id == pluginId {
			return &e, true
		}
	}

	return nil, false
}

func FindVersionPrefixed(doc *Document, prefix string) (*TableEntry, bool) {
	versionsTable, ok := FindTable(doc, "versions")

	if !ok {
		return nil, false
	}

	for _, e := range versionsTable.Entries {
		if e.Kind == StringValue && strings.HasPrefix(e.Key, prefix) {
			return &e, true
		}
	}

	return nil, false
}

func FindTable(doc *Document, name string) (*Table, bool) {
	for _, t := range doc.Tables.List {
		if t.Name == name {
			return &t, true
		}
	}

	return nil, false
}
