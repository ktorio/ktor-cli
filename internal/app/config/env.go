package config

import (
	"os"
	"path"
)

func GenBaseUrl() string {
	envUrl := os.Getenv("GEN_BASEURL")

	if len(envUrl) == 0 {
		return "https://ktor-plugin.europe-north1-gke.intellij.net"
	}

	return envUrl
}

func KtorDir(homeDir string) string {
	return path.Join(homeDir, ".ktor")
}

func LogPath(homeDir string) string {
	return path.Join(KtorDir(homeDir), "run.log")
}
