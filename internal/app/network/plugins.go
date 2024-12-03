package network

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ktorio/ktor-cli/internal/app"
	"github.com/ktorio/ktor-cli/internal/app/config"
	"net/http"
)

type Plugin struct {
	Id              string   `json:"xmlId"`
	Name            string   `json:"name"`
	Group           string   `json:"group"`
	Description     string   `json:"description"`
	RequiredPlugins []string `json:"requiredFeatures"`
}

func FetchPlugins(client *http.Client, ktorVersion string, ctx context.Context) ([]Plugin, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/features/%s", config.GenBaseUrl(), ktorVersion), nil)

	if err != nil {
		return nil, &app.Error{Err: err, Kind: app.InternalError}
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", ctx.Value("user-agent").(string))

	resp, err := client.Do(req)

	if err != nil {
		return nil, ConvertResponseError(err, app.GenServerError)
	}

	defer resp.Body.Close()

	tag := fmt.Sprintf("fetch plugins for %s", ktorVersion)
	if err = CheckResponseStatus(resp, tag, app.GenServerError); err != nil {
		return nil, err
	}

	dec := json.NewDecoder(resp.Body)
	var plugins []Plugin
	err = dec.Decode(&plugins)

	if err != nil {
		return nil, app.Error{Err: errors.New(fmt.Sprintf("%s: %s", tag, err.Error())), Kind: app.GenServerError}
	}

	return plugins, nil
}
