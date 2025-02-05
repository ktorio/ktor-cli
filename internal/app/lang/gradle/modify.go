package gradle

import (
	"fmt"
	"github.com/ktorio/ktor-cli/internal/app/lang"
)

func AddCatalogDep(build *BuildRoot, libKey string) (string, error) {
	ktorDep, ok := FindCatalogDepPrefixed(build, "libs.ktor")

	if !ok && build.Repositories.Statement != nil {
		build.Rewriter.InsertAfterDefault(
			build.Repositories.Statement.GetStop().GetTokenIndex(),
			fmt.Sprintf("\n\n"+NewDepsWithKtor(libKey)),
		)
		return build.Rewriter.GetTextDefault(), nil
	}

	lang.InsertLnAfter(
		build.Rewriter,
		ktorDep.Statement.GetStop(),
		lang.HiddenTokensToLeft(build.Stream, ktorDep.Statement.GetStart().GetTokenIndex()),
		CatalogDependency(libKey),
	)

	return build.Rewriter.GetTextDefault(), nil
}
