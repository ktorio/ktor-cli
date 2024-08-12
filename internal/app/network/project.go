package network

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ktorio/ktor-cli/internal/app"
	"github.com/ktorio/ktor-cli/internal/app/config"
	"io"
	"net/http"
)

type ProjectPayload struct {
	Settings      ProjectSettings `json:"settings"`
	Plugins       []string        `json:"features"`
	HasSampleCode bool            `json:"addDefaultRoutes"`
	ConfigType    string          `json:"configurationOption"`
	HasWrapper    bool            `json:"addWrapper"`
}

type ProjectSettings struct {
	Name            string                     `json:"project_name"`
	CompanyWebsite  string                     `json:"company_website"`
	Engine          string                     `json:"engine"`
	BuildSystem     string                     `json:"build_system"`
	KtorVersion     string                     `json:"ktor_version"`
	KotlinVersion   string                     `json:"kotlin_version"`
	BuildSystemArgs map[BuildSystemArgs]string `json:"build_system_args"`
}

type BuildSystemArgs string

const VersionCatalogBuildArg BuildSystemArgs = "version_catalog"

func NewProject(client *http.Client, payload ProjectPayload) ([]byte, error) {
	var body bytes.Buffer
	err := json.NewEncoder(&body).Encode(payload)
	if err != nil {
		return nil, &app.Error{Err: err, Kind: app.InternalError}
	}

	resp, err := client.Post(fmt.Sprintf("%s/project/generate", config.GenBaseUrl()), "application/json", &body)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bodyBytes, nil
}
