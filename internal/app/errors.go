package app

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

const (
	NetworkError ErrorKind = iota
	InternalError
	GenServerError
	GenServerTimeoutError
	UnknownError
	ProjectDirError
	JdksDirError
	ProjectExtractError
	JdkExtractError
	GradlewChmodError
	ExtractRootDirExistError
	UnableLocateJdkError
	JdkServerError
	JdkServerDownloadError
	JdkVerificationFailed
	ExternalCommandError
	OpenApiDownloadJarError
	OpenApiExecuteJarError
	ArtifactSearchError
	ArtifactListError
	ParsingSyntaxError
)

type ErrorKind int

type Error struct {
	Err  error
	Kind ErrorKind
}

func (e Error) Error() string {
	return fmt.Sprintf("%v: %v", e.Kind, e.Err)
}

func StatusError(resp *http.Response, endpoint string) error {
	bodyBytes, err := io.ReadAll(resp.Body)

	body := ""
	if err == nil {
		b := string(bodyBytes)
		if len(b) > 0 {
			body = fmt.Sprintf(": %s", b)
		}
	}

	return errors.New(fmt.Sprintf("%s: unexpected response status %d from the server%s", endpoint, resp.StatusCode, body))
}
