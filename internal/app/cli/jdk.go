package cli

import (
	"errors"
	"fmt"
	"github.com/ktorio/ktor-cli/internal/app"
	"github.com/ktorio/ktor-cli/internal/app/config"
	"github.com/ktorio/ktor-cli/internal/app/jdk"
	"log"
	"net/http"
	"os"
)

func DownloadJdk(homeDir string, client *http.Client, logger *log.Logger, attempt int) (string, error) {
	jdkPath, err := jdk.FetchRecommendedJdk(client, homeDir, logger)

	if err != nil {
		var e *app.Error
		var pe *os.PathError

		isAppError := errors.As(err, &e)
		// JDK is already downloaded but the config doesn't reflect that
		if isAppError && errors.As(e.Err, &pe) && errors.Is(pe.Err, os.ErrExist) {
			jdkPath = pe.Path
			config.SetValue("jdk", pe.Path)
			_ = config.Commit()
			return pe.Path, nil
		}

		if isAppError && e.Kind == app.JdkVerificationFailed {
			fmt.Println("JDK Verification Failed. Trying again...")

			if attempt >= 2 {
				return "", err
			}

			return DownloadJdk(homeDir, client, logger, attempt+1)
		}

		return "", err
	}

	return jdkPath, nil
}
