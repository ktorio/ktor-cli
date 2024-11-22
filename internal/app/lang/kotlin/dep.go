package kotlin

import (
	"errors"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/ktorio/ktor-cli/internal/app/ktor"
	"github.com/ktorio/ktor-cli/internal/app/lang"
	parser "github.com/ktorio/ktor-cli/internal/app/lang/parsers/kotlin"
	"strings"
)

func NewParser(fp string) (*parser.KotlinParser, error) {
	input, err := antlr.NewFileStream(fp)

	if err != nil {
		return nil, err
	}

	lexer := parser.NewKotlinLexer(&input.InputStream)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	return parser.NewKotlinParser(stream), nil
}

func AddRawDepAfter(p *parser.KotlinParser, st parser.IStatementContext, mc ktor.MavenCoords) string {
	stream := p.GetTokenStream().(*antlr.CommonTokenStream)
	rewriter := antlr.NewTokenStreamRewriter(stream)
	indent := lang.HiddenTokensToLeft(stream, st.GetStart().GetTokenIndex())
	rewriter.InsertAfterDefault(
		st.GetStop().GetTokenIndex(),
		"\n"+indent+fmt.Sprintf("implementation(%s)", lang.Quote(mc.Group+":"+mc.Artifact)),
	)
	return rewriter.GetTextDefault()
}

func FindBom(p *parser.KotlinParser) (parser.IStatementContext, bool) {
	sts, ok := findDepsStatements(p.Script())

	if !ok {
		return nil, false
	}

	for _, st := range sts {
		pus, ok := lang.FindChild[parser.IPostfixUnaryExpressionContext](st)

		if !ok {
			continue
		}

		if pus.GetChildCount() == 0 {
			continue
		}

		pe, ok := pus.GetChild(0).(parser.IPrimaryExpressionContext)

		if !ok {
			continue
		}

		if pe.SimpleIdentifier().GetText() != "implementation" {
			continue
		}

		pus2, ok := pus.GetChild(1).(parser.IPostfixUnarySuffixContext)

		if !ok || pus2.CallSuffix() == nil {
			continue
		}

		for _, va := range findValueArguments(pus2.CallSuffix()) {
			ps, ok := lang.FindChild[parser.IPostfixUnaryExpressionContext](va)

			if !ok {
				continue
			}

			if id := ps.PrimaryExpression().SimpleIdentifier(); id == nil || id.GetText() != "platform" {
				continue
			}

			pus2, ok = ps.GetChild(1).(parser.IPostfixUnarySuffixContext)

			if !ok || pus2.CallSuffix() == nil {
				continue
			}

			for _, va := range findValueArguments(pus2.CallSuffix()) {
				if strings.HasPrefix(lang.Unquote(va.GetText()), "io.ktor:ktor-bom") {
					return st, true
				}
			}
		}
	}

	return nil, false
}

func findValueArguments(cf parser.ICallSuffixContext) []parser.IValueArgumentContext {
	vas := cf.ValueArguments()

	if vas == nil {
		return nil
	}

	return vas.AllValueArgument()
}

func FindCatalogDep(p *parser.KotlinParser, catalogKey string) bool {
	sts, ok := findDepsStatements(p.Script())

	if !ok {
		return false
	}

	for _, st := range sts {
		if st.GetText() == fmt.Sprintf("implementation(libs.%s)", strings.ReplaceAll(catalogKey, "-", ".")) {
			return true
		}
	}

	return false
}

func AddDependency(p *parser.KotlinParser, versionKey string) (string, error) {
	stream := p.GetTokenStream().(*antlr.CommonTokenStream)

	sts, ok := findDepsStatements(p.Script())

	if !ok {
		return "", errors.New("kotlin: could not find dependencies")
	}

	st, _, ok := findDep(sts, "libs.ktor")

	if !ok {
		return "", errors.New("kotlin: could not find catalog Ktor dependencies")
	}

	indent := ""
	for _, t := range stream.GetHiddenTokensToLeft(st.GetStart().GetTokenIndex(), antlr.TokenHiddenChannel) {
		indent += t.GetText()
	}

	obj := strings.ReplaceAll(versionKey, "-", ".")
	rewriter := antlr.NewTokenStreamRewriter(stream)
	rewriter.InsertAfterDefault(st.GetStop().GetTokenIndex(), "\n"+indent+fmt.Sprintf("implementation(libs.%s)", obj))

	return rewriter.GetTextDefault(), nil
}

func findDep(depStatements []parser.IStatementContext, depPrefix string) (parser.IStatementContext, parser.IValueArgumentContext, bool) {
	for _, st := range depStatements {
		e, ok := lang.FindChild[parser.IPostfixUnaryExpressionContext](st)

		if !ok {
			continue
		}

		id, ok := e.GetChild(0).GetChild(0).(parser.ISimpleIdentifierContext)

		if !ok {
			continue
		}

		if id.GetText() != "implementation" {
			continue
		}

		suf, ok := e.GetChild(1).(parser.IPostfixUnarySuffixContext)

		if !ok {
			continue
		}

		vargs, ok := lang.FindChild[parser.IValueArgumentsContext](suf)

		if !ok {
			continue
		}

		for _, va := range lang.ChildrenOfType[parser.IValueArgumentContext](vargs) {
			if strings.HasPrefix(va.GetText(), depPrefix) {
				return st, va, true
			}
		}
	}

	return nil, nil, false
}

func FindRawDep(sts parser.IStatementsContext, mc ktor.MavenCoords) bool {
	_, _, ok := findDep(sts.AllStatement(), lang.Quote(mc.Group+":"+mc.Artifact))
	return ok
}

func findDepsStatements(script parser.IScriptContext) ([]parser.IStatementContext, bool) {
	for _, st := range script.AllStatement() {
		e, ok := lang.FindChild[*parser.PostfixUnaryExpressionContext](st)

		if !ok {
			continue
		}

		id, ok := e.GetChild(0).GetChild(0).(parser.ISimpleIdentifierContext)

		if !ok {
			continue
		}

		if id.GetText() != "dependencies" {
			continue
		}

		suf, ok := e.GetChild(1).(*parser.PostfixUnarySuffixContext)

		if !ok {
			continue
		}

		lit, ok := lang.FindChild[*parser.LambdaLiteralContext](suf)

		if !ok {
			continue
		}

		return lit.Statements().AllStatement(), true
	}
	return nil, false
}
