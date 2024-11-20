// Code generated from TomlParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // TomlParser

import "github.com/antlr4-go/antlr/v4"

// BaseTomlParserListener is a complete listener for a parse tree produced by TomlParser.
type BaseTomlParserListener struct{}

var _ TomlParserListener = &BaseTomlParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTomlParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTomlParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTomlParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTomlParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterDocument is called when production document is entered.
func (s *BaseTomlParserListener) EnterDocument(ctx *DocumentContext) {}

// ExitDocument is called when production document is exited.
func (s *BaseTomlParserListener) ExitDocument(ctx *DocumentContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseTomlParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseTomlParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterComment is called when production comment is entered.
func (s *BaseTomlParserListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BaseTomlParserListener) ExitComment(ctx *CommentContext) {}

// EnterKey_value is called when production key_value is entered.
func (s *BaseTomlParserListener) EnterKey_value(ctx *Key_valueContext) {}

// ExitKey_value is called when production key_value is exited.
func (s *BaseTomlParserListener) ExitKey_value(ctx *Key_valueContext) {}

// EnterKey is called when production key is entered.
func (s *BaseTomlParserListener) EnterKey(ctx *KeyContext) {}

// ExitKey is called when production key is exited.
func (s *BaseTomlParserListener) ExitKey(ctx *KeyContext) {}

// EnterSimple_key is called when production simple_key is entered.
func (s *BaseTomlParserListener) EnterSimple_key(ctx *Simple_keyContext) {}

// ExitSimple_key is called when production simple_key is exited.
func (s *BaseTomlParserListener) ExitSimple_key(ctx *Simple_keyContext) {}

// EnterUnquoted_key is called when production unquoted_key is entered.
func (s *BaseTomlParserListener) EnterUnquoted_key(ctx *Unquoted_keyContext) {}

// ExitUnquoted_key is called when production unquoted_key is exited.
func (s *BaseTomlParserListener) ExitUnquoted_key(ctx *Unquoted_keyContext) {}

// EnterQuoted_key is called when production quoted_key is entered.
func (s *BaseTomlParserListener) EnterQuoted_key(ctx *Quoted_keyContext) {}

// ExitQuoted_key is called when production quoted_key is exited.
func (s *BaseTomlParserListener) ExitQuoted_key(ctx *Quoted_keyContext) {}

// EnterDotted_key is called when production dotted_key is entered.
func (s *BaseTomlParserListener) EnterDotted_key(ctx *Dotted_keyContext) {}

// ExitDotted_key is called when production dotted_key is exited.
func (s *BaseTomlParserListener) ExitDotted_key(ctx *Dotted_keyContext) {}

// EnterValue is called when production value is entered.
func (s *BaseTomlParserListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseTomlParserListener) ExitValue(ctx *ValueContext) {}

// EnterString is called when production string is entered.
func (s *BaseTomlParserListener) EnterString(ctx *StringContext) {}

// ExitString is called when production string is exited.
func (s *BaseTomlParserListener) ExitString(ctx *StringContext) {}

// EnterInteger is called when production integer is entered.
func (s *BaseTomlParserListener) EnterInteger(ctx *IntegerContext) {}

// ExitInteger is called when production integer is exited.
func (s *BaseTomlParserListener) ExitInteger(ctx *IntegerContext) {}

// EnterFloating_point is called when production floating_point is entered.
func (s *BaseTomlParserListener) EnterFloating_point(ctx *Floating_pointContext) {}

// ExitFloating_point is called when production floating_point is exited.
func (s *BaseTomlParserListener) ExitFloating_point(ctx *Floating_pointContext) {}

// EnterBool_ is called when production bool_ is entered.
func (s *BaseTomlParserListener) EnterBool_(ctx *Bool_Context) {}

// ExitBool_ is called when production bool_ is exited.
func (s *BaseTomlParserListener) ExitBool_(ctx *Bool_Context) {}

// EnterDate_time is called when production date_time is entered.
func (s *BaseTomlParserListener) EnterDate_time(ctx *Date_timeContext) {}

// ExitDate_time is called when production date_time is exited.
func (s *BaseTomlParserListener) ExitDate_time(ctx *Date_timeContext) {}

// EnterArray_ is called when production array_ is entered.
func (s *BaseTomlParserListener) EnterArray_(ctx *Array_Context) {}

// ExitArray_ is called when production array_ is exited.
func (s *BaseTomlParserListener) ExitArray_(ctx *Array_Context) {}

// EnterArray_values is called when production array_values is entered.
func (s *BaseTomlParserListener) EnterArray_values(ctx *Array_valuesContext) {}

// ExitArray_values is called when production array_values is exited.
func (s *BaseTomlParserListener) ExitArray_values(ctx *Array_valuesContext) {}

// EnterComment_or_nl is called when production comment_or_nl is entered.
func (s *BaseTomlParserListener) EnterComment_or_nl(ctx *Comment_or_nlContext) {}

// ExitComment_or_nl is called when production comment_or_nl is exited.
func (s *BaseTomlParserListener) ExitComment_or_nl(ctx *Comment_or_nlContext) {}

// EnterNl_or_comment is called when production nl_or_comment is entered.
func (s *BaseTomlParserListener) EnterNl_or_comment(ctx *Nl_or_commentContext) {}

// ExitNl_or_comment is called when production nl_or_comment is exited.
func (s *BaseTomlParserListener) ExitNl_or_comment(ctx *Nl_or_commentContext) {}

// EnterTable is called when production table is entered.
func (s *BaseTomlParserListener) EnterTable(ctx *TableContext) {}

// ExitTable is called when production table is exited.
func (s *BaseTomlParserListener) ExitTable(ctx *TableContext) {}

// EnterStandard_table is called when production standard_table is entered.
func (s *BaseTomlParserListener) EnterStandard_table(ctx *Standard_tableContext) {}

// ExitStandard_table is called when production standard_table is exited.
func (s *BaseTomlParserListener) ExitStandard_table(ctx *Standard_tableContext) {}

// EnterInline_table is called when production inline_table is entered.
func (s *BaseTomlParserListener) EnterInline_table(ctx *Inline_tableContext) {}

// ExitInline_table is called when production inline_table is exited.
func (s *BaseTomlParserListener) ExitInline_table(ctx *Inline_tableContext) {}

// EnterInline_table_keyvals is called when production inline_table_keyvals is entered.
func (s *BaseTomlParserListener) EnterInline_table_keyvals(ctx *Inline_table_keyvalsContext) {}

// ExitInline_table_keyvals is called when production inline_table_keyvals is exited.
func (s *BaseTomlParserListener) ExitInline_table_keyvals(ctx *Inline_table_keyvalsContext) {}

// EnterInline_table_keyvals_non_empty is called when production inline_table_keyvals_non_empty is entered.
func (s *BaseTomlParserListener) EnterInline_table_keyvals_non_empty(ctx *Inline_table_keyvals_non_emptyContext) {
}

// ExitInline_table_keyvals_non_empty is called when production inline_table_keyvals_non_empty is exited.
func (s *BaseTomlParserListener) ExitInline_table_keyvals_non_empty(ctx *Inline_table_keyvals_non_emptyContext) {
}

// EnterArray_table is called when production array_table is entered.
func (s *BaseTomlParserListener) EnterArray_table(ctx *Array_tableContext) {}

// ExitArray_table is called when production array_table is exited.
func (s *BaseTomlParserListener) ExitArray_table(ctx *Array_tableContext) {}
