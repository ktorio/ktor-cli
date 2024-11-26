package gradle

import (
	"errors"
	"fmt"
	"github.com/ktorio/ktor-cli/internal/app/ktor"
	"github.com/ktorio/ktor-cli/internal/app/lang"
	parser "github.com/ktorio/ktor-cli/internal/app/lang/parsers/kotlin"
	"strings"
)

func AddRawDepAfter(build *BuildRoot, st parser.IStatementContext, mc ktor.MavenCoords) {
	indent := lang.HiddenTokensToLeft(build.Stream, st.GetStart().GetTokenIndex())
	build.Rewriter.InsertAfterDefault(
		st.GetStop().GetTokenIndex(),
		"\n"+indent+fmt.Sprintf("implementation(%s)", lang.Quote(mc.Group+":"+mc.Artifact)),
	)
}

func AddDependency(build *BuildRoot, versionKey string) (string, error) {
	var ktorSt parser.IStatementContext
	for _, dep := range build.Dependencies.List {
		if dep.Kind == VersionCatalogDep && strings.HasPrefix(dep.Path, "libs.ktor") {
			ktorSt = dep.Statement
			break
		}
	}

	if ktorSt == nil {
		return "", errors.New("kotlin: could not find Ktor catalog dependencies")
	}

	obj := strings.ReplaceAll(versionKey, "-", ".")
	indent := lang.HiddenTokensToLeft(build.Stream, ktorSt.GetStart().GetTokenIndex())
	build.Rewriter.InsertAfterDefault(ktorSt.GetStop().GetTokenIndex(), "\n"+indent+fmt.Sprintf("implementation(libs.%s)", obj))

	return build.Rewriter.GetTextDefault(), nil
}
