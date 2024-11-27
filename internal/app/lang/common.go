package lang

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"strings"
)

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

func FindParent[T any](tree antlr.Tree) (T, bool) {
	var zero T

	for ; tree != nil; tree = tree.GetParent() {
		if expr, ok := tree.(T); ok {
			return expr, true
		}
	}

	return zero, false
}

func ChildrenOfType[T any](tree antlr.Tree) []T {
	var result []T
	for _, ch := range tree.GetChildren() {
		if x, ok := ch.(T); ok {
			result = append(result, x)
		}
	}

	return result
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
