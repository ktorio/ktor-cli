package network

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ktorio/ktor-cli/internal/app"
	"github.com/ktorio/ktor-cli/internal/app/config"
	"github.com/ktorio/ktor-cli/internal/app/progress"
	"github.com/ktorio/ktor-cli/internal/app/utils"
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
	Name               string                     `json:"project_name"`
	CompanyWebsite     string                     `json:"company_website"`
	Engine             string                     `json:"engine"`
	BuildSystem        string                     `json:"build_system"`
	KtorVersion        string                     `json:"ktor_version"`
	KotlinVersion      string                     `json:"kotlin_version"`
	BuildSystemArgs    map[BuildSystemArgs]string `json:"build_system_args"`
	CompilationTargets []string                   `json:"compile_targets"`
	ProjectStructure   string                     `json:"project_structure"`
}

type BuildSystemArgs string

const VersionCatalogBuildArg BuildSystemArgs = "version_catalog"

func NewProject(client *http.Client, payload ProjectPayload, ctx context.Context) ([]byte, error) {
	var body bytes.Buffer
	err := json.NewEncoder(&body).Encode(payload)
	if err != nil {
		return nil, &app.Error{Err: err, Kind: app.InternalError}
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/project/generate", config.GenBaseUrl()), &body)

	if err != nil {
		return nil, &app.Error{Err: err, Kind: app.InternalError}
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", ctx.Value("user-agent").(string))

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if err = checkResponseStatus(resp, "new project"); err != nil {
		return nil, err
	}

	reader, progressBar := progress.NewReader(
		resp.Body,
		"Downloading project archive... ",
		utils.ContentLength(resp),
		true,
	)
	defer progressBar.Done()

	bodyBytes, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	return bodyBytes, nil
}
