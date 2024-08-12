package network

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ktorio/ktor-cli/internal/app"
	"github.com/ktorio/ktor-cli/internal/app/config"
	"net"
	"net/http"
	"syscall"
)

func FetchSettings(client *http.Client) (*DefaultSettings, error) {
	resp, err := client.Get(fmt.Sprintf("%s/project/settings", config.GenBaseUrl()))

	if err != nil {
		if errors.Is(err, syscall.ECONNREFUSED) || errors.Is(err, syscall.ECONNABORTED) || errors.Is(err, syscall.ECONNRESET) {
			return nil, &app.Error{Err: err, Kind: app.NetworkError}
		}

		var dnsErr *net.DNSError
		if errors.As(err, &dnsErr) {
			return nil, &app.Error{Err: err, Kind: app.NetworkError}
		}

		if errors.Is(err, context.DeadlineExceeded) {
			return nil, &app.Error{Err: err, Kind: app.NetworkError}
		}

		var certErr *tls.CertificateVerificationError
		if errors.As(err, &certErr) {
			return nil, &app.Error{Err: err, Kind: app.ServerError}
		}

		return nil, &app.Error{Err: err, Kind: app.UnknownError}
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		statusErr := app.StatusError(resp, "fetch settings")
		if resp.StatusCode == http.StatusNotFound || resp.StatusCode >= 500 {
			return nil, &app.Error{Err: statusErr, Kind: app.ServerError}
		}

		if resp.StatusCode == http.StatusBadRequest {
			return nil, &app.Error{Err: statusErr, Kind: app.InternalError}
		}

		return nil, &app.Error{Err: statusErr, Kind: app.ServerError}
	}

	dec := json.NewDecoder(resp.Body)
	dec.DisallowUnknownFields()
	var settings DefaultSettings
	err = dec.Decode(&settings)

	if err != nil {
		return nil, app.Error{Err: errors.New("fetch settings: " + err.Error()), Kind: app.ServerError}
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
