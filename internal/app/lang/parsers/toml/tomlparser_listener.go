// Code generated from TomlParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // TomlParser

import "github.com/antlr4-go/antlr/v4"

// TomlParserListener is a complete listener for a parse tree produced by TomlParser.
type TomlParserListener interface {
	antlr.ParseTreeListener

	// EnterDocument is called when entering the document production.
	EnterDocument(c *DocumentContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterComment is called when entering the comment production.
	EnterComment(c *CommentContext)

	// EnterKey_value is called when entering the key_value production.
	EnterKey_value(c *Key_valueContext)

	// EnterKey is called when entering the key production.
	EnterKey(c *KeyContext)

	// EnterSimple_key is called when entering the simple_key production.
	EnterSimple_key(c *Simple_keyContext)

	// EnterUnquoted_key is called when entering the unquoted_key production.
	EnterUnquoted_key(c *Unquoted_keyContext)

	// EnterQuoted_key is called when entering the quoted_key production.
	EnterQuoted_key(c *Quoted_keyContext)

	// EnterDotted_key is called when entering the dotted_key production.
	EnterDotted_key(c *Dotted_keyContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterString is called when entering the string production.
	EnterString(c *StringContext)

	// EnterInteger is called when entering the integer production.
	EnterInteger(c *IntegerContext)

	// EnterFloating_point is called when entering the floating_point production.
	EnterFloating_point(c *Floating_pointContext)

	// EnterBool_ is called when entering the bool_ production.
	EnterBool_(c *Bool_Context)

	// EnterDate_time is called when entering the date_time production.
	EnterDate_time(c *Date_timeContext)

	// EnterArray_ is called when entering the array_ production.
	EnterArray_(c *Array_Context)

	// EnterArray_values is called when entering the array_values production.
	EnterArray_values(c *Array_valuesContext)

	// EnterComment_or_nl is called when entering the comment_or_nl production.
	EnterComment_or_nl(c *Comment_or_nlContext)

	// EnterNl_or_comment is called when entering the nl_or_comment production.
	EnterNl_or_comment(c *Nl_or_commentContext)

	// EnterTable is called when entering the table production.
	EnterTable(c *TableContext)

	// EnterStandard_table is called when entering the standard_table production.
	EnterStandard_table(c *Standard_tableContext)

	// EnterInline_table is called when entering the inline_table production.
	EnterInline_table(c *Inline_tableContext)

	// EnterInline_table_keyvals is called when entering the inline_table_keyvals production.
	EnterInline_table_keyvals(c *Inline_table_keyvalsContext)

	// EnterInline_table_keyvals_non_empty is called when entering the inline_table_keyvals_non_empty production.
	EnterInline_table_keyvals_non_empty(c *Inline_table_keyvals_non_emptyContext)

	// EnterArray_table is called when entering the array_table production.
	EnterArray_table(c *Array_tableContext)

	// ExitDocument is called when exiting the document production.
	ExitDocument(c *DocumentContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitComment is called when exiting the comment production.
	ExitComment(c *CommentContext)

	// ExitKey_value is called when exiting the key_value production.
	ExitKey_value(c *Key_valueContext)

	// ExitKey is called when exiting the key production.
	ExitKey(c *KeyContext)

	// ExitSimple_key is called when exiting the simple_key production.
	ExitSimple_key(c *Simple_keyContext)

	// ExitUnquoted_key is called when exiting the unquoted_key production.
	ExitUnquoted_key(c *Unquoted_keyContext)

	// ExitQuoted_key is called when exiting the quoted_key production.
	ExitQuoted_key(c *Quoted_keyContext)

	// ExitDotted_key is called when exiting the dotted_key production.
	ExitDotted_key(c *Dotted_keyContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitString is called when exiting the string production.
	ExitString(c *StringContext)

	// ExitInteger is called when exiting the integer production.
	ExitInteger(c *IntegerContext)

	// ExitFloating_point is called when exiting the floating_point production.
	ExitFloating_point(c *Floating_pointContext)

	// ExitBool_ is called when exiting the bool_ production.
	ExitBool_(c *Bool_Context)

	// ExitDate_time is called when exiting the date_time production.
	ExitDate_time(c *Date_timeContext)

	// ExitArray_ is called when exiting the array_ production.
	ExitArray_(c *Array_Context)

	// ExitArray_values is called when exiting the array_values production.
	ExitArray_values(c *Array_valuesContext)

	// ExitComment_or_nl is called when exiting the comment_or_nl production.
	ExitComment_or_nl(c *Comment_or_nlContext)

	// ExitNl_or_comment is called when exiting the nl_or_comment production.
	ExitNl_or_comment(c *Nl_or_commentContext)

	// ExitTable is called when exiting the table production.
	ExitTable(c *TableContext)

	// ExitStandard_table is called when exiting the standard_table production.
	ExitStandard_table(c *Standard_tableContext)

	// ExitInline_table is called when exiting the inline_table production.
	ExitInline_table(c *Inline_tableContext)

	// ExitInline_table_keyvals is called when exiting the inline_table_keyvals production.
	ExitInline_table_keyvals(c *Inline_table_keyvalsContext)

	// ExitInline_table_keyvals_non_empty is called when exiting the inline_table_keyvals_non_empty production.
	ExitInline_table_keyvals_non_empty(c *Inline_table_keyvals_non_emptyContext)

	// ExitArray_table is called when exiting the array_table production.
	ExitArray_table(c *Array_tableContext)
}
