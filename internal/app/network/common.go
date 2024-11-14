package network

import (
	"context"
	"crypto/tls"
	"errors"
	"github.com/ktorio/ktor-cli/internal/app"
	"net"
	"net/http"
	"syscall"
)

func convertResponseError(err error) error {
	if errors.Is(err, syscall.ECONNREFUSED) || errors.Is(err, syscall.ECONNABORTED) || errors.Is(err, syscall.ECONNRESET) {
		return &app.Error{Err: err, Kind: app.NetworkError}
	}

	var dnsErr *net.DNSError
	if errors.As(err, &dnsErr) {
		return &app.Error{Err: err, Kind: app.NetworkError}
	}

	if errors.Is(err, context.DeadlineExceeded) {
		return &app.Error{Err: err, Kind: app.NetworkError}
	}

	var certErr *tls.CertificateVerificationError
	if errors.As(err, &certErr) {
		return &app.Error{Err: err, Kind: app.GenServerError}
	}

	return &app.Error{Err: err, Kind: app.UnknownError}
}

func checkResponseStatus(resp *http.Response, endpoint string) error {
	if resp.StatusCode != http.StatusOK {
		statusErr := app.StatusError(resp, endpoint)
		if resp.StatusCode == http.StatusNotFound || resp.StatusCode >= 500 {
			return &app.Error{Err: statusErr, Kind: app.GenServerError}
		}

		if resp.StatusCode == http.StatusBadRequest {
			return &app.Error{Err: statusErr, Kind: app.InternalError}
		}

		return &app.Error{Err: statusErr, Kind: app.GenServerError}
	}

	return nil
}
