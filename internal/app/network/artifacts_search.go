package network

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ktorio/ktor-cli/internal/app"
	"github.com/ktorio/ktor-cli/internal/app/config"
	"net/http"
)

type Artifact struct {
	Name     string `json:"artifact"`
	Group    string `json:"group"`
	IsTest   bool   `json:"isTest"`
	Distance int    `json:"distance"`
}

func SearchArtifacts(client *http.Client, ktorVersion string, searches []string) (map[string][]Artifact, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/artifacts/%s/search", config.GenBaseUrl(), ktorVersion), nil)

	if err != nil {
		return nil, &app.Error{Err: err, Kind: app.InternalError}
	}

	qVals := req.URL.Query()

	for _, s := range searches {
		qVals.Add("q", s)
	}

	req.URL.RawQuery = qVals.Encode()

	resp, err := client.Do(req)

	if err != nil {
		return nil, ConvertResponseError(err, app.ArtifactSearchError)
	}

	defer resp.Body.Close()

	tag := fmt.Sprintf("search artifacts for %s", ktorVersion)
	if err = CheckResponseStatus(resp, tag, app.ArtifactSearchError); err != nil {
		return nil, err
	}

	dec := json.NewDecoder(resp.Body)
	result := make(map[string][]Artifact)
	err = dec.Decode(&result)

	if err != nil {
		return nil, app.Error{Err: errors.New(fmt.Sprintf("%s: %s", tag, err.Error())), Kind: app.ArtifactSearchError}
	}

	return result, nil
}
