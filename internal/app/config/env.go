package config

import (
	"fmt"
	"os"
	"path/filepath"
)

func GenBaseUrl() string {
	if e, ok := os.LookupEnv("GEN_BASEURL"); ok && e != "" {
		return e
	}

	return "https://ktor-plugin.europe-north1-gke.intellij.net"
}

func CorrettoBaseUrl() string {
	if e, ok := os.LookupEnv("CORRETTO_BASEURL"); ok && e != "" {
		return e
	}

	return "https://corretto.aws"
}

func OpenApiJarUrl() string {
	if e, ok := os.LookupEnv("OPENAPI_JAR_URL"); ok && e != "" {
		return e
	}

	version := "7.11.0"

	return fmt.Sprintf("https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/%s/openapi-generator-cli-%s.jar", version, version)
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

func TempDir(homeDir string) string {
	return filepath.Join(KtorDir(homeDir), "temp")
}

func ktorConfigPath(homeDir string) string {
	return filepath.Join(KtorDir(homeDir), "config.json")
}
