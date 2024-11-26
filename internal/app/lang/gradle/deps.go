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

func HasPlugin(plugins []Plugin, pluginId string) bool {
	for _, p := range plugins {
		if p.IsKotlin && p.KotlinId == pluginId {
			return true
		}
	}

	return false
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
