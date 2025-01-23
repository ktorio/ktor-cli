package lang

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"path/filepath"
	"strings"
)

type SyntaxError struct {
	Line, Col int
	Msg       string
}

type SyntaxErrors struct {
	SyntaxErrors []SyntaxError
	File         string
}

func (s SyntaxErrors) Error() string {
	var sb strings.Builder

	sep := ""
	for _, e := range s.SyntaxErrors {
		sb.WriteString(sep)
		sb.WriteString(fmt.Sprintf("line %d:%d %s", e.Line, e.Col, e.Msg))
		sep = "\n"
	}

	return fmt.Sprintf("%s: syntax error[s]:\n%s", filepath.Base(s.File), sb.String())
}

type ErrorListener struct {
	*antlr.DefaultErrorListener
	Errors *SyntaxErrors
	File   string
}

func NewErrorListener(file string) *ErrorListener {
	return &ErrorListener{DefaultErrorListener: antlr.NewDefaultErrorListener(), File: file}
}

func (d *ErrorListener) SyntaxError(_ antlr.Recognizer, _ interface{}, line, col int, msg string, _ antlr.RecognitionException) {
	if d.Errors == nil {
		d.Errors = &SyntaxErrors{File: d.File}
	}

	d.Errors.SyntaxErrors = append(d.Errors.SyntaxErrors, SyntaxError{Line: line, Col: col, Msg: msg})
}

var DefaultIndent = strings.Repeat(" ", 4)

func Quote(s string) string {
	if strings.HasPrefix(s, "\"") && strings.HasSuffix(s, "\"") {
		return s
	}

	return "\"" + s + "\""
}

func Unquote(s string) string {
	if strings.HasPrefix(s, `"`) && strings.HasSuffix(s, `"`) {
		runes := []rune(s)
		return string(runes[1 : len(runes)-1])
	}
	return s
}

func FindChild[T any](tree antlr.Tree) (T, bool) {
	var zero T

	for ; tree.GetChildCount() > 0; tree = tree.GetChild(0) {
		if ch, ok := tree.(T); ok {
			return ch, true
		}
	}

	return zero, false
}

func HiddenTokensToLeft(stream *antlr.CommonTokenStream, tokenIndex int) string {
	indent := ""
	for _, t := range stream.GetHiddenTokensToLeft(tokenIndex, antlr.TokenHiddenChannel) {
		indent += t.GetText()
	}
	return indent
}

// ToIndentedStringTree This function is useful for debugging
//
// goland:noinspection GoUnusedFunction
func ToIndentedStringTree(tree antlr.Tree, ruleNames []string, level int) string {
	if tree == nil {
		return ""
	}

	indent := ""
	for i := 0; i < level; i++ {
		indent += "  "
	}

	switch t := tree.(type) {
	case antlr.TerminalNode:
		token := t.GetSymbol()
		return fmt.Sprintf("%sTOKEN: %s\n", indent, token.GetText())
	case antlr.RuleNode:
		ruleName := ruleNames[t.GetRuleContext().GetRuleIndex()]
		result := fmt.Sprintf("%sRULE: %s\n", indent, ruleName)
		for i := 0; i < t.GetChildCount(); i++ {
			result += ToIndentedStringTree(t.GetChild(i), ruleNames, level+1)
		}
		return result
	default:
		return fmt.Sprintf("%sUNKNOWN NODE\n", indent)
	}
}
