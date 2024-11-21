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

func (e AddModuleError) Error() string {
	return fmt.Sprintf("cannot add module. error #%d: %v", e.Kind, e.Err)
}

func Add(mc ktor.MavenCoords, projectDir string) error {
	versionsPath := filepath.Join(projectDir, "gradle", "libs.versions.toml")
	modified, err := toml.AddLib(versionsPath, mc)

	if err != nil {
		return err
	}

	fmt.Println(getDiff(versionsPath, modified))

	buildPath := filepath.Join(projectDir, "build.gradle.kts")
	modified, err = kotlin.AddDependency(buildPath, mc.Artifact)

	if err != nil {
		return err
	}

	fmt.Println(getDiff(buildPath, modified))

	return nil
}

func getDiff(fp string, new string) string {
	old, err := os.ReadFile(fp)

	if err != nil {
		return ""
	}

	edits := myers.ComputeEdits(span.URIFromPath(fp), string(old), new)
	return fmt.Sprint(gotextdiff.ToUnified(filepath.Base(fp), filepath.Base(fp)+"~new", string(old), edits))
}
