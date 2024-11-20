package gradle

import (
	parser "github.com/ktorio/ktor-cli/internal/app/lang/parsers/kotlin"
)

func FindBom(deps []Dep) (parser.IStatementContext, bool) {
	for _, dep := range deps {
		if dep.IsBom {
			return dep.Statement, true
		}
	}

	return nil, false
}
