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

	//p := build.Parser
	//sts, ok := findDepsStatements(p.Script())
	//
	//if !ok {
	//	return nil, false
	//}
	//
	//for _, st := range sts {
	//	pus, ok := lang.FindChild[parser.IPostfixUnaryExpressionContext](st)
	//
	//	if !ok {
	//		continue
	//	}
	//
	//	if pus.GetChildCount() == 0 {
	//		continue
	//	}
	//
	//	pe, ok := pus.GetChild(0).(parser.IPrimaryExpressionContext)
	//
	//	if !ok {
	//		continue
	//	}
	//
	//	if pe.SimpleIdentifier().GetText() != "implementation" {
	//		continue
	//	}
	//
	//	pus2, ok := pus.GetChild(1).(parser.IPostfixUnarySuffixContext)
	//
	//	if !ok || pus2.CallSuffix() == nil {
	//		continue
	//	}
	//
	//	for _, va := range findValueArguments(pus2.CallSuffix()) {
	//		ps, ok := lang.FindChild[parser.IPostfixUnaryExpressionContext](va)
	//
	//		if !ok {
	//			continue
	//		}
	//
	//		if id := ps.PrimaryExpression().SimpleIdentifier(); id == nil || id.GetText() != "platform" {
	//			continue
	//		}
	//
	//		pus2, ok = ps.GetChild(1).(parser.IPostfixUnarySuffixContext)
	//
	//		if !ok || pus2.CallSuffix() == nil {
	//			continue
	//		}
	//
	//		for _, va := range findValueArguments(pus2.CallSuffix()) {
	//			if strings.HasPrefix(lang.Unquote(va.GetText()), "io.ktor:ktor-bom") {
	//				return st, true
	//			}
	//		}
	//	}
	//}
	//
	//return nil, false
}
