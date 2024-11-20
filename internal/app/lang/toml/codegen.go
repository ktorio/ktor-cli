package toml

import (
	"fmt"
	"github.com/ktorio/ktor-cli/internal/app/ktor"
)

func PluginEntry(key, pluginId, version string) string {
	return fmt.Sprintf("%s = { id = \"%s\", version.ref = \"%s\" }", key, pluginId, version)
}

func LibEntryModule(versionKey string, mc ktor.MavenCoords) string {
	return fmt.Sprintf("%s = { module = \"%s:%s\", version.ref = \"%s\" }", mc.Artifact, mc.Group, mc.Artifact, versionKey)
}

func LibEntryGroupName(versionKey string, mc ktor.MavenCoords) string {
	return fmt.Sprintf("%s = { group = \"%s\", name = \"%s\", version.ref = \"%s\" }", mc.Artifact, mc.Group, mc.Artifact, versionKey)
}

func NewTomlWithKtor(mc ktor.MavenCoords) string {
	return fmt.Sprintf(`[versions]
ktor = "%s"

[libraries]
%s = { module = "%s:%s", version.ref = "ktor" }
`, mc.Version, mc.Artifact, mc.Group, mc.Artifact)
}

func NewLibraryTableWithKtor(mc ktor.MavenCoords) string {
	return fmt.Sprintf("[libraries]\n%s = { module = \"%s:%s\", version.ref = \"ktor\" }", mc.Artifact, mc.Group, mc.Artifact)
}

func VersionEntry(key, version string) string {
	return fmt.Sprintf("%s = \"%s\"", key, version)
}
