package gradle

import (
	"github.com/ktorio/ktor-cli/internal/app/ktor"
	"strings"
)

func FindKtorPlugin(plugins []Plugin) (*Plugin, bool) {
	for _, p := range plugins {
		if p.Prefix == "id" && p.Id == "io.ktor.plugin" {
			return &p, true
		}
	}

	return nil, false
}

func FindKotlinPlugin(plugins []Plugin) (*Plugin, bool) {
	for _, p := range plugins {
		if p.Prefix == "kotlin" && p.Id == "jvm" {
			return &p, true
			//indent := lang.HiddenTokensToLeft(build.Stream, p.Statement.GetStart().GetTokenIndex())
			//code := fmt.Sprintf("kotlin(\"%s\") version \"%s\"", ktor.SerPluginKotlinId, p.Version)
			//build.Rewriter.InsertAfterDefault(p.Statement.GetStop().GetTokenIndex(), "\n"+indent+code)
			//break
		}
	}

	return nil, false
}

func HasSerializationPlugin(plugins []Plugin) bool {
	for _, p := range plugins {
		if (p.Prefix == "kotlin" && p.Id == ktor.SerPluginKotlinId) || (p.Prefix == "id" && p.Id == ktor.SerPluginId) {
			return true
		}
	}

	return false
}

func FindKtorDep(deps []Dep, preferTest bool) (*Dep, bool) {
	var lastKtorDep *Dep
	for _, dep := range deps {
		if dep.Kind == HardcodedDep {
			mc, ok := ktor.ParseMavenCoords(dep.Path)

			if !ok {
				continue
			}

			if mc.Group == "io.ktor" {
				lastKtorDep = &dep

				if !preferTest {
					return &dep, true
				}
			}

			if preferTest && dep.IsTest {
				return &dep, true
			}
		}

		if dep.Kind == VersionCatalogDep {
			if strings.HasPrefix(dep.Path, "libs.ktor") {
				return &dep, true
			}
		}
	}

	if lastKtorDep != nil {
		return lastKtorDep, true
	}

	return nil, false
}

func FindCatalogDep(build *BuildRoot, catalogKey string) (*Dep, bool) {
	for _, dep := range build.Dependencies.List {
		if dep.Kind == VersionCatalogDep && dep.Path == "libs."+strings.ReplaceAll(catalogKey, "-", ".") {
			return &dep, true
		}
	}

	return nil, false
}

func FindCatalogDepPrefixed(build *BuildRoot, prefix string) (*Dep, bool) {
	for _, dep := range build.Dependencies.List {
		if dep.Kind == VersionCatalogDep && strings.HasPrefix(dep.Path, prefix) {
			return &dep, true
		}
	}

	return nil, false
}