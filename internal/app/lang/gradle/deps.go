package gradle

import (
	"github.com/ktorio/ktor-cli/internal/app/ktor"
	"strings"
)

func FindCatalogDep(build *BuildRoot, catalogKey string) bool {
	for _, dep := range build.Dependencies.List {
		if dep.Kind == VersionCatalogDep && dep.Path == "libs."+strings.ReplaceAll(catalogKey, "-", ".") {
			return true
		}
	}

	return false
}

func FindRawDep(build *BuildRoot, mavenCoords ktor.MavenCoords) bool {
	for _, dep := range build.Dependencies.List {
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
