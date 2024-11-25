package gradle

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/ktorio/ktor-cli/internal/app/lang"
	"github.com/ktorio/ktor-cli/internal/app/lang/kotlin"
	parser "github.com/ktorio/ktor-cli/internal/app/lang/parsers/kotlin"
	"strings"
)

type DepKind int

const (
	VersionCatalogDep DepKind = iota
)

type BuildRoot struct {
	Dependencies Dependencies
	Stream       *antlr.CommonTokenStream
	Parser       *parser.KotlinParser
}

type Dependencies struct {
	List       []Dep
	Statements parser.IStatementsContext
}

type Dep struct {
	Kind        DepKind
	IsTest      bool
	IsBom       bool
	CatalogPath string
	Statement   parser.IStatementContext
}

func ParseBuildFile(fp string) (*BuildRoot, error) {
	p, err := kotlin.NewParser(fp)
	root := BuildRoot{}

	if err != nil {
		return &root, err
	}

	root.Stream = p.GetTokenStream().(*antlr.CommonTokenStream)
	root.Parser = p

	for _, st := range p.Script().AllStatement() {
		e, ok := lang.FindChild[parser.IPostfixUnaryExpressionContext](st)

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

		suf, ok := e.GetChild(1).(parser.IPostfixUnarySuffixContext)

		if !ok {
			continue
		}

		lit, ok := lang.FindChild[parser.ILambdaLiteralContext](suf)

		if !ok {
			continue
		}

		root.Dependencies.Statements = lit.Statements()
		for _, depSt := range lit.Statements().AllStatement() {
			pus, ok := lang.FindChild[parser.IPostfixUnaryExpressionContext](depSt)

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

			if pe.SimpleIdentifier().GetText() != "implementation" && pe.SimpleIdentifier().GetText() != "testImplementation" {
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

				d := Dep{
					IsTest:      pe.SimpleIdentifier().GetText() == "testImplementation",
					Kind:        VersionCatalogDep,
					CatalogPath: ps.GetText(),
					Statement:   depSt,
				}
				root.Dependencies.List = append(root.Dependencies.List, d)

				if id := ps.PrimaryExpression().SimpleIdentifier(); id == nil || id.GetText() != "platform" {
					continue
				}

				pus2, ok = ps.GetChild(1).(parser.IPostfixUnarySuffixContext)

				if !ok || pus2.CallSuffix() == nil {
					continue
				}

				for _, va := range findValueArguments(pus2.CallSuffix()) {
					if strings.HasPrefix(lang.Unquote(va.GetText()), "io.ktor:ktor-bom") {
						d.IsBom = true
						break
					}
				}
			}
		}
	}

	return &root, nil
}

func findValueArguments(cf parser.ICallSuffixContext) []parser.IValueArgumentContext {
	vas := cf.ValueArguments()

	if vas == nil {
		return nil
	}

	return vas.AllValueArgument()
}
