package command

import (
	"errors"
	"fmt"
	"github.com/ktorio/ktor-cli/internal/app/ktor"
	"io/fs"
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

		if versionBytes, err := os.ReadFile(filepath.Join(projDir, "ktor-version.txt")); err == nil {
			ktorVersion := strings.TrimSpace(string(versionBytes))
			actualVersion, ok := SearchKtorVersion(projDir)

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

		files, result, err := addDependency(mc, projDir, serPlugin)

		if err != nil {
			t.Fatal(err)
		}

		switch e.Name() {
		case "multi-platform-catalog-projects-not-supported", "multi-platform-projects-not-supported":
			if result != MultiplatformProjectNotSupported {
				t.Fatalf("%s: Expected Multiplatform project error, got %v", e.Name(), result)
			}
		case "maven-projects-not-supported":
			if result != MavenProjectNotSupported {
				t.Fatalf("%s: Expected Maven project error, got %v", e.Name(), result)
			}
		case "groovy-dsl-projects-not-supported":
			if result != GroovyDslNotSupported {
				t.Fatalf("%s: Expected Groovy DSL project error, got %v", e.Name(), result)
			}
		case "empty-project":
			if result != BuildGradleKtsNotFound {
				t.Fatalf("%s: Expected Build Gradle KTS project error, got %v", e.Name(), result)
			}
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

			if fc == nil {
				return errors.New(fmt.Sprintf("%s: content for file %s not found", e.Name(), filepath.Base(srcPath)))
			}

			if string(expBytes) != fc.Content {
				rel, err := filepath.Rel(filepath.Dir(projDir), srcPath)

				if err != nil {
					return err
				}

				t.Fatalf("File %s has unexpected content:\n%s", rel, getDiff(p, fc.Content))
			}

			return nil
		})

		if err != nil {
			t.Fatal(err)
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
