package command

import (
	"errors"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/hexops/gotextdiff"
	"github.com/hexops/gotextdiff/myers"
	"github.com/hexops/gotextdiff/span"
	"github.com/ktorio/ktor-cli/internal/app/ktor"
	"github.com/ktorio/ktor-cli/internal/app/lang"
	"github.com/ktorio/ktor-cli/internal/app/lang/gradle"
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

func Add(mc ktor.MavenCoords, projectDir string, serPlugin *ktor.GradlePlugin) error {
	files, err := addDependency(mc, projectDir, serPlugin)

	if err != nil {
		return err
	}

	for _, f := range files {
		fmt.Println(getDiff(f.Path, f.Content))
	}

	return nil
}

func addDependency(mc ktor.MavenCoords, projectDir string, serPlugin *ktor.GradlePlugin) ([]FileContent, error) {
	versionsPath := filepath.Join(projectDir, "gradle", "libs.versions.toml")
	buildPath := filepath.Join(projectDir, "build.gradle.kts")
	var changes []FileContent

	build, err := gradle.ParseBuildFile(buildPath)
	if err != nil {
		return changes, nil
	}

	bom, hasBom := gradle.FindBom(build.Dependencies.List)
	_, hasKtorPlugin := gradle.FindKtorPlugin(build.Plugins.List)

	if hasBom || hasKtorPlugin {
		if serPlugin != nil && !gradle.HasPlugin(build.Plugins.List, "plugin.serialization") {
			for _, p := range build.Plugins.List {
				if p.Prefix == "kotlin" && p.Id == "jvm" {
					indent := lang.HiddenTokensToLeft(build.Stream, p.Statement.GetStart().GetTokenIndex())
					code := fmt.Sprintf("kotlin(\"plugin.serialization\") version \"%s\"", p.Version)
					build.Rewriter.InsertAfterDefault(p.Statement.GetStop().GetTokenIndex(), "\n"+indent+code)
					break
				}
			}
		}

		if gradle.FindRawDep(build.Dependencies.List, mc) {
			return changes, nil
		}

		if hasBom {
			gradle.AddRawDepAfter(build, bom, mc)
		} else {
			if kDep, ok := gradle.FindKtorDep(build.Dependencies.List, mc.IsTest); ok {
				gradle.AddRawDepAfter(build, kDep.Statement, mc)
			}
		}

		changes = append(changes, FileContent{Path: buildPath, Content: build.Rewriter.GetTextDefault()})
		return changes, nil
	}

	hasKtorDeps := false
	for _, dep := range build.Dependencies.List {
		if dep.Kind == gradle.VersionCatalogDep && strings.HasPrefix(dep.Path, "libs.ktor") {
			hasKtorDeps = true
			break
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
		rewriter.InsertAfterDefault(build.Dependencies.Element.GetStop().GetTokenIndex(), "\n"+indent+fmt.Sprintf("implementation(libs.%s)\n", strings.ReplaceAll(mc.Artifact, "-", ".")))

		changes = append(changes, FileContent{Path: buildPath, Content: rewriter.GetTextDefault()})

		return changes, nil
	}

	tomlDoc, err := toml.ParseToml(versionsPath)

	if err != nil {
		return changes, err
	}

	key, ok := toml.FindCatalogLib(tomlDoc.Tables.List, mc)

	if ok {
		ok = gradle.FindCatalogDep(build.Dependencies.List, key)

		if ok {
			return changes, nil
		}
	}

	modified, err := toml.AddLib(tomlDoc, mc)

	if err != nil {
		return changes, err
	}

	changes = append(changes, FileContent{Path: versionsPath, Content: modified})

	modified, err = gradle.AddDependency(build, mc.Artifact)

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
