package cli

import (
	"errors"
	"fmt"
	"github.com/ktorio/ktor-cli/internal/app"
	"github.com/ktorio/ktor-cli/internal/app/config"
	"github.com/ktorio/ktor-cli/internal/app/i18n"
	"github.com/ktorio/ktor-cli/internal/app/jdk"
	"log"
	"net/http"
	"os"
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

func ObtainJdk(client *http.Client, verboseLogger *log.Logger, homeDir string) (jdk.Source, string, error) {
	if jh, ok := jdk.JavaHome(); ok {
		if v, err := jdk.GetJavaMajorVersion(jh, homeDir); err == nil && v >= jdk.MinJavaVersion {
			return jdk.FromJavaHome, jh, nil
		}
	}

	if jdkPath, ok := config.GetValue("jdk"); ok {
		if st, err := os.Stat(jdkPath); err == nil && st.IsDir() {
			return jdk.FromConfig, jdkPath, nil
		}
	}

	if jdkPath, ok := jdk.FindLocally(jdk.MinJavaVersion); ok {
		config.SetValue("jdk", jdkPath)
		_ = config.Commit()
		return jdk.Locally, jdkPath, nil
	}

	jdkPath, err := DownloadJdk(homeDir, client, verboseLogger, 0)

	if err != nil {
		return 0, "", err
	}

	return jdk.Downloaded, jdkPath, nil
}
