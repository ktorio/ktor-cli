package project

import (
	"fmt"
	"github.com/ktorio/ktor-cli/internal/app/lang/gradle"
	"github.com/ktorio/ktor-cli/internal/app/lang/toml"
	"github.com/ktorio/ktor-cli/internal/app/utils"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"unicode"
)

type semver struct {
	valid               bool
	major, minor, patch int
}

func (v *semver) String() string {
	return fmt.Sprintf("%d.%d.%d", v.major, v.minor, v.patch)
}

var DevModeSincePluginVersion = &semver{major: 3, minor: 1, patch: 1}

type searchResult struct {
	found   bool
	dirName string
	version semver
}

func GuessGradleTasks(projectDir string) (runTask, buildTask string, guessed bool) {
	buildPath := filepath.Join(projectDir, "build.gradle.kts")
	if !utils.Exists(buildPath) {
		return
	}

	if pVersion, found := searchKtorPlugin(projectDir); found && pVersion.valid && pVersion.aboveOrEqual(DevModeSincePluginVersion) {
		runTask = "run"
		buildTask = "classes"
		guessed = true
	} else {
		entries, err := os.ReadDir(projectDir)
		if err != nil {
			return
		}

		var candidateDirs []string
		for _, e := range entries {
			if e.IsDir() && !strings.HasPrefix(e.Name(), ".") && e.Name() != "gradle" {
				if utils.Exists(filepath.Join(projectDir, e.Name(), "build.gradle.kts")) {
					candidateDirs = append(candidateDirs, e.Name())
				}
			}
		}

		if len(candidateDirs) > 0 {
			searchChan := make(chan searchResult)

			for _, dirName := range candidateDirs {
				candidateDir := filepath.Join(projectDir, dirName)
				go func(dirPath, dirName string) {
					pVersion, found = searchKtorPlugin(dirPath)
					searchChan <- searchResult{found: found, version: pVersion, dirName: dirName}
				}(candidateDir, dirName)
			}

			numResults := 0
			running := true
			for running {
				result := <-searchChan
				numResults++

				if result.found && pVersion.valid && pVersion.aboveOrEqual(DevModeSincePluginVersion) {
					runTask = fmt.Sprintf(":%s:run", result.dirName)
					buildTask = fmt.Sprintf(":%s:classes", result.dirName)
					guessed = true
					running = false
				} else if numResults == len(candidateDirs) {
					running = false
				}
			}
		}
	}

	return
}
func searchKtorPlugin(dir string) (version semver, found bool) {
	buildPath := filepath.Join(dir, "build.gradle.kts")
	buildRoot, err, _ := gradle.ParseBuildFile(buildPath)

	if err != nil {
		return
	}

	found = true
	var catalogKeys []string
	for _, p := range buildRoot.Plugins.List {
		if p.IsCatalog && p.Applied {
			if segments := strings.Split(p.Id, "."); len(segments) > 0 {
				off := len(segments) - 1

				for ; off >= 0; off-- {
					if segments[off] == "plugins" {
						off++
						break
					}
				}

				catalogKeys = append(catalogKeys, strings.Join(segments[off:], "-"))
			}
		}

		if p.Prefix == "id" && p.Id == "io.ktor.plugin" && p.Applied {
			if p.Version != "" {
				version = parseVersion(p.Version)
			}
			return
		}
	}

	if len(catalogKeys) > 0 {
		hasCatalog := false
		catalogName := "libs.versions.toml"
		catalogPath := filepath.Join(dir, "gradle", catalogName)
		if utils.Exists(catalogPath) {
			hasCatalog = true
		} else {
			catalogPath = filepath.Join(dir, "..", "gradle", catalogName)

			if utils.Exists(catalogPath) {
				hasCatalog = true
			}
		}

		if hasCatalog {
			tomlDoc, err, _ := toml.ParseCatalogToml(catalogPath)

			if err != nil {
				found = false
				return
			}

			if t, ok := toml.FindTable(tomlDoc, "plugins"); ok {
				for _, te := range t.Entries {
					if te.Kind == toml.ValueMap {
						if id, ok := te.Get("id"); ok && id == "io.ktor.plugin" {

							pluginApplied := false
							for _, catalogKey := range catalogKeys {
								if te.Key == catalogKey {
									pluginApplied = true
									break
								}
							}

							if pluginApplied {
								found = true

								if versionRef, ok := te.Get("version.ref"); ok {
									if vTable, ok := toml.FindTable(tomlDoc, "versions"); ok {
										for _, vt := range vTable.Entries {
											if vt.Key == versionRef && vt.Kind == toml.StringValue {
												version = parseVersion(vt.String)
												break
											}
										}
									}
								}

								return
							}
						}
					}
				}
			}
		}
	}

	found = false
	return
}

func parseVersion(str string) (version semver) {
	comps := strings.Split(str, ".")
	if len(comps) == 3 {
		majorComp, minorComp, patchComp := comps[0], comps[1], comps[2]

		major, err := strconv.Atoi(majorComp)
		if err != nil {
			return
		}

		minor, err := strconv.Atoi(minorComp)
		if err != nil {
			return
		}

		patch := 0
		for _, c := range patchComp {
			if unicode.IsNumber(c) { // May be beta or eap
				patch += patch*10 + (int(c) - '0')
			} else {
				break
			}
		}

		version = semver{major: major, minor: minor, patch: patch, valid: true}
	}

	return
}

func (v *semver) aboveOrEqual(other *semver) bool {
	if v.major > other.major {
		return true
	} else if v.major == other.major {
		if v.minor > other.minor {
			return true
		} else if v.minor == other.minor {
			if v.patch >= other.patch {
				return true
			} else {
				return false
			}
		} else {
			return false
		}
	} else {
		return false
	}
}
