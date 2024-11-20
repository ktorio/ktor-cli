package kotlin

import (
	"errors"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/ktorio/ktor-cli/internal/app/lang"
	parser "github.com/ktorio/ktor-cli/internal/app/lang/parsers/kotlin"
	"strings"
)

func AddDependency(buildPath, versionKey string) (string, error) {
	input, err := antlr.NewFileStream(buildPath)

	if err != nil {
		return "", err
	}

	lexer := parser.NewKotlinLexer(&input.InputStream)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewKotlinParser(stream)
	rewriter := antlr.NewTokenStreamRewriter(stream)

	sts, ok := findDepsStatements(p.Script())

	if !ok {
		return "", errors.New("could not find dependencies")
	}

	st, _, ok := findKtorDep(sts)

	if !ok {
		return "", errors.New("could not find catalog Ktor dependencies")
	}

	indent := ""
	for _, t := range stream.GetHiddenTokensToLeft(st.GetStart().GetTokenIndex(), antlr.TokenHiddenChannel) {
		indent += t.GetText()
	}

	obj := strings.ReplaceAll(versionKey, "-", ".")
	rewriter.InsertAfterDefault(st.GetStop().GetTokenIndex(), "\n"+indent+fmt.Sprintf("implementation(libs.%s)", obj))

	return rewriter.GetTextDefault(), nil
}

func findKtorDep(depStatements []parser.IStatementContext) (parser.IStatementContext, parser.IValueArgumentContext, bool) {
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
			if strings.HasPrefix(va.GetText(), "libs.ktor") {
				return st, va, true
			}
		}
	}

	return nil, nil, false
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
