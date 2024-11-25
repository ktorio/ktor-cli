package command

import (
	"errors"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/hexops/gotextdiff"
	"github.com/hexops/gotextdiff/myers"
	"github.com/hexops/gotextdiff/span"
	"github.com/ktorio/ktor-cli/internal/app/ktor"
	"github.com/ktorio/ktor-cli/internal/app/lang/gradle"
	"github.com/ktorio/ktor-cli/internal/app/lang/kotlin"
	parser "github.com/ktorio/ktor-cli/internal/app/lang/parsers/kotlin"
	"github.com/ktorio/ktor-cli/internal/app/lang/toml"
	"os"
	"path/filepath"
	"strings"
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

	build, err := gradle.ParseBuildFile(buildPath)
	if err != nil {
		return changes, nil
	}

	buildParser, err := kotlin.NewParser(buildPath)

	if err != nil {
		return nil, err
	}

	if bom, ok := gradle.FindBom(build); ok {
		if sts, ok := bom.GetParent().(parser.IStatementsContext); ok && kotlin.FindRawDep(sts, mc) {
			return changes, nil
		}

		changes = append(changes, FileContent{Path: buildPath, Content: kotlin.AddRawDepAfter(buildParser, bom, mc)})
		return changes, nil
	}

	hasKtorDeps := false
	for _, dep := range build.Dependencies.List {
		switch dep.Kind {
		case gradle.VersionCatalogDep:
			if strings.HasPrefix(dep.CatalogPath, "libs.ktor") {
				hasKtorDeps = true
				break
			}
		}
	}

	if _, err := os.Stat(versionsPath); errors.Is(err, os.ErrNotExist) && !hasKtorDeps {
		tomlContent := fmt.Sprintf(`[versions]
ktor = "%s"

[libraries]
%s = { module = "%s:%s", version.ref = "ktor" }
`, mc.Version, mc.Artifact, mc.Group, mc.Artifact)

		changes = append(changes, FileContent{Path: versionsPath, Content: tomlContent})

		rewriter := antlr.NewTokenStreamRewriter(build.Stream)
		indent := strings.Repeat(" ", 4)
		rewriter.InsertAfterDefault(build.Dependencies.Statements.GetStop().GetTokenIndex(), "\n"+indent+fmt.Sprintf("implementation(libs.%s)\n", strings.ReplaceAll(mc.Artifact, "-", ".")))

		changes = append(changes, FileContent{Path: buildPath, Content: rewriter.GetTextDefault()})

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
