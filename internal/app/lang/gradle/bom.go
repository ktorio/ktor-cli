package gradle

import (
	parser "github.com/ktorio/ktor-cli/internal/app/lang/parsers/kotlin"
)

func FindBom(build *BuildRoot) (parser.IStatementContext, bool) {
	for _, dep := range build.Dependencies.List {
		if dep.IsBom {
			return dep.Statement, true
		}
	}

	return nil, false
}
