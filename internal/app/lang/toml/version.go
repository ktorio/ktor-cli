package toml

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/ktorio/ktor-cli/internal/app/ktor"
	"github.com/ktorio/ktor-cli/internal/app/lang"
	parser "github.com/ktorio/ktor-cli/internal/app/lang/parsers/toml"
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

func ParseToml(fp string) (*Document, error) {
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

	lang.InsertLnAfter(
		rewriter,
		libTable.Element.GetStop(),
		"",
		LibEntry(versionKey, mc),
	)

	return rewriter.GetTextDefault(), nil
}
