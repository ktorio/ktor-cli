package gradle

import (
	"bytes"
	"errors"
	"github.com/antlr4-go/antlr/v4"
	"github.com/ktorio/ktor-cli/internal/app"
	"github.com/ktorio/ktor-cli/internal/app/lang"
	parser "github.com/ktorio/ktor-cli/internal/app/lang/parsers/kotlin"
	"io"
	"os"
	"strings"
)

type DepKind int

const (
	VersionCatalogDep DepKind = iota
	HardcodedDep
)

type BuildRoot struct {
	Dependencies Dependencies
	Repositories Repositories
	Plugins      Plugins
	TopLevelVars []VarDecl
	Stream       *antlr.CommonTokenStream
	Rewriter     *antlr.TokenStreamRewriter
	Parser       *parser.KotlinParser
}

type VarDecl struct {
	IsDelegate bool
	Delegate   string
	Id         string
	StringVal  string
}

type Repositories struct {
	Statement parser.IStatementContext
}

type Plugins struct {
	List []Plugin
}

type Plugin struct {
	Statement parser.IStatementContext
	Prefix    string
	Id        string
	IsCatalog bool
	Version   string
	Applied   bool
}

type Dependencies struct {
	List       []Dep
	Statements parser.IStatementsContext
}

type Dep struct {
	Kind         DepKind
	IsTest       bool
	IsKtorBom    bool
	Path         string
	PlatformPath string
	Statement    parser.IStatementContext
}

func ParseBuildFile(fp string) (*BuildRoot, error, []lang.SyntaxError) {
	reader, err := fixTrailingNewLine(fp)

	if err != nil {
		var pe *os.PathError
		if errors.As(err, &pe) {
			if errors.Is(pe.Err, os.ErrPermission) {
				err = &app.Error{Err: err, Kind: app.NoPermsForFile}
			}
		}

		return nil, err, nil
	}

	input := antlr.NewIoStream(reader)
	lexer := parser.NewKotlinLexer(input)
	lexer.RemoveErrorListeners()
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewKotlinParser(stream)
	p.RemoveErrorListeners()
	errListener := lang.NewErrorListener()
	p.AddErrorListener(errListener)
	root := BuildRoot{}

	root.Stream = p.GetTokenStream().(*antlr.CommonTokenStream)
	root.Parser = p
	root.Rewriter = antlr.NewTokenStreamRewriter(root.Stream)

	for _, st := range p.Script().AllStatement() {
		if pd, ok := lang.FindChild[parser.IPropertyDeclarationContext](st); ok {
			if vd, ok := parseVarDecl(pd); ok {
				root.TopLevelVars = append(root.TopLevelVars, vd)
			}
		}

		e, ok := lang.FindChild[parser.IPostfixUnaryExpressionContext](st)

		if !ok {
			continue
		}

		id, ok := e.GetChild(0).GetChild(0).(parser.ISimpleIdentifierContext)

		if !ok {
			continue
		}

		if e.GetChildCount() < 2 {
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

		if id.GetText() == "dependencies" {
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

				d := Dep{
					IsTest:    pe.SimpleIdentifier().GetText() == "testImplementation",
					Statement: depSt,
				}

				for _, va := range findValueArguments(pus2.CallSuffix()) {
					ps, ok := lang.FindChild[parser.IPostfixUnaryExpressionContext](va)

					if !ok {
						continue
					}

					k := HardcodedDep
					if strings.HasPrefix(va.GetText(), "libs.") {
						k = VersionCatalogDep
					}

					d.Kind = k
					d.Path = lang.Unquote(va.GetText())

					if id := ps.PrimaryExpression().SimpleIdentifier(); id == nil || id.GetText() != "platform" {
						continue
					}

					pus2, ok = ps.GetChild(1).(parser.IPostfixUnarySuffixContext)

					if !ok || pus2.CallSuffix() == nil {
						continue
					}

					for _, va := range findValueArguments(pus2.CallSuffix()) {
						if pp := lang.Unquote(va.GetText()); strings.HasPrefix(pp, "io.ktor:ktor-bom") {
							d.IsKtorBom = true
							d.PlatformPath = pp
							break
						}
					}
				}

				root.Dependencies.List = append(root.Dependencies.List, d)
			}
		} else if id.GetText() == "plugins" {
			for _, depSt := range lit.Statements().AllStatement() {
				ifc, ok := lang.FindChild[parser.IInfixFunctionCallContext](depSt)

				plugin := Plugin{Statement: depSt, Applied: true}
				if ok && ifc.GetChildCount() > 2 {
					callId, ok := ifc.GetChild(1).(parser.ISimpleIdentifierContext)

					if ok {
						switch callId.GetText() {
						case "version":
							if lit, ok := lang.FindChild[parser.IStringLiteralContext](ifc.GetChild(2)); ok {
								plugin.Version = lang.Unquote(lit.GetText())
							}
						case "apply":
							if lit, ok := lang.FindChild[parser.ILiteralConstantContext](ifc.GetChild(2)); ok {
								plugin.Applied = lit.GetText() != "false"
							}
						}
					}
				}

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

				if pe.SimpleIdentifier().GetText() == "alias" {
					plugin.IsCatalog = true
				}

				if pe.SimpleIdentifier().GetText() == "kotlin" {
					plugin.IsCatalog = false
				}

				plugin.Prefix = pe.SimpleIdentifier().GetText()

				if pus.GetChildCount() < 2 {
					continue
				}

				pus2, ok := pus.GetChild(1).(parser.IPostfixUnarySuffixContext)

				if !ok || pus2.CallSuffix() == nil {
					continue
				}

				for _, va := range findValueArguments(pus2.CallSuffix()) {
					plugin.Id = lang.Unquote(va.GetText())
				}

				root.Plugins.List = append(root.Plugins.List, plugin)
			}
		} else if id.GetText() == "repositories" {
			root.Repositories.Statement = st
		}

	}

	return &root, nil, errListener.Errors
}

func parseVarDecl(pd parser.IPropertyDeclarationContext) (VarDecl, bool) {
	vd := VarDecl{}

	if pd.VariableDeclaration() == nil {
		return vd, false
	}

	vd.Id = pd.VariableDeclaration().SimpleIdentifier().GetText()

	if pd.PropertyDelegate() == nil {
		expr := pd.Expression().GetText()
		if strings.HasPrefix(expr, "\"") && strings.HasSuffix(expr, "\"") {
			vd.StringVal = lang.Unquote(expr)
		}
	} else {
		if id, ok := lang.FindChild[parser.ISimpleIdentifierContext](pd.PropertyDelegate().Expression()); ok {
			vd.IsDelegate = true
			vd.Delegate = id.GetText()
		}
	}

	return vd, true
}

func fixTrailingNewLine(fp string) (io.Reader, error) {
	fBytes, err := os.ReadFile(fp)

	if err != nil {
		return nil, err
	}

	if len(fBytes) > 0 && fBytes[len(fBytes)-1] != '\n' {
		fBytes = append(fBytes, '\n')
	}

	return bytes.NewReader(fBytes), nil
}

func findValueArguments(cf parser.ICallSuffixContext) []parser.IValueArgumentContext {
	vas := cf.ValueArguments()

	if vas == nil {
		return nil
	}

	return vas.AllValueArgument()
}
