package jdk

import (
	"errors"
	"fmt"
	"github.com/ktorio/ktor-cli/internal/app"
	"io"
	"log"
	"net/http"
)

func DownloadJdk(client *http.Client, d *Descriptor, logger *log.Logger) ([]byte, error) {
	if !hasJdkBuild(d) {
		return nil, &app.Error{Err: Error{d}, Kind: app.UnableLocateJdkError}
	}

	ext := "tar.gz"
	if d.Platform == "windows" {
		ext = "zip"
	}

	url := fmt.Sprintf("https://corretto.aws/downloads/latest/amazon-corretto-%s-%s-%s-jdk.%s", d.Version, d.Arch, d.Platform, ext)
	logger.Printf("Downloading %s from %s\n", d, url)

	resp, err := client.Get(url)
	if err != nil {
		return nil, &app.Error{Err: err, Kind: app.JdkServerError}
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, &app.Error{
			Err:  errors.New(fmt.Sprintf("download jdk: bad status code %d", resp.StatusCode)),
			Kind: app.JdkServerError,
		}
	}

	return io.ReadAll(resp.Body)
}

func hasJdkBuild(d *Descriptor) bool {
	if d.Version != "11" && d.Version != "17" && d.Version != "21" {
		return false
	}

	switch d.Platform {
	case "linux":
		if d.Arch == "x64" || d.Arch == "aarch64" {
			return true
		}
	case "windows":
		if d.Arch == "x64" {
			return true
		}
	case "macos":
		if d.Arch == "x64" || d.Arch == "aarch64" {
			return true
		}
	case "alpine":
		if d.Arch == "x64" || d.Arch == "aarch64" {
			return true
		}
	}
	return false
}
