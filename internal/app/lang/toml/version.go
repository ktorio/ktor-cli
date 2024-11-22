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

func NewParser(fp string) (*parser.TomlParser, error) {
	input, err := antlr.NewFileStream(fp)

	if err != nil {
		return nil, err
	}

	lexer := parser.NewTomlLexer(&input.InputStream)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	return parser.NewTomlParser(stream), nil
}

func FindCatalogLib(p *parser.TomlParser, mavenCoords ktor.MavenCoords) (string, bool) {
	doc := p.Document()

	libTable, ok := findTable(doc, "libraries")

	if !ok {
		return "", false
	}

	entries, err := findTableEntries(doc, libTable)

	if err != nil {
		return "", false
	}

	for _, e := range entries {
		if kv := e.Key_value(); kv != nil && kv.Value().Inline_table() != nil {
			tkv, ok := findTableKvByKey(kv, "module")
			if !ok {
				continue
			}

			artefact := lang.Unquote(tkv.Value().String_().GetText())

			if !strings.HasPrefix(artefact, "io.ktor") {
				continue
			}

			if mc, ok := ktor.ParseMavenCoords(artefact); ok && mavenCoords.RoughlySame(mc) {
				return lang.Unquote(kv.Key().GetText()), true
			}
		}
	}

	return "", false
}

func AddLib(p *parser.TomlParser, mc ktor.MavenCoords) (string, error) {
	stream := p.GetTokenStream().(*antlr.CommonTokenStream)
	rewriter := antlr.NewTokenStreamRewriter(stream)

	doc := p.Document()
	libTable, ok := findTable(doc, "libraries")

	if !ok {
		return "", errors.New("toml: unable to find the [libraries] section")
	}

	dep, vr, ok := findKtorDep(doc, libTable)

	if !ok {
		versionsTable, ok := findTable(doc, "versions")

		if !ok {
			return "", errors.New("toml: unable to find the [versions] section")
		}

		entries, err := findTableEntries(doc, versionsTable)

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
			key = "ktor"
			v := fmt.Sprintf("%s = \"%s\"", key, mc.Version)
			rewriter.InsertAfterDefault(versionsTable.GetStop().GetTokenIndex(), "\n"+lang.HiddenTokensToLeft(stream, versionsTable.GetStart().GetTokenIndex())+v)
		}

		lib := fmt.Sprintf("%s = { module = \"%s\", version.ref = \"%s\" }", mc.Artifact, mc.String(), key)
		rewriter.InsertAfterDefault(libTable.GetStop().GetTokenIndex(), "\n"+lang.HiddenTokensToLeft(stream, libTable.GetStop().GetTokenIndex())+lib)

		return rewriter.GetTextDefault(), nil
	}

	lib := fmt.Sprintf("%s = { module = \"%s\", version.ref = %s }", mc.Artifact, mc.String(), vr)
	rewriter.InsertAfterDefault(dep.GetStop().GetTokenIndex(), "\n"+lang.HiddenTokensToLeft(stream, dep.GetStart().GetTokenIndex())+lib)

	return rewriter.GetTextDefault(), nil
}

func ruleIndex(tree antlr.Tree) int {
	if tree.GetParent() == nil {
		return -1
	}

	for i, ch := range tree.GetParent().GetChildren() {
		if ch == tree {
			return i
		}
	}

	return -1
}

func findTableEntries(doc antlr.ParseTree, table antlr.Tree) ([]parser.IExpressionContext, error) {
	tableIndex := ruleIndex(table.GetParent())

	if tableIndex == -1 {
		return nil, errors.New(fmt.Sprintf("toml: unable to find the [%s] section", table))
	}

	var exprs []parser.IExpressionContext
	for _, ch := range doc.GetChildren()[tableIndex+1:] {
		if ch.GetChildCount() == 0 {
			continue
		}

		if _, ok := ch.GetChild(0).(parser.ITableContext); ok { // if next table is met
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

func findTableKvByKey(kv parser.IKey_valueContext, key string) (parser.IInline_table_keyvals_non_emptyContext, bool) {
	it := kv.Value().Inline_table()

	if it == nil {
		return nil, false
	}

	kvs := it.Inline_table_keyvals()

	for vals := kvs.Inline_table_keyvals_non_empty(); vals != nil; vals = vals.Inline_table_keyvals_non_empty() {
		if vals.Key().GetText() == key {
			return vals, true
		}
	}

	return nil, false
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

func findTable(doc antlr.Tree, name string) (parser.ITableContext, bool) {
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
