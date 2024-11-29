package gradle

import (
	"fmt"
	"github.com/ktorio/ktor-cli/internal/app/ktor"
	"github.com/ktorio/ktor-cli/internal/app/lang"
	"strings"
)

func KotlinPrefixedPlugin(id, version string) string {
	return fmt.Sprintf("kotlin(\"%s\") version \"%s\"", id, version)
}

func CatalogPlugin(catalogKey string) string {
	return fmt.Sprintf("alias(libs.plugins.%s)", strings.ReplaceAll(catalogKey, "-", "."))
}

func RawDependencyNoVersion(mc ktor.MavenCoords, suffix string) string {
	fn := "implementation"
	if mc.IsTest {
		fn = "testImplementation"
	}

	return fmt.Sprintf("%s(%s)", fn, lang.Quote(mc.Group+":"+mc.Artifact+suffix))
}

func DependencyWithVersionVar(mc ktor.MavenCoords, version string, suffix string) string {
	fn := "implementation"
	if mc.IsTest {
		fn = "testImplementation"
	}

	return fmt.Sprintf("%s(%s)", fn, lang.Quote(mc.Group+":"+mc.Artifact+suffix+":$"+version))
}

func CatalogDependency(artifact string) string {
	return fmt.Sprintf("implementation(libs.%s)", strings.ReplaceAll(artifact, "-", "."))
}

func NewDepsWithKtor(libKey string) string {
	return fmt.Sprintf(`dependencies {
    implementation(libs.%s)
}`, strings.ReplaceAll(libKey, "-", "."))
}
