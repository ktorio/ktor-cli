package jdk

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

func DownloadJdk(client *http.Client, platform, arch, version string, w io.Writer) error {
	if !hasJdkBuild(platform, arch, version) {
		return errors.New(fmt.Sprintf("cannot download JDK for the platform %s and architecture %s", platform, arch))
	}

	ext := "tar.gz"
	if platform == "windows" {
		ext = "zip"
	}

	url := fmt.Sprintf("https://corretto.aws/downloads/latest/amazon-corretto-%s-%s-%s-jdk.%s", version, arch, platform, ext)

	resp, err := client.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	_, err = io.Copy(w, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func hasJdkBuild(platform, arch, version string) bool {
	if version != "11" && version != "17" && version != "21" {
		return false
	}

	switch platform {
	case "linux":
		if arch == "x64" || arch == "aarch64" {
			return true
		}
	case "windows":
		if arch == "x64" {
			return true
		}
	case "macos":
		if arch == "x64" || arch == "aarch64" {
			return true
		}
	case "alpine":
		if arch == "x64" || arch == "aarch64" {
			return true
		}
	}
	return false
}
