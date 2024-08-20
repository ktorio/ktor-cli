package jdk

import (
	"github.com/ktorio/ktor-cli/internal/app"
	"github.com/ktorio/ktor-cli/internal/app/config"
	"log"
	"net/http"
	"os"
	"runtime"
)

func FetchRecommendedJdk(client *http.Client, homeDir string, logger *log.Logger) (string, error) {
	jdksDir := config.JdksDir(homeDir)
	err := os.MkdirAll(jdksDir, 0755)

	if err != nil {
		return "", &app.Error{Err: err, Kind: app.JdksDirError}
	}

	logger.Printf("Fetching %s\n", getRecommendedJdk())
	jdkPath, err := fetch(client, getRecommendedJdk(), config.JdksDir(homeDir), logger)
	if err != nil {
		return "", err
	}

	config.SetValue("jdk", jdkPath)
	_ = config.Commit()

	return jdkPath, err
}

func getRecommendedJdk() *Descriptor {
	platform := runtime.GOOS

	if platform == "darwin" {
		platform = "macos"
	}

	arch := runtime.GOARCH
	switch runtime.GOARCH {
	case "amd64":
		arch = "x64"
	case "arm":
		arch = "aarch64"
	}

	return &Descriptor{Platform: platform, Arch: arch, Version: "17"}
}
