package project

import (
	"fmt"
	"github.com/ktorio/ktor-cli/internal/app/ktor"
	"github.com/ktorio/ktor-cli/internal/app/lang/gradle"
	"github.com/ktorio/ktor-cli/internal/app/lang/toml"
	"path/filepath"
	"strings"
)

func SearchKtorVersion(projectDir string, build *gradle.BuildRoot, tomlDoc *toml.Document, tomlSuccessParsed bool) (version string, found bool) {
	found = true

	for _, p := range build.Plugins.List {
		if p.Prefix == "id" && p.Id == "io.ktor.plugin" {
			version = p.Version
			return
		}
	}

	for _, d := range build.Dependencies.List {
		if d.IsKtorBom {
			if mc, ok := ktor.ParseMavenCoords(d.PlatformPath); ok && mc.Version != "" {
				version = mc.Version
				break
			}
		} else if d.Kind == gradle.HardcodedDep {
			if mc, ok := ktor.ParseMavenCoords(d.Path); ok && mc.Group == ktor.MavenGroup && mc.Version != "" {
				version = mc.Version
				break
			}
		}
	}

	if version != "" && !strings.HasPrefix(version, "$") {
		return
	} else {
		props := gradle.ParseProps(filepath.Join(projectDir, "gradle.properties"))

		for _, v := range build.TopLevelVars {
			if v.Id == strings.TrimPrefix(version, "$") {
				if v.IsDelegate && v.Delegate == "project" {
					if val, ok := props[strings.TrimPrefix(version, "$")]; ok {
						version = val
						break
					}
				} else {
					version = v.StringVal
					break
				}
			}
		}

		if version != "" {
			return
		}
	}

	if tomlSuccessParsed {
		if t, ok := toml.FindTable(tomlDoc, "versions"); ok {
			for _, te := range t.Entries {
				if strings.HasPrefix(te.Key, "ktor") && te.Kind == toml.StringValue && te.String != "" {
					version = te.String
					return
				}
			}
		}
	}

	found = false
	return
}

func IsKmp(build *gradle.BuildRoot, tomlDoc *toml.Document, validToml bool) bool {
	for _, p := range build.Plugins.List {
		if (p.Prefix == "kotlin" && p.Id == "multiplatform") || (p.Prefix == "id" && p.Id == "org.jetbrains.kotlin.multiplatform") {
			return true
		}
	}

	if validToml {
		plugin, ok := toml.FindPlugin(tomlDoc, ktor.KmpPluginId)

		if !ok {
			return false
		}

		for _, p := range build.Plugins.List {
			if p.Prefix == "kotlin" && p.Id == "multiplatform" {
				return true
			}

			if p.Prefix == "alias" && p.Id == fmt.Sprintf("libs.plugins.%s", plugin.Key) {
				return true
			}
		}
	}

	return false
}
