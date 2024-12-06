package toml

import (
	"errors"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/ktorio/ktor-cli/internal/app/ktor"
	"github.com/ktorio/ktor-cli/internal/app/lang"
	parser "github.com/ktorio/ktor-cli/internal/app/lang/parsers/toml"
	"github.com/ktorio/ktor-cli/internal/app/utils"
	"path/filepath"
	"strings"
)

type Document struct {
	Stream   *antlr.CommonTokenStream
	Rewriter *antlr.TokenStreamRewriter
	Tables   Tables
}

type Tables struct {
	List []Table
}

type Table struct {
	Name    string
	Element parser.ITableContext
	Entries []TableEntry
}

type TableEntryKind int

const (
	ValueMap TableEntryKind = iota
	StringValue
)

type TableEntry struct {
	Kind       TableEntryKind
	KeyValue   map[string]string
	Key        string
	String     string
	Expression parser.IExpressionContext
}

func (te *TableEntry) Get(key string) (string, bool) {
	if te.Kind != ValueMap {
		return "", false
	}

	v, ok := te.KeyValue[key]
	return v, ok
}

func ParseCatalogToml(projectDir string) (*Document, error) {
	fp, ok := FindVersionsPath(projectDir)

	if !ok {
		return nil, errors.New(fmt.Sprintf("catalog: cannot find TOML file for the project %s", projectDir))
	}

	input, err := antlr.NewFileStream(fp)

	if err != nil {
		return nil, err
	}

	lexer := parser.NewTomlLexer(&input.InputStream)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewTomlParser(stream)

	doc := Document{Stream: stream, Rewriter: antlr.NewTokenStreamRewriter(stream)}

	table := Table{}
	for _, ch := range p.Document().GetChildren() {
		if ch.GetChildCount() == 0 {
			continue
		}

		if _, ok := ch.GetChild(0).(parser.ICommentContext); ok {
			continue
		}

		t, ok := ch.GetChild(0).(parser.ITableContext)

		if ok && table.Name == "" {
			table.Name = tableName(t)
			table.Element = t
			continue
		}

		if !ok {
			table.Entries = append(table.Entries, parseEntry(ch))
		}

		if ok {
			tName := tableName(t)
			if table.Name != tName {
				doc.Tables.List = append(doc.Tables.List, table)
				table = Table{Name: tName, Element: t}
			}
		}
	}

	doc.Tables.List = append(doc.Tables.List, table)

	return &doc, nil
}

func tableName(t parser.ITableContext) string {
	if t.Standard_table() != nil && t.Standard_table().Key() != nil {
		return t.Standard_table().Key().GetText()
	}

	return t.GetText()
}

func parseEntry(tree antlr.Tree) TableEntry {
	exp, ok := tree.(parser.IExpressionContext)
	entry := TableEntry{Expression: exp}

	if !ok {
		return entry
	}

	kv := exp.Key_value()
	if kv == nil {
		return entry
	}

	entry.Key = kv.Key().GetText()

	entry.KeyValue = make(map[string]string)
	if kv.Value().Inline_table() != nil {
		entry.Kind = ValueMap

		it := kv.Value().Inline_table()

		if it == nil {
			return entry
		}

		kvs := it.Inline_table_keyvals()

		for vals := kvs.Inline_table_keyvals_non_empty(); vals != nil; vals = vals.Inline_table_keyvals_non_empty() {
			entry.KeyValue[vals.Key().GetText()] = lang.Unquote(vals.Value().String_().GetText())
		}
		return entry
	}

	entry.Kind = StringValue
	entry.String = lang.Unquote(kv.Value().String_().GetText())

	return entry
}

func AddLib(doc *Document, mc ktor.MavenCoords) (string, error) {
	versionKey := ""
	versionsTable, hasVersionsTable := FindTable(doc, "versions")

	if hasVersionsTable {
		if e, ok := FindVersionPrefixed(doc, "ktor"); ok {
			versionKey = e.Key
		}
	}

	rewriter := doc.Rewriter

	if versionKey == "" && hasVersionsTable {
		versionKey = "ktor"

		lang.InsertLnAfter(
			doc.Rewriter,
			versionsTable.Element.GetStop(),
			lang.HiddenTokensToLeft(doc.Stream, versionsTable.Element.GetStart().GetTokenIndex()),
			VersionEntry(versionKey, mc.Version),
		)
	}

	libTable, hasLibTable := FindTable(doc, "libraries")
	if !hasLibTable {
		if hasVersionsTable && len(versionsTable.Entries) > 0 {
			lastVersion := versionsTable.Entries[len(versionsTable.Entries)-1].Expression
			rewriter.InsertAfterDefault(lastVersion.GetStop().GetTokenIndex(), "\n\n"+NewLibraryTableWithKtor(mc))
		}

		return rewriter.GetTextDefault(), nil
	}

	var firstKtorLib *TableEntry
	isModule := false
	for _, e := range libTable.Entries {
		if e.Kind != ValueMap {
			continue
		}

		if g, ok := e.KeyValue["group"]; ok && g == "io.ktor" {
			firstKtorLib = &e
			break
		}

		if m, ok := e.KeyValue["module"]; ok && strings.HasPrefix(m, "io.ktor") {
			firstKtorLib = &e
			isModule = true
			break
		}
	}

	if firstKtorLib != nil && !isModule {
		lang.InsertLnAfter(rewriter, firstKtorLib.Expression.GetStop(), "", LibEntryGroupName(versionKey, mc))
		return rewriter.GetTextDefault(), nil
	}

	if firstKtorLib != nil {
		lang.InsertLnAfter(rewriter, firstKtorLib.Expression.GetStop(), "", LibEntryModule(versionKey, mc))
		return rewriter.GetTextDefault(), nil
	}

	lang.InsertLnAfter(rewriter, libTable.Element.GetStop(), "", LibEntryModule(versionKey, mc))

	return rewriter.GetTextDefault(), nil
}

func FindVersionsPath(projectDir string) (string, bool) {
	inCurrentDir := filepath.Join(projectDir, "gradle", "libs.versions.toml")

	if utils.Exists(inCurrentDir) {
		return inCurrentDir, true
	}

	if utils.Exists(filepath.Join(projectDir, "..", "gradle", "libs.versions.toml")) {
		return filepath.Join(projectDir, "..", "gradle", "libs.versions.toml"), true
	}

	return inCurrentDir, false
}
