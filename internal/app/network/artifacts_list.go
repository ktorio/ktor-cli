package network

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ktorio/ktor-cli/internal/app"
	"github.com/ktorio/ktor-cli/internal/app/config"
	"net/http"
)

func ListArtifacts(client *http.Client, ktorVersion string) ([]string, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/artifacts/%s/list", config.GenBaseUrl(), ktorVersion), nil)

	if err != nil {
		return nil, &app.Error{Err: err, Kind: app.InternalError}
	}

	resp, err := client.Do(req)

	if err != nil {
		return nil, ConvertResponseError(err, app.ArtifactListError)
	}

	defer resp.Body.Close()

	tag := fmt.Sprintf("list artifacts for %s", ktorVersion)
	if err = CheckResponseStatus(resp, tag, app.ArtifactListError); err != nil {
		return nil, err
	}

	dec := json.NewDecoder(resp.Body)
	var result []string
	err = dec.Decode(&result)

	if err != nil {
		return nil, app.Error{Err: errors.New(fmt.Sprintf("%s: %s", tag, err.Error())), Kind: app.ArtifactListError}
	}

	return result, nil
}
