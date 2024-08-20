package config

import (
	"os"
	"path/filepath"
)

func GenBaseUrl() string {
	envUrl := os.Getenv("GEN_BASEURL")

	if len(envUrl) == 0 {
		return "https://ktor-plugin.europe-north1-gke.intellij.net"
	}

	return envUrl
}

func KtorDir(homeDir string) string {
	return filepath.Join(homeDir, ".ktor")
}

func LogPath(homeDir string) string {
	return filepath.Join(KtorDir(homeDir), "run.log")
}

func JdksDir(homeDir string) string {
	return filepath.Join(KtorDir(homeDir), "jdks")
}

func ktorConfigPath(homeDir string) string {
	return filepath.Join(KtorDir(homeDir), "config.json")
}
