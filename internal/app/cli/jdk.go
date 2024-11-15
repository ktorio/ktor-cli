package cli

import (
	"errors"
	"fmt"
	"github.com/ktorio/ktor-cli/internal/app"
	"github.com/ktorio/ktor-cli/internal/app/i18n"
	"github.com/ktorio/ktor-cli/internal/app/jdk"
	"log"
	"net/http"
)

func DownloadJdk(homeDir string, client *http.Client, logger *log.Logger, attempt int) (string, error) {
	jdkPath, err := jdk.FetchRecommendedJdk(client, homeDir, logger)

	if err != nil {
		var e *app.Error

		if errors.As(err, &e) && e.Kind == app.JdkVerificationFailed {
			fmt.Println(i18n.Get(i18n.JdkVerificationFailed))

			if attempt >= 2 {
				return "", err
			}

			return DownloadJdk(homeDir, client, logger, attempt+1)
		}

		return "", err
	}

	return jdkPath, nil
}
