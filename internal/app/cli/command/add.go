package command

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/hexops/gotextdiff"
	"github.com/hexops/gotextdiff/myers"
	"github.com/hexops/gotextdiff/span"
	"github.com/ktorio/ktor-cli/internal/app/ktor"
	"github.com/ktorio/ktor-cli/internal/app/lang"
	"github.com/ktorio/ktor-cli/internal/app/lang/gradle"
	"github.com/ktorio/ktor-cli/internal/app/lang/toml"
	"github.com/ktorio/ktor-cli/internal/app/utils"
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

	if len(files) == 0 {
		fmt.Println("Nothing to change.")
		fmt.Println("Goodbye!")
		return nil
	}

	for _, f := range files {
		fmt.Println(getDiff(f.Path, f.Content))
	}

	fmt.Print("Do you want to apply the changes above (y/n)? ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	answer := scanner.Text()

	if answer == "y" || answer == "Y" || answer == "yes" || answer == "Yes" {
		err = applyChanges(files)

		if err == nil {
			fmt.Println("The changes are successfully applied.")
		} else {
			fmt.Println("An error occurred applying the changes.")
		}

		return err
	}

	fmt.Println("Goodbye!")

	return nil
}

func applyChanges(files []FileContent) error {
	// Load all current files content into memory
	var savedContents [][]byte
	for _, f := range files {
		b, err := os.ReadFile(f.Path)
		if err != nil {
			return err
		}

		savedContents = append(savedContents, b)
	}

	// Write changes to all files
	var lastErr error
	for _, fc := range files {
		err := os.WriteFile(fc.Path, []byte(fc.Content), 0777)

		if err != nil {
			lastErr = err
			break
		}
	}

	// If at least one error -> roll everything back
	if lastErr != nil {
		for i, b := range savedContents {
			fc := files[i]
			_ = os.WriteFile(fc.Path, b, 0777)
		}
	}

	return lastErr
}

func addDependency(mc ktor.MavenCoords, projectDir string, serPlugin *ktor.GradlePlugin) ([]FileContent, error) {
	versionsPath := filepath.Join(projectDir, "gradle", "libs.versions.toml")
	buildPath := filepath.Join(projectDir, "build.gradle.kts")
	var changes []FileContent

	if !utils.Exists(buildPath) {
		if utils.Exists(filepath.Join(projectDir, "pom.xml")) {
			fmt.Println("Adding Ktor dependency to a Maven project is not supported.")
		}

		if utils.Exists(filepath.Join(projectDir, "build.gradle")) {
			fmt.Println("Adding Ktor dependency to a Gradle project with Groovy DSL is not supported.")
		}

		fmt.Printf("Build file %s is expected but not found.\n", buildPath)

		os.Exit(1)
	}

	build, err := gradle.ParseBuildFile(buildPath)
	if err != nil {
		return changes, err
	}

	// Looking for a hardcoded dependency
	for _, dep := range build.Dependencies.List {
		if dep.Kind != gradle.HardcodedDep {
			continue
		}

		if coords, ok := ktor.ParseMavenCoords(dep.Path); ok && mc.RoughlySame(coords) {
			return changes, nil
		}
	}

	tomlDoc, tomlErr := toml.ParseToml(versionsPath)

	// Check if catalog dependency is already present
	if tomlErr == nil {
		libEntry, ok := toml.FindLib(tomlDoc, mc)

		if ok {
			if _, ok := gradle.FindCatalogDep(build, libEntry.Key); ok {
				return changes, nil
			}
		}
	}

	bom, hasBom := gradle.FindBom(build.Dependencies.List)
	_, hasKtorPlugin := gradle.FindKtorPlugin(build.Plugins.List)

	// Add serialization plugin
	if serPlugin != nil {
		if (hasBom || hasKtorPlugin) && !gradle.HasSerializationPlugin(build.Plugins.List) {
			if kotlinPlugin, ok := gradle.FindKotlinPlugin(build.Plugins.List); ok {
				lang.InsertLnAfter(
					build.Rewriter,
					kotlinPlugin.Statement.GetStop(),
					lang.HiddenTokensToLeft(build.Stream, kotlinPlugin.Statement.GetStart().GetTokenIndex()),
					gradle.KotlinPrefixedPlugin(ktor.SerPluginKotlinId, kotlinPlugin.Version),
				)
			}
		} else if tomlErr == nil {
			_, hasSerPlugin := toml.FindPlugin(tomlDoc, ktor.SerPluginId)
			kotlinPluginEntry, hasKotlinPlugin := toml.FindPlugin(tomlDoc, ktor.KotlinJvmPluginId)

			if !hasSerPlugin && hasKotlinPlugin {
				if vRef, ok := kotlinPluginEntry.Get("version.ref"); ok {
					key := "kotlin-serialization"

					lang.InsertLnAfter(
						tomlDoc.Rewriter,
						kotlinPluginEntry.Expression.GetStop(),
						lang.HiddenTokensToLeft(tomlDoc.Stream, kotlinPluginEntry.Expression.GetStart().GetTokenIndex()),
						toml.PluginEntry(key, ktor.SerPluginId, vRef),
					)

					if len(build.Plugins.List) > 0 {
						lastPlugin := build.Plugins.List[len(build.Plugins.List)-1]

						lang.InsertLnAfter(
							build.Rewriter,
							lastPlugin.Statement.GetStop(),
							lang.HiddenTokensToLeft(build.Stream, lastPlugin.Statement.GetStart().GetTokenIndex()),
							gradle.CatalogPlugin(key),
						)
					}
				}
			}
		}
	}

	// Add dependency with BOM defined
	if hasBom || hasKtorPlugin {
		if kDep, ok := gradle.FindKtorDep(build.Dependencies.List, mc.IsTest); ok {
			suffix := ""
			if strings.HasSuffix(kDep.Path, "-jvm") {
				suffix = "-jvm"
			}

			lang.InsertLnAfter(
				build.Rewriter,
				kDep.Statement.GetStop(),
				lang.HiddenTokensToLeft(build.Stream, kDep.Statement.GetStart().GetTokenIndex()),
				gradle.RawDependencyNoVersion(mc, suffix),
			)

		} else if hasBom {
			lang.InsertLnAfter(
				build.Rewriter,
				bom.GetStop(),
				lang.HiddenTokensToLeft(build.Stream, bom.GetStart().GetTokenIndex()),
				gradle.RawDependencyNoVersion(mc, ""),
			)
		}

		changes = append(changes, FileContent{Path: buildPath, Content: build.Rewriter.GetTextDefault()})
		return changes, nil
	}

	// versions catalog file doesn't exist
	if _, err := os.Stat(versionsPath); errors.Is(err, os.ErrNotExist) {
		changes = append(changes, FileContent{Path: versionsPath, Content: toml.NewTomlWithKtor(mc)})

		suffix := ""
		if len(build.Dependencies.List) == 0 {
			suffix = "\n"
		}

		lang.InsertLnAfter(
			build.Rewriter,
			build.Dependencies.Statements.GetStop(),
			lang.DefaultIndent,
			gradle.CatalogDependency(mc.Artifact)+suffix,
		)

		changes = append(changes, FileContent{Path: buildPath, Content: build.Rewriter.GetTextDefault()})

		return changes, nil
	}

	modified, err := toml.AddLib(tomlDoc, mc)

	if err != nil {
		return changes, err
	}

	changes = append(changes, FileContent{Path: versionsPath, Content: modified})

	modified, err = gradle.AddCatalogDep(build, mc.Artifact)

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
