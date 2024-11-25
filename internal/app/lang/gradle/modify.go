package gradle

import (
	"errors"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/ktorio/ktor-cli/internal/app/ktor"
	"github.com/ktorio/ktor-cli/internal/app/lang"
	parser "github.com/ktorio/ktor-cli/internal/app/lang/parsers/kotlin"
	"strings"
)

func AddRawDepAfter(build *BuildRoot, st parser.IStatementContext, mc ktor.MavenCoords) string {
	rewriter := antlr.NewTokenStreamRewriter(build.Stream)
	indent := lang.HiddenTokensToLeft(build.Stream, st.GetStart().GetTokenIndex())
	rewriter.InsertAfterDefault(
		st.GetStop().GetTokenIndex(),
		"\n"+indent+fmt.Sprintf("implementation(%s)", lang.Quote(mc.Group+":"+mc.Artifact)),
	)
	return rewriter.GetTextDefault()
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
	rewriter := antlr.NewTokenStreamRewriter(build.Stream)
	rewriter.InsertAfterDefault(ktorSt.GetStop().GetTokenIndex(), "\n"+indent+fmt.Sprintf("implementation(libs.%s)", obj))

	return rewriter.GetTextDefault(), nil
}
