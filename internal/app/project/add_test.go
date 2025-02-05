package project

import (
	"errors"
	"github.com/ktorio/ktor-cli/internal/app/ktor"
	"github.com/ktorio/ktor-cli/internal/app/lang"
	"github.com/ktorio/ktor-cli/internal/app/lang/gradle"
	"github.com/ktorio/ktor-cli/internal/app/lang/toml"
	"github.com/ktorio/ktor-cli/internal/app/utils"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"
)

func TestAddProjectDependencies(t *testing.T) {
	testDir := filepath.Join("internal", "app", "cli", "command", "testData")

	if _, err := os.Stat(testDir); errors.Is(err, os.ErrNotExist) {
		testDir = "testData"
	}

	entries, err := os.ReadDir(testDir)

	if err != nil {
		t.Fatal(err)
	}

	for _, e := range entries {
		if !e.IsDir() {
			continue
		}

		projDir := filepath.Join(testDir, e.Name())

		b, err := os.ReadFile(filepath.Join(projDir, "ktor-module.txt"))
		if err != nil {
			t.Fatalf("Expected ktor-module.txt file in the %s", projDir)
		}

		buildPath := filepath.Join(projDir, "build.gradle.kts")
		buildRoot, buildErr, buildSyntaxErrors := gradle.ParseBuildFile(buildPath)

		tomlPath, tomlFound := toml.FindVersionsPath(projDir)
		tomlSuccessParsed := false
		var tomlDoc *toml.Document

		if tomlFound {
			tomlDoc, err, _ = toml.ParseCatalogToml(tomlPath)

			if err == nil {
				tomlSuccessParsed = true
			}
		}

		if e.Name() == "multi-platform-catalog-projects-not-supported" || e.Name() == "multi-platform-projects-not-supported" {
			if IsKmp(buildRoot, tomlDoc, tomlSuccessParsed) {
				continue
			} else {
				log.Fatalf("%s: expected multiplatform project to be unsupported", e.Name())
			}
		}

		if versionBytes, err := os.ReadFile(filepath.Join(projDir, "ktor-version.txt")); err == nil && buildErr == nil && utils.Exists(buildPath) {
			ktorVersion := strings.TrimSpace(string(versionBytes))
			actualVersion, ok := SearchKtorVersion(projDir, buildRoot, tomlDoc, tomlSuccessParsed)

			if len(ktorVersion) != 0 && !ok {
				t.Fatalf("%s: expected Ktor version to be %s, found nothing", e.Name(), ktorVersion)
			}

			if actualVersion != ktorVersion {
				t.Fatalf("%s: expected Ktor version to be %s, got %s", e.Name(), ktorVersion, actualVersion)
			}
		}

		parts := strings.Split(strings.TrimSpace(string(b)), ":")
		version := ""
		artifact := parts[0]
		if len(parts) > 1 {
			version = parts[1]
		}

		mc := ktor.MavenCoords{Artifact: artifact, Group: "io.ktor", Version: version, IsTest: artifact == "ktor-server-test-host"}
		depPlugins := ktor.DependentPlugins(mc)
		var serPlugin *ktor.GradlePlugin
		if len(depPlugins) > 0 {
			serPlugin = &depPlugins[0]
		}

		if buildErr == nil && utils.Exists(buildPath) {
			files, err := AddKtorModule(mc, buildRoot, tomlDoc, tomlSuccessParsed, serPlugin, buildPath, tomlPath, projDir)

			if len(buildSyntaxErrors) > 0 && !utils.Exists(filepath.Join(projDir, "expect-error.txt")) {
				t.Fatalf("%s: unexpected syntax errors\n%s", e.Name(), lang.StringifySyntaxErrors(buildSyntaxErrors))
			}

			err = filepath.WalkDir(projDir, func(p string, d fs.DirEntry, err error) error {
				if err != nil {
					return err
				}

				if !strings.HasSuffix(p, ".expected") {
					return nil
				}

				srcPath := strings.TrimSuffix(p, ".expected")

				srcBytes, err := os.ReadFile(srcPath)

				if err != nil {
					srcBytes = []byte{}
				}

				expBytes, err := os.ReadFile(p)

				if err != nil {
					return err
				}

				fc := findFileContent(files, srcPath)
				if slices.Equal(srcBytes, expBytes) && fc == nil {
					return nil
				}

				expContent := ""
				if fc != nil {
					expContent = fc.Content
				}

				if string(expBytes) != expContent {
					rel, err := filepath.Rel(filepath.Dir(projDir), srcPath)

					if err != nil {
						return err
					}

					t.Fatalf("File %s has unexpected content:\n%s", rel, utils.GetDiff(p, expContent))
				}

				return nil
			})

			if err != nil {
				t.Fatal(err)
			}
		}
	}
}

func findFileContent(files []FileContent, fp string) *FileContent {
	for _, fc := range files {
		if fc.Path == fp {
			return &fc
		}
	}

	return nil
}
