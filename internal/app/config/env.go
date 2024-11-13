package config

import (
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

	return "https://oss.sonatype.org/content/repositories/snapshots/org/openapitools/openapi-generator-cli/7.10.0-SNAPSHOT/openapi-generator-cli-7.10.0-20241106.082455-66.jar"
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
