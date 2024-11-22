package command

import (
	"fmt"
	"github.com/hexops/gotextdiff"
	"github.com/hexops/gotextdiff/myers"
	"github.com/hexops/gotextdiff/span"
	"github.com/ktorio/ktor-cli/internal/app/ktor"
	"github.com/ktorio/ktor-cli/internal/app/lang/kotlin"
	"github.com/ktorio/ktor-cli/internal/app/lang/toml"
	"os"
	"path/filepath"
)

type ErrorKind int

const (
	VersionsFileAbsent ErrorKind = iota
)

type AddModuleError struct {
	Kind ErrorKind
	Err  error
}

type FileContent struct {
	Path    string
	Content string
}

func (e AddModuleError) Error() string {
	return fmt.Sprintf("cannot add module. error #%d: %v", e.Kind, e.Err)
}

func Add(mc ktor.MavenCoords, projectDir string) error {
	files, err := addDependency(mc, projectDir)

	if err != nil {
		return err
	}

	for _, f := range files {
		fmt.Println(getDiff(f.Path, f.Content))
	}

	return nil
}

func addDependency(mc ktor.MavenCoords, projectDir string) ([]FileContent, error) {
	versionsPath := filepath.Join(projectDir, "gradle", "libs.versions.toml")
	buildPath := filepath.Join(projectDir, "build.gradle.kts")
	var changes []FileContent

	buildParser, err := kotlin.NewParser(buildPath)

	if err != nil {
		return nil, err
	}

	if bom, ok := kotlin.FindBom(buildParser); ok {
		changes = append(changes, FileContent{Path: buildPath, Content: kotlin.AddRawDepAfter(buildParser, bom, mc)})
		return changes, nil
	}

	versionsParser, err := toml.NewParser(versionsPath)

	if err != nil {
		return changes, err
	}

	key, ok := toml.FindCatalogLib(versionsParser, mc)

	buildParser, err = kotlin.NewParser(buildPath)

	if err != nil {
		return changes, err
	}

	if ok {
		ok = kotlin.FindCatalogDep(buildParser, key)

		if ok {
			return changes, nil
		}
	}

	versionsParser, _ = toml.NewParser(versionsPath)
	modified, err := toml.AddLib(versionsParser, mc)

	if err != nil {
		return changes, err
	}

	changes = append(changes, FileContent{Path: versionsPath, Content: modified})

	buildParser, _ = kotlin.NewParser(buildPath)
	buildParser.GetTokenStream().Reset()
	modified, err = kotlin.AddDependency(buildParser, mc.Artifact)

	if err != nil {
		return changes, err
	}

	changes = append(changes, FileContent{Path: buildPath, Content: modified})

	return changes, nil
}

func getDiff(fp string, new string) string {
	old, err := os.ReadFile(fp)

	if err != nil {
		return ""
	}

	edits := myers.ComputeEdits(span.URIFromPath(fp), string(old), new)
	return fmt.Sprint(gotextdiff.ToUnified(filepath.Base(fp), filepath.Base(fp)+"~new", string(old), edits))
}
