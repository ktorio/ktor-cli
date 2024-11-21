package command

import (
	"errors"
	"github.com/ktorio/ktor-cli/internal/app/ktor"
	"os"
	"path/filepath"
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

		files, err := addDependency(ktor.MavenCoords{Artifact: strings.TrimSpace(string(b)), Group: "io.ktor"}, projDir)

		if err != nil {
			t.Fatal(err)
		}

		for _, f := range files {
			rel, err := filepath.Rel(filepath.Dir(projDir), f.Path)

			if err != nil {
				t.Fatal(err)
			}

			expectedPath := f.Path + ".expected"

			b, err := os.ReadFile(expectedPath)

			if err != nil {
				t.Fatal(err)
			}

			if string(b) != f.Content {
				t.Fatalf("File %s has unexpected content:\n%s", rel, getDiff(expectedPath, f.Content))
			}
		}
	}
}
