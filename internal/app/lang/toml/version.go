package toml

import (
	"errors"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/ktorio/ktor-cli/internal/app/ktor"
	"github.com/ktorio/ktor-cli/internal/app/lang"
	parser "github.com/ktorio/ktor-cli/internal/app/lang/parsers/toml"
	"strings"
)

func AddLib(versionsPath string, mc ktor.MavenCoords) (string, error) {
	input, err := antlr.NewFileStream(versionsPath)

	if err != nil {
		return "", err
	}

	lexer := parser.NewTomlLexer(&input.InputStream)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewTomlParser(stream)
	rewriter := antlr.NewTokenStreamRewriter(stream)

	doc := p.Document()
	libTable, ok := findTable(doc, "libraries")

	if !ok {
		return "", errors.New("unable to find the [libraries] section")
	}

	dep, vr, ok := findKtorDep(doc, libTable)

	if !ok {
		entries, err := findTableEntries(doc, "versions")

		if err != nil {
			return "", err
		}

		key := ""
		for _, e := range entries {
			if kv := e.Key_value(); kv != nil && kv.Key() != nil && strings.HasPrefix(kv.Key().Simple_key().GetText(), "ktor") {
				key = kv.Key().Simple_key().GetText()
				break
			}
		}

		if key == "" {
			return "", errors.New("toml: cannot find Ktor version")
		}

		lib := fmt.Sprintf("%s = { module = \"%s\", version.ref = \"%s\" }", mc.Artifact, mc.String(), key)
		rewriter.InsertAfterDefault(libTable.GetStop().GetTokenIndex(), "\n"+lang.HiddenTokensToLeft(stream, libTable.GetStop().GetTokenIndex())+lib)

		return rewriter.GetTextDefault(), nil
	}

	lib := fmt.Sprintf("%s = { module = \"%s\", version.ref = %s }", mc.Artifact, mc.String(), vr)
	rewriter.InsertAfterDefault(dep.GetStop().GetTokenIndex(), "\n"+lang.HiddenTokensToLeft(stream, dep.GetStart().GetTokenIndex())+lib)

	return rewriter.GetTextDefault(), nil
}

func findTableEntries(doc antlr.ParseTree, table string) ([]parser.IExpressionContext, error) {
	tableIndex := -1

	for i, ch := range doc.GetChildren() {
		if ch.GetChildCount() == 0 {
			continue
		}

		t, ok := ch.GetChild(0).(parser.ITableContext)

		if !ok {
			continue
		}

		if t.GetText() != fmt.Sprintf("[%s]", table) {
			continue
		}

		tableIndex = i
		break
	}

	if tableIndex == -1 {
		return nil, errors.New(fmt.Sprintf("toml: unable to find the [%s] section", table))
	}

	var exprs []parser.IExpressionContext
	for _, ch := range doc.GetChildren()[tableIndex+1:] {
		if ch.GetChildCount() == 0 {
			continue
		}

		if _, ok := ch.GetChild(0).(parser.ITableContext); ok { // next table
			break
		}

		if expr, ok := ch.(parser.IExpressionContext); ok {
			exprs = append(exprs, expr)
		}
	}

	return exprs, nil
}

func findKtorDep(doc parser.IDocumentContext, table parser.ITableContext) (parser.IExpressionContext, string, bool) {
	tableExpr := table.GetParent().(parser.IExpressionContext)
	foundTable := false
	for _, e := range doc.GetChildren() {
		if e == tableExpr {
			foundTable = true
			continue
		}

		if !foundTable {
			continue
		}

		if _, ok := e.GetChild(0).(parser.ITableContext); ok {
			break
		}

		kv, ok := e.GetChild(0).(parser.IKey_valueContext)

		if !ok {
			continue
		}

		vr, ok := findVersionRef(kv)

		if !ok {
			continue
		}

		if strings.HasPrefix(vr, "\"ktor") {
			return e.(parser.IExpressionContext), vr, true
		}
	}

	return nil, "", false
}

func findVersionRef(kv parser.IKey_valueContext) (string, bool) {
	it := kv.Value().Inline_table()

	if it == nil {
		return "", false
	}

	kvs := it.Inline_table_keyvals()

	for vals := kvs.Inline_table_keyvals_non_empty(); vals != nil; vals = vals.Inline_table_keyvals_non_empty() {
		if vals.Key().GetText() == "version.ref" {
			return vals.Value().String_().BASIC_STRING().GetSymbol().GetText(), true
		}
	}

	return "", false
}

func findTable(doc parser.IDocumentContext, name string) (parser.ITableContext, bool) {
	for _, ch := range doc.GetChildren() {
		if ch.GetChildCount() == 0 {
			continue
		}

		t, ok := ch.GetChild(0).(parser.ITableContext)

		if !ok {
			continue
		}

		if t.GetText() == fmt.Sprintf("[%s]", name) {
			return t, true
		}
	}

	return nil, false
}
