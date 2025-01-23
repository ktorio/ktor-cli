package command

import (
	"bufio"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/hexops/gotextdiff"
	"github.com/hexops/gotextdiff/myers"
	"github.com/hexops/gotextdiff/span"
	"github.com/ktorio/ktor-cli/internal/app/ktor"
	"github.com/ktorio/ktor-cli/internal/app/lang"
	"github.com/ktorio/ktor-cli/internal/app/lang/gradle"
	"github.com/ktorio/ktor-cli/internal/app/lang/kotlin"
	"github.com/ktorio/ktor-cli/internal/app/lang/toml"
	"github.com/ktorio/ktor-cli/internal/app/utils"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type ErrorKind int

type AddModuleError struct {
	Kind ErrorKind
	Err  error
}

type FileContent struct {
	Path    string
	Content string
}

type AddDependencyResult string

const (
	Success                          AddDependencyResult = "success"
	Error                                                = "error"
	MultiplatformProjectNotSupported                     = "multiplatform-not-supported"
	MavenProjectNotSupported                             = "maven-not-supported"
	GroovyDslNotSupported                                = "groovy-dsl-not-supported"
	BuildGradleKtsNotFound                               = "build-gradle-kts-not-found"
)

func (e AddModuleError) Error() string {
	return fmt.Sprintf("cannot add module. error #%d: %v", e.Kind, e.Err)
}

func Add(mc ktor.MavenCoords, projectDir string, serPlugin *ktor.GradlePlugin) error {
	files, result, err := addDependency(mc, projectDir, serPlugin)

	switch result {
	case MultiplatformProjectNotSupported:
		fmt.Fprintln(os.Stderr, "Unable to add the Ktor module to a Kotlin Multiplatform project (not supported yet).")
		os.Exit(1)
	case MavenProjectNotSupported:
		fmt.Fprintln(os.Stderr, "Unable to add the Ktor module to a Maven project (not supported yet).")
		os.Exit(1)
	case GroovyDslNotSupported:
		fmt.Fprintln(os.Stderr, "Unable to add the Ktor module to a Gradle project with Groovy DSL (not supported yet).")
		os.Exit(1)
	case BuildGradleKtsNotFound:
		fmt.Fprintf(os.Stderr, "Unable to find build.gradle.kts file in the project directory %s.\n", projectDir)
		os.Exit(1)
	case Error:
		return err
	}

	if len(files) == 0 {
		fmt.Println("Nothing to change.")
		return nil
	}

	fmt.Printf("Below you can find suggested changes to add '%s'.\n", mc.String())
	fmt.Println("If you consider them incorrect, please file an issue at https://youtrack.jetbrains.com/newIssue?project=ktor.")
	fmt.Println()
	for _, f := range files {
		fmt.Println(getDiff(f.Path, f.Content))
	}

	fmt.Print("Do you want to apply the changes (y/n)? ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	answer := scanner.Text()

	if answer == "y" || answer == "Y" || answer == "yes" || answer == "Yes" {
		err = applyChanges(files)

		if err == nil {
			fmt.Println("The changes have been successfully applied.")
		} else {
			fmt.Println("An error occurred applying the changes.")
			// TODO: Report the error
		}

		return err
	}

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

func SearchKtorVersion(projectDir string) (version string, found bool) {
	build, err := gradle.ParseBuildFile(filepath.Join(projectDir, "build.gradle.kts"))

	if err != nil {
		log.Println(err)
		return
	}

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

	tomlPath, ok := toml.FindVersionsPath(projectDir)

	if !ok {
		found = false
		return
	}

	tomlDoc, tomlErr := toml.ParseCatalogToml(tomlPath)

	if tomlErr == nil {
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

func addDependency(mc ktor.MavenCoords, projectDir string, serPlugin *ktor.GradlePlugin) ([]FileContent, AddDependencyResult, error) {
	buildPath := filepath.Join(projectDir, "build.gradle.kts")
	var changes []FileContent

	if !utils.Exists(buildPath) {
		if utils.Exists(filepath.Join(projectDir, "pom.xml")) {
			return changes, MavenProjectNotSupported, nil
		}

		if utils.Exists(filepath.Join(projectDir, "build.gradle")) {
			return changes, GroovyDslNotSupported, nil
		}

		return changes, BuildGradleKtsNotFound, nil
	}

	build, err := gradle.ParseBuildFile(buildPath)
	if err != nil {
		return changes, Error, err
	}

	if isKmpProject(build, projectDir) {
		return changes, MultiplatformProjectNotSupported, nil
	}

	// Check {BOM, Hardcoded, variable as version} dependency exist
	for _, d := range build.Dependencies.List {
		if d.Kind == gradle.VersionCatalogDep {
			continue
		}

		if m, ok := ktor.ParseMavenCoords(d.Path); ok && mc.RoughlySame(m) {
			return changes, Success, nil
		}
	}

	tomlPath, tomlFound := toml.FindVersionsPath(projectDir)

	var tomlDoc *toml.Document
	var tomlError error
	if tomlFound {
		tomlDoc, tomlError = toml.ParseCatalogToml(tomlPath)
	}

	// Check Catalog dependency exist
	if tomlFound && tomlError == nil {
		libEntry, ok := toml.FindLib(tomlDoc, mc)

		if ok {
			if _, ok := gradle.FindCatalogDep(build, libEntry.Key); ok {
				return changes, Success, nil
			}
		}
	}

	if ktorDep, coords, ok := gradle.FindDepFunc(build.Dependencies.List, func(mc ktor.MavenCoords) bool {
		return mc.Group == "io.ktor" && strings.HasPrefix(mc.Version, "$")
	}); ok {

		if vd, ok := gradle.FindVarDecl(build.TopLevelVars, func(v *gradle.VarDecl) bool {
			return v.Id == kotlin.GetVarId(coords.Version)
		}); ok {
			lang.InsertLnAfter(
				build.Rewriter,
				ktorDep.Statement.GetStop(),
				lang.HiddenTokensToLeft(build.Stream, ktorDep.Statement.GetStart().GetTokenIndex()),
				gradle.DependencyWithVersionVar(mc, vd.Id, gradle.PlatformSuffix(coords.Artifact)),
			)
			changes = append(changes, FileContent{Path: buildPath, Content: build.Rewriter.GetTextDefault()})
			return changes, Success, nil
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
		} else if tomlFound && tomlError == nil {
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
			lang.InsertLnAfter(
				build.Rewriter,
				kDep.Statement.GetStop(),
				lang.HiddenTokensToLeft(build.Stream, kDep.Statement.GetStart().GetTokenIndex()),
				gradle.RawDependencyNoVersion(mc, gradle.PlatformSuffix(kDep.Path)),
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
		return changes, Success, nil
	}

	versionsPath, ok := toml.FindVersionsPath(projectDir)
	// versions catalog file doesn't exist
	if !ok {
		changes = append(changes, FileContent{Path: versionsPath, Content: toml.NewTomlWithKtor(mc)})

		if build.Dependencies.Statements != nil {
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
		}

		if len(build.Rewriter.GetProgram(antlr.DefaultProgramName)) > 0 {
			changes = append(changes, FileContent{Path: buildPath, Content: build.Rewriter.GetTextDefault()})
		}

		return changes, Success, nil
	}

	if tomlFound && tomlError == nil {
		modified, err := toml.AddLib(tomlDoc, mc)

		if err != nil {
			return changes, Error, err
		}

		changes = append(changes, FileContent{Path: versionsPath, Content: modified})

		modified, err = gradle.AddCatalogDep(build, mc.Artifact)

		if err != nil {
			return changes, Error, err
		}

		changes = append(changes, FileContent{Path: buildPath, Content: modified})

		return changes, Success, nil
	}

	return changes, Error, nil
}

func isKmpProject(build *gradle.BuildRoot, projectDir string) bool {
	for _, p := range build.Plugins.List {
		if (p.Prefix == "kotlin" && p.Id == "multiplatform") || (p.Prefix == "id" && p.Id == "org.jetbrains.kotlin.multiplatform") {
			return true
		}
	}

	tomlPath, tomlFound := toml.FindVersionsPath(projectDir)

	if tomlFound {
		tomlDoc, err := toml.ParseCatalogToml(tomlPath)

		if err != nil {
			return false
		}

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

func getDiff(fp string, new string) string {
	old, err := os.ReadFile(fp)

	if err != nil {
		old = []byte{}
	}

	edits := myers.ComputeEdits(span.URIFromPath(fp), string(old), new)
	return fmt.Sprint(gotextdiff.ToUnified(filepath.Base(fp), filepath.Base(fp)+"~new", string(old), edits))
}
