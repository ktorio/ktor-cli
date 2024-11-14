package network

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ktorio/ktor-cli/internal/app"
	"github.com/ktorio/ktor-cli/internal/app/config"
	"net/http"
)

func FetchSettings(client *http.Client) (*DefaultSettings, error) {
	resp, err := client.Get(fmt.Sprintf("%s/project/settings", config.GenBaseUrl()))

	if err != nil {
		return nil, convertResponseError(err)
	}

	defer resp.Body.Close()

	if err = checkResponseStatus(resp, "fetch settings"); err != nil {
		return nil, err
	}

	dec := json.NewDecoder(resp.Body)
	var settings DefaultSettings
	err = dec.Decode(&settings)

	if err != nil {
		return nil, app.Error{Err: errors.New("fetch settings: " + err.Error()), Kind: app.GenServerError}
	}

	return &settings, nil
}

type DefaultSettings struct {
	ProjectName    StringParam `json:"project_name"`
	CompanyWebsite StringParam `json:"company_website"`
	Engine         EnumParam   `json:"engine"`
	KtorVersion    EnumParam   `json:"ktor_version"`
	KotlinVersion  EnumParam   `json:"kotlin_version"`
	BuildSystem    EnumParam   `json:"build_system"`
	ConfigType     EnumParam   `json:"configuration_in"`
}

type StringParam struct {
	DefaultVal string `json:"default"`
}

type EnumParam struct {
	Options   []EnumOption `json:"options"`
	DefaultId string       `json:"default_id"`
}

type EnumOption struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
