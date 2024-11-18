package openapi

import (
	"github.com/ktorio/ktor-cli/internal/app"
	"github.com/ktorio/ktor-cli/internal/app/i18n"
	"github.com/ktorio/ktor-cli/internal/app/network"
	"github.com/ktorio/ktor-cli/internal/app/progress"
	"github.com/ktorio/ktor-cli/internal/app/utils"
	"io"
	"net/http"
)

func DownloadJar(client *http.Client, jarUrl string) ([]byte, error) {
	resp, err := client.Get(jarUrl)

	if err != nil {
		return nil, network.ConvertResponseError(err, app.OpenApiDownloadJarError)
	}

	defer resp.Body.Close()

	if err = network.CheckResponseStatus(resp, "fetch OpenAPI JAR", app.OpenApiDownloadJarError); err != nil {
		return nil, err
	}

	reader, progressBar := progress.NewReader(
		resp.Body,
		i18n.Get(i18n.DownloadingOpenApiJarProgress),
		utils.ContentLength(resp),
		true,
	)
	defer progressBar.Done()

	bodyBytes, err := io.ReadAll(reader)
	if err != nil {
		return nil, &app.Error{Err: err, Kind: app.OpenApiDownloadJarError}
	}

	return bodyBytes, nil
}
