package generate

import (
	"bytes"
	"context"
	"errors"
	"github.com/ktorio/ktor-cli/internal/app"
	"github.com/ktorio/ktor-cli/internal/app/archive"
	"github.com/ktorio/ktor-cli/internal/app/i18n"
	"github.com/ktorio/ktor-cli/internal/app/network"
	"github.com/ktorio/ktor-cli/internal/app/progress"
	"github.com/ktorio/ktor-cli/internal/app/utils"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

// Project Returns *app.Error on error
func Project(client *http.Client, logger *log.Logger, projectDir, project string, plugins []string, versionCatalog bool, ctx context.Context) error {
	err := os.MkdirAll(projectDir, 0755)
	logger.Printf(i18n.Get(i18n.CreatingDir, projectDir))

	if _, err = os.Stat(projectDir); errors.Is(err, os.ErrNotExist) {
		if err != nil {
			var pe *os.PathError
			errors.As(err, &pe)

			if errors.Is(pe.Err, os.ErrExist) || errors.Is(pe.Err, os.ErrPermission) {
				return &app.Error{Err: err, Kind: app.ProjectDirError}
			}

			return &app.Error{Err: err, Kind: app.UnknownError}
		}
	} else if !utils.IsDirEmpty(projectDir) {
		return &app.Error{Err: &os.PathError{Err: os.ErrExist, Path: projectDir}, Kind: app.ProjectDirError}
	}

	logger.Printf(i18n.Get(i18n.CreatingDir, projectDir))

	settings, err := network.FetchSettings(client)

	if err != nil {
		return err
	}

	logger.Println(i18n.Get(i18n.RequestGenServer))

	versionCatalogValue := "false"
	if versionCatalog {
		versionCatalogValue = "true"
	}

	projectPayload := network.ProjectPayload{
		Settings: network.ProjectSettings{
			Name:           project,
			CompanyWebsite: "com.example.com",
			Engine:         settings.Engine.DefaultId,
			BuildSystem:    settings.BuildSystem.DefaultId,
			KtorVersion:    settings.KtorVersion.DefaultId,
			KotlinVersion:  settings.KotlinVersion.DefaultId,
			BuildSystemArgs: map[network.BuildSystemArgs]string{
				network.VersionCatalogBuildArg: versionCatalogValue,
			},
		},
		Plugins:       plugins,
		HasSampleCode: true,
		ConfigType:    settings.ConfigType.DefaultId,
		HasWrapper:    true,
	}

	zipBytes, err := network.NewProject(client, projectPayload, ctx)

	if err != nil {
		if os.IsTimeout(err) {
			return &app.Error{Err: err, Kind: app.GenServerTimeoutError}
		}

		return &app.Error{Err: err, Kind: app.GenServerError}
	}

	logger.Printf(i18n.Get(i18n.ExtractingArchiveToDir, projectDir))

	reader, progressBar := progress.NewReaderAt(
		bytes.NewReader(zipBytes),
		i18n.Get(i18n.ExtractProjectArchive),
		len(zipBytes),
		logger.Writer() == io.Discard,
	)
	defer progressBar.Done()

	_, err = archive.ExtractZip(reader, int64(len(zipBytes)), projectDir, logger)

	if err != nil {
		return &app.Error{Err: err, Kind: app.ProjectExtractError}
	}

	gradlewPath := path.Join(projectDir, "gradlew")
	logger.Printf(i18n.Get(i18n.MakeFileExec, gradlewPath))
	err = os.Chmod(gradlewPath, 0764)

	if err != nil {
		return &app.Error{Err: err, Kind: app.GradlewChmodError}
	}

	return nil
}
