package gradle

import (
	"github.com/ktorio/ktor-cli/internal/app/ktor"
	"strings"
)

func FindCatalogDep(deps []Dep, catalogKey string) bool {
	for _, dep := range deps {
		if dep.Kind == VersionCatalogDep && dep.Path == "libs."+strings.ReplaceAll(catalogKey, "-", ".") {
			return true
		}
	}

	return false
}

func FindKtorPlugin(plugins []Plugin) (*Plugin, bool) {
	for _, p := range plugins {
		if p.Prefix == "id" && p.Id == "io.ktor.plugin" {
			return &p, true
		}
	}

	return nil, false
}

func HasPlugin(plugins []Plugin, pluginId string) bool {
	for _, p := range plugins {
		if p.Prefix == "kotlin" && p.Id == pluginId {
			return true
		}
	}

	return false
}

func FindKtorDep(deps []Dep) (*Dep, bool) {
	for _, dep := range deps {
		if dep.Kind == HardcodedRep {
			mc, ok := ktor.ParseMavenCoords(dep.Path)

			if !ok {
				continue
			}

			if mc.Group == "io.ktor" {
				return &dep, true
			}
		}

		if dep.Kind == VersionCatalogDep {
			if strings.HasPrefix(dep.Path, "libs.ktor") {
				return &dep, true
			}
		}
	}

	return nil, false
}

func FindRawDep(deps []Dep, mavenCoords ktor.MavenCoords) bool {
	for _, dep := range deps {
		if dep.Kind != HardcodedRep {
			continue
		}

		mc, ok := ktor.ParseMavenCoords(dep.Path)

		if !ok {
			continue
		}

		if mavenCoords.RoughlySame(mc) {
			return true
		}
	}

	return false
}
