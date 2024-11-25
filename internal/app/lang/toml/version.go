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

type Document struct {
	Stream *antlr.CommonTokenStream
	Tables Tables
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
	Kind     TableEntryKind
	KeyValue map[string]string
	Key      string
	String   string
}

func ParseToml(fp string) (*Document, error) {
	p, err := NewParser(fp)

	if err != nil {
		return nil, err
	}

	doc := Document{Stream: p.GetTokenStream().(*antlr.CommonTokenStream)}

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
	entry := TableEntry{}

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

func NewParser(fp string) (*parser.TomlParser, error) {
	input, err := antlr.NewFileStream(fp)

	if err != nil {
		return nil, err
	}

	lexer := parser.NewTomlLexer(&input.InputStream)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	return parser.NewTomlParser(stream), nil
}

func FindCatalogLib(doc *Document, mavenCoords ktor.MavenCoords) (string, bool) {
	for _, t := range doc.Tables.List {
		for _, e := range t.Entries {
			if e.Kind != ValueMap {
				continue
			}
			m, ok := e.KeyValue["module"]

			if !ok || !strings.HasPrefix(m, "io.ktor") {
				return m, true
			}

			if mc, ok := ktor.ParseMavenCoords(m); ok && mavenCoords.RoughlySame(mc) {
				return e.Key, true
			}
		}
	}

	return "", false
}

func AddLib(doc *Document, mc ktor.MavenCoords) (string, error) {
	key := ""
	var versionsTable parser.ITableContext
	var libTable parser.ITableContext
	stream := doc.Stream

	for _, t := range doc.Tables.List {
		if t.Name == "versions" {
			versionsTable = t.Element
		}

		if t.Name == "libraries" {
			libTable = t.Element
		}

		for _, e := range t.Entries {
			if e.Kind == StringValue && strings.HasPrefix(e.Key, "ktor") {
				key = e.Key
				break
			}
		}
	}

	rewriter := antlr.NewTokenStreamRewriter(stream)
	if key == "" && versionsTable != nil {
		key = "ktor"
		v := fmt.Sprintf("%s = \"%s\"", key, mc.Version)
		rewriter.InsertAfterDefault(versionsTable.GetStop().GetTokenIndex(), "\n"+lang.HiddenTokensToLeft(stream, versionsTable.GetStart().GetTokenIndex())+v)
	}

	if libTable == nil {
		return "", errors.New("toml: unable to find [libraries] section")
	}

	lib := fmt.Sprintf("%s = { module = \"%s\", version.ref = \"%s\" }", mc.Artifact, mc.String(), key)
	rewriter.InsertAfterDefault(libTable.GetStop().GetTokenIndex(), "\n"+lang.HiddenTokensToLeft(stream, libTable.GetStop().GetTokenIndex())+lib)

	return rewriter.GetTextDefault(), nil
}
