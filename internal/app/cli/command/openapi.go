package command

import (
	"errors"
	"fmt"
	"github.com/ktorio/ktor-cli/internal/app"
	"github.com/ktorio/ktor-cli/internal/app/cli"
	"github.com/ktorio/ktor-cli/internal/app/config"
	"github.com/ktorio/ktor-cli/internal/app/i18n"
	"github.com/ktorio/ktor-cli/internal/app/jdk"
	"github.com/ktorio/ktor-cli/internal/app/network"
	"github.com/ktorio/ktor-cli/internal/app/openapi"
	"github.com/ktorio/ktor-cli/internal/app/utils"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)

func OpenApi(client *http.Client, specPath string, projectName, projectDir string, homeDir string, logger *log.Logger) error {
	err := os.MkdirAll(projectDir, 0755)
	logger.Printf(i18n.Get(i18n.CreatingDir, projectDir))

	var pe *os.PathError

	if errors.As(err, &pe) && (errors.Is(pe.Err, syscall.EROFS) || errors.Is(pe.Err, syscall.EPERM)) {
		return &app.Error{Err: err, Kind: app.ProjectDirError}
	}

	if _, err := os.Stat(projectDir); errors.Is(err, os.ErrNotExist) {
		return &app.Error{Err: err, Kind: app.UnknownError}
	} else if !utils.IsDirEmpty(projectDir) {
		return &app.Error{Err: &os.PathError{Err: os.ErrExist, Path: projectDir}, Kind: app.ProjectDirError}
	}

	jarName := filepath.Base(config.OpenApiJarUrl())
	jarPath := filepath.Join(config.TempDir(homeDir), jarName)

	if _, err := os.Stat(jarPath); errors.Is(err, os.ErrNotExist) {
		jarBytes, err := openapi.DownloadJar(client, config.OpenApiJarUrl())

		if err != nil {
			return err
		}

		f, err := os.Create(jarPath)
		logger.Printf(i18n.Get(i18n.CreateOpenApiJar, jarPath))

		if err != nil {
			return &app.Error{Err: err, Kind: app.OpenApiDownloadJarError}
		}

		defer f.Close()

		_, err = f.Write(jarBytes)

		if err != nil {
			return &app.Error{Err: err, Kind: app.OpenApiDownloadJarError}
		}
	}

	src, jdkPath, err := cli.ObtainJdk(client, logger, homeDir)

	if err != nil {
		return err
	}

	if src == jdk.Downloaded {
		fmt.Printf(i18n.Get(i18n.JdkDownloaded, jdkPath))
	}

	settings, err := network.FetchSettings(client)

	if err != nil {
		return err
	}

	javaExec := filepath.Join(jdkPath, "bin", "java")

	c := []string{javaExec, "-jar", jarPath, "generate", "-g", "kotlin-server", "-i", specPath,
		"--artifact-id", projectName, "--package-name", utils.GetPackage(settings.CompanyWebsite.DefaultVal), "-o", projectDir}

	logger.Printf(i18n.Get(i18n.ExecutingCommand, strings.Join(c, " ")))

	cmd := exec.Command(javaExec, c[1:]...)

	stdout, err := cmd.Output()

	var ee *exec.ExitError
	if errors.As(err, &ee) {
		msg := string(ee.Stderr)

		if strings.Contains(msg, "Unable to access jarfile") {
			return &app.Error{Err: err, Kind: app.OpenApiExecuteJarError}
		}

		return &app.Error{Err: errors.New(msg), Kind: app.ExternalCommandError}
	}

	if err != nil {
		return &app.Error{Err: err, Kind: app.UnknownError}
	}

	logger.Println(string(stdout))
	return nil
}
