// Code generated from TomlParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // TomlParser

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type TomlParser struct {
	*antlr.BaseParser
}

var TomlParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func tomlparserParserInit() {
	staticData := &TomlParserParserStaticData
	staticData.LiteralNames = []string{
		"", "", "", "", "'['", "'[['", "']'", "']]'", "'='", "'.'", "','", "",
		"", "", "", "'{'", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "'}'",
	}
	staticData.SymbolicNames = []string{
		"", "WS", "NL", "COMMENT", "L_BRACKET", "DOUBLE_L_BRACKET", "R_BRACKET",
		"DOUBLE_R_BRACKET", "EQUALS", "DOT", "COMMA", "BASIC_STRING", "LITERAL_STRING",
		"UNQUOTED_KEY", "VALUE_WS", "L_BRACE", "BOOLEAN", "ML_BASIC_STRING",
		"ML_LITERAL_STRING", "FLOAT", "INF", "NAN", "DEC_INT", "HEX_INT", "OCT_INT",
		"BIN_INT", "OFFSET_DATE_TIME", "LOCAL_DATE_TIME", "LOCAL_DATE", "LOCAL_TIME",
		"INLINE_TABLE_WS", "R_BRACE", "ARRAY_WS",
	}
	staticData.RuleNames = []string{
		"document", "expression", "comment", "key_value", "key", "simple_key",
		"unquoted_key", "quoted_key", "dotted_key", "value", "string", "integer",
		"floating_point", "bool_", "date_time", "array_", "array_values", "comment_or_nl",
		"nl_or_comment", "table", "standard_table", "inline_table", "inline_table_keyvals",
		"inline_table_keyvals_non_empty", "array_table",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 32, 181, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 1, 0, 1, 0, 1, 0,
		5, 0, 54, 8, 0, 10, 0, 12, 0, 57, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 3, 1, 68, 8, 1, 1, 2, 3, 2, 71, 8, 2, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 4, 1, 4, 3, 4, 79, 8, 4, 1, 5, 1, 5, 3, 5, 83, 8, 5, 1,
		6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 4, 8, 92, 8, 8, 11, 8, 12, 8, 93,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 103, 8, 9, 1, 10, 1, 10,
		1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 3,
		15, 117, 8, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 133, 8, 16, 3, 16, 135,
		8, 16, 1, 17, 3, 17, 138, 8, 17, 1, 17, 5, 17, 141, 8, 17, 10, 17, 12,
		17, 144, 9, 17, 1, 18, 1, 18, 3, 18, 148, 8, 18, 5, 18, 150, 8, 18, 10,
		18, 12, 18, 153, 9, 18, 1, 19, 1, 19, 3, 19, 157, 8, 19, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 3, 22, 168, 8, 22, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 3, 23, 175, 8, 23, 1, 24, 1, 24, 1, 24,
		1, 24, 1, 24, 0, 0, 25, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
		26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 0, 5, 1, 0, 11, 12, 2,
		0, 11, 12, 17, 18, 1, 0, 22, 25, 1, 0, 19, 21, 1, 0, 26, 29, 178, 0, 50,
		1, 0, 0, 0, 2, 67, 1, 0, 0, 0, 4, 70, 1, 0, 0, 0, 6, 72, 1, 0, 0, 0, 8,
		78, 1, 0, 0, 0, 10, 82, 1, 0, 0, 0, 12, 84, 1, 0, 0, 0, 14, 86, 1, 0, 0,
		0, 16, 88, 1, 0, 0, 0, 18, 102, 1, 0, 0, 0, 20, 104, 1, 0, 0, 0, 22, 106,
		1, 0, 0, 0, 24, 108, 1, 0, 0, 0, 26, 110, 1, 0, 0, 0, 28, 112, 1, 0, 0,
		0, 30, 114, 1, 0, 0, 0, 32, 134, 1, 0, 0, 0, 34, 142, 1, 0, 0, 0, 36, 151,
		1, 0, 0, 0, 38, 156, 1, 0, 0, 0, 40, 158, 1, 0, 0, 0, 42, 162, 1, 0, 0,
		0, 44, 167, 1, 0, 0, 0, 46, 169, 1, 0, 0, 0, 48, 176, 1, 0, 0, 0, 50, 55,
		3, 2, 1, 0, 51, 52, 5, 2, 0, 0, 52, 54, 3, 2, 1, 0, 53, 51, 1, 0, 0, 0,
		54, 57, 1, 0, 0, 0, 55, 53, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56, 58, 1,
		0, 0, 0, 57, 55, 1, 0, 0, 0, 58, 59, 5, 0, 0, 1, 59, 1, 1, 0, 0, 0, 60,
		61, 3, 6, 3, 0, 61, 62, 3, 4, 2, 0, 62, 68, 1, 0, 0, 0, 63, 64, 3, 38,
		19, 0, 64, 65, 3, 4, 2, 0, 65, 68, 1, 0, 0, 0, 66, 68, 3, 4, 2, 0, 67,
		60, 1, 0, 0, 0, 67, 63, 1, 0, 0, 0, 67, 66, 1, 0, 0, 0, 68, 3, 1, 0, 0,
		0, 69, 71, 5, 3, 0, 0, 70, 69, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 5, 1,
		0, 0, 0, 72, 73, 3, 8, 4, 0, 73, 74, 5, 8, 0, 0, 74, 75, 3, 18, 9, 0, 75,
		7, 1, 0, 0, 0, 76, 79, 3, 10, 5, 0, 77, 79, 3, 16, 8, 0, 78, 76, 1, 0,
		0, 0, 78, 77, 1, 0, 0, 0, 79, 9, 1, 0, 0, 0, 80, 83, 3, 14, 7, 0, 81, 83,
		3, 12, 6, 0, 82, 80, 1, 0, 0, 0, 82, 81, 1, 0, 0, 0, 83, 11, 1, 0, 0, 0,
		84, 85, 5, 13, 0, 0, 85, 13, 1, 0, 0, 0, 86, 87, 7, 0, 0, 0, 87, 15, 1,
		0, 0, 0, 88, 91, 3, 10, 5, 0, 89, 90, 5, 9, 0, 0, 90, 92, 3, 10, 5, 0,
		91, 89, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 91, 1, 0, 0, 0, 93, 94, 1,
		0, 0, 0, 94, 17, 1, 0, 0, 0, 95, 103, 3, 20, 10, 0, 96, 103, 3, 22, 11,
		0, 97, 103, 3, 24, 12, 0, 98, 103, 3, 26, 13, 0, 99, 103, 3, 28, 14, 0,
		100, 103, 3, 30, 15, 0, 101, 103, 3, 42, 21, 0, 102, 95, 1, 0, 0, 0, 102,
		96, 1, 0, 0, 0, 102, 97, 1, 0, 0, 0, 102, 98, 1, 0, 0, 0, 102, 99, 1, 0,
		0, 0, 102, 100, 1, 0, 0, 0, 102, 101, 1, 0, 0, 0, 103, 19, 1, 0, 0, 0,
		104, 105, 7, 1, 0, 0, 105, 21, 1, 0, 0, 0, 106, 107, 7, 2, 0, 0, 107, 23,
		1, 0, 0, 0, 108, 109, 7, 3, 0, 0, 109, 25, 1, 0, 0, 0, 110, 111, 5, 16,
		0, 0, 111, 27, 1, 0, 0, 0, 112, 113, 7, 4, 0, 0, 113, 29, 1, 0, 0, 0, 114,
		116, 5, 4, 0, 0, 115, 117, 3, 32, 16, 0, 116, 115, 1, 0, 0, 0, 116, 117,
		1, 0, 0, 0, 117, 118, 1, 0, 0, 0, 118, 119, 3, 34, 17, 0, 119, 120, 5,
		6, 0, 0, 120, 31, 1, 0, 0, 0, 121, 122, 3, 34, 17, 0, 122, 123, 3, 18,
		9, 0, 123, 124, 3, 36, 18, 0, 124, 125, 5, 10, 0, 0, 125, 126, 3, 32, 16,
		0, 126, 127, 3, 34, 17, 0, 127, 135, 1, 0, 0, 0, 128, 129, 3, 34, 17, 0,
		129, 130, 3, 18, 9, 0, 130, 132, 3, 36, 18, 0, 131, 133, 5, 10, 0, 0, 132,
		131, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 135, 1, 0, 0, 0, 134, 121,
		1, 0, 0, 0, 134, 128, 1, 0, 0, 0, 135, 33, 1, 0, 0, 0, 136, 138, 5, 3,
		0, 0, 137, 136, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0,
		139, 141, 5, 2, 0, 0, 140, 137, 1, 0, 0, 0, 141, 144, 1, 0, 0, 0, 142,
		140, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 35, 1, 0, 0, 0, 144, 142, 1,
		0, 0, 0, 145, 147, 5, 2, 0, 0, 146, 148, 5, 3, 0, 0, 147, 146, 1, 0, 0,
		0, 147, 148, 1, 0, 0, 0, 148, 150, 1, 0, 0, 0, 149, 145, 1, 0, 0, 0, 150,
		153, 1, 0, 0, 0, 151, 149, 1, 0, 0, 0, 151, 152, 1, 0, 0, 0, 152, 37, 1,
		0, 0, 0, 153, 151, 1, 0, 0, 0, 154, 157, 3, 40, 20, 0, 155, 157, 3, 48,
		24, 0, 156, 154, 1, 0, 0, 0, 156, 155, 1, 0, 0, 0, 157, 39, 1, 0, 0, 0,
		158, 159, 5, 4, 0, 0, 159, 160, 3, 8, 4, 0, 160, 161, 5, 6, 0, 0, 161,
		41, 1, 0, 0, 0, 162, 163, 5, 15, 0, 0, 163, 164, 3, 44, 22, 0, 164, 165,
		5, 31, 0, 0, 165, 43, 1, 0, 0, 0, 166, 168, 3, 46, 23, 0, 167, 166, 1,
		0, 0, 0, 167, 168, 1, 0, 0, 0, 168, 45, 1, 0, 0, 0, 169, 170, 3, 8, 4,
		0, 170, 171, 5, 8, 0, 0, 171, 174, 3, 18, 9, 0, 172, 173, 5, 10, 0, 0,
		173, 175, 3, 46, 23, 0, 174, 172, 1, 0, 0, 0, 174, 175, 1, 0, 0, 0, 175,
		47, 1, 0, 0, 0, 176, 177, 5, 5, 0, 0, 177, 178, 3, 8, 4, 0, 178, 179, 5,
		7, 0, 0, 179, 49, 1, 0, 0, 0, 17, 55, 67, 70, 78, 82, 93, 102, 116, 132,
		134, 137, 142, 147, 151, 156, 167, 174,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// TomlParserInit initializes any static state used to implement TomlParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewTomlParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func TomlParserInit() {
	staticData := &TomlParserParserStaticData
	staticData.once.Do(tomlparserParserInit)
}

// NewTomlParser produces a new parser instance for the optional input antlr.TokenStream.
func NewTomlParser(input antlr.TokenStream) *TomlParser {
	TomlParserInit()
	this := new(TomlParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &TomlParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "TomlParser.g4"

	return this
}

// TomlParser tokens.
const (
	TomlParserEOF               = antlr.TokenEOF
	TomlParserWS                = 1
	TomlParserNL                = 2
	TomlParserCOMMENT           = 3
	TomlParserL_BRACKET         = 4
	TomlParserDOUBLE_L_BRACKET  = 5
	TomlParserR_BRACKET         = 6
	TomlParserDOUBLE_R_BRACKET  = 7
	TomlParserEQUALS            = 8
	TomlParserDOT               = 9
	TomlParserCOMMA             = 10
	TomlParserBASIC_STRING      = 11
	TomlParserLITERAL_STRING    = 12
	TomlParserUNQUOTED_KEY      = 13
	TomlParserVALUE_WS          = 14
	TomlParserL_BRACE           = 15
	TomlParserBOOLEAN           = 16
	TomlParserML_BASIC_STRING   = 17
	TomlParserML_LITERAL_STRING = 18
	TomlParserFLOAT             = 19
	TomlParserINF               = 20
	TomlParserNAN               = 21
	TomlParserDEC_INT           = 22
	TomlParserHEX_INT           = 23
	TomlParserOCT_INT           = 24
	TomlParserBIN_INT           = 25
	TomlParserOFFSET_DATE_TIME  = 26
	TomlParserLOCAL_DATE_TIME   = 27
	TomlParserLOCAL_DATE        = 28
	TomlParserLOCAL_TIME        = 29
	TomlParserINLINE_TABLE_WS   = 30
	TomlParserR_BRACE           = 31
	TomlParserARRAY_WS          = 32
)

// TomlParser rules.
const (
	TomlParserRULE_document                       = 0
	TomlParserRULE_expression                     = 1
	TomlParserRULE_comment                        = 2
	TomlParserRULE_key_value                      = 3
	TomlParserRULE_key                            = 4
	TomlParserRULE_simple_key                     = 5
	TomlParserRULE_unquoted_key                   = 6
	TomlParserRULE_quoted_key                     = 7
	TomlParserRULE_dotted_key                     = 8
	TomlParserRULE_value                          = 9
	TomlParserRULE_string                         = 10
	TomlParserRULE_integer                        = 11
	TomlParserRULE_floating_point                 = 12
	TomlParserRULE_bool_                          = 13
	TomlParserRULE_date_time                      = 14
	TomlParserRULE_array_                         = 15
	TomlParserRULE_array_values                   = 16
	TomlParserRULE_comment_or_nl                  = 17
	TomlParserRULE_nl_or_comment                  = 18
	TomlParserRULE_table                          = 19
	TomlParserRULE_standard_table                 = 20
	TomlParserRULE_inline_table                   = 21
	TomlParserRULE_inline_table_keyvals           = 22
	TomlParserRULE_inline_table_keyvals_non_empty = 23
	TomlParserRULE_array_table                    = 24
)

// IDocumentContext is an interface to support dynamic dispatch.
type IDocumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	EOF() antlr.TerminalNode
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode

	// IsDocumentContext differentiates from other interfaces.
	IsDocumentContext()
}

type DocumentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDocumentContext() *DocumentContext {
	var p = new(DocumentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TomlParserRULE_document
	return p
}

func InitEmptyDocumentContext(p *DocumentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TomlParserRULE_document
}

func (*DocumentContext) IsDocumentContext() {}

func NewDocumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DocumentContext {
	var p = new(DocumentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TomlParserRULE_document

	return p
}

func (s *DocumentContext) GetParser() antlr.Parser { return s.parser }

func (s *DocumentContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *DocumentContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DocumentContext) EOF() antlr.TerminalNode {
	return s.GetToken(TomlParserEOF, 0)
}

func (s *DocumentContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(TomlParserNL)
}

func (s *DocumentContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(TomlParserNL, i)
}

func (s *DocumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DocumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DocumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlParserListener); ok {
		listenerT.EnterDocument(s)
	}
}

func (s *DocumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlParserListener); ok {
		listenerT.ExitDocument(s)
	}
}

func (p *TomlParser) Document() (localctx IDocumentContext) {
	localctx = NewDocumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TomlParserRULE_document)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(50)
		p.Expression()
	}
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == TomlParserNL {
		{
			p.SetState(51)
			p.Match(TomlParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(52)
			p.Expression()
		}

		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(58)
		p.Match(TomlParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Key_value() IKey_valueContext
	Comment() ICommentContext
	Table() ITableContext

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TomlParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TomlParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TomlParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Key_value() IKey_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKey_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKey_valueContext)
}

func (s *ExpressionContext) Comment() ICommentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommentContext)
}

func (s *ExpressionContext) Table() ITableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITableContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlParserListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlParserListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *TomlParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TomlParserRULE_expression)
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TomlParserBASIC_STRING, TomlParserLITERAL_STRING, TomlParserUNQUOTED_KEY:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(60)
			p.Key_value()
		}
		{
			p.SetState(61)
			p.Comment()
		}

	case TomlParserL_BRACKET, TomlParserDOUBLE_L_BRACKET:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(63)
			p.Table()
		}
		{
			p.SetState(64)
			p.Comment()
		}

	case TomlParserEOF, TomlParserNL, TomlParserCOMMENT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(66)
			p.Comment()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICommentContext is an interface to support dynamic dispatch.
type ICommentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COMMENT() antlr.TerminalNode

	// IsCommentContext differentiates from other interfaces.
	IsCommentContext()
}

type CommentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommentContext() *CommentContext {
	var p = new(CommentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TomlParserRULE_comment
	return p
}

func InitEmptyCommentContext(p *CommentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TomlParserRULE_comment
}

func (*CommentContext) IsCommentContext() {}

func NewCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentContext {
	var p = new(CommentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TomlParserRULE_comment

	return p
}

func (s *CommentContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentContext) COMMENT() antlr.TerminalNode {
	return s.GetToken(TomlParserCOMMENT, 0)
}

func (s *CommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlParserListener); ok {
		listenerT.EnterComment(s)
	}
}

func (s *CommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlParserListener); ok {
		listenerT.ExitComment(s)
	}
}

func (p *TomlParser) Comment() (localctx ICommentContext) {
	localctx = NewCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, TomlParserRULE_comment)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TomlParserCOMMENT {
		{
			p.SetState(69)
			p.Match(TomlParserCOMMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IKey_valueContext is an interface to support dynamic dispatch.
type IKey_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Key() IKeyContext
	EQUALS() antlr.TerminalNode
	Value() IValueContext

	// IsKey_valueContext differentiates from other interfaces.
	IsKey_valueContext()
}

type Key_valueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKey_valueContext() *Key_valueContext {
	var p = new(Key_valueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TomlParserRULE_key_value
	return p
}

func InitEmptyKey_valueContext(p *Key_valueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TomlParserRULE_key_value
}

func (*Key_valueContext) IsKey_valueContext() {}

func NewKey_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Key_valueContext {
	var p = new(Key_valueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TomlParserRULE_key_value

	return p
}

func (s *Key_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Key_valueContext) Key() IKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeyContext)
}

func (s *Key_valueContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(TomlParserEQUALS, 0)
}

func (s *Key_valueContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *Key_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Key_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Key_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlParserListener); ok {
		listenerT.EnterKey_value(s)
	}
}

func (s *Key_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlParserListener); ok {
		listenerT.ExitKey_value(s)
	}
}

func (p *TomlParser) Key_value() (localctx IKey_valueContext) {
	localctx = NewKey_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, TomlParserRULE_key_value)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)
		p.Key()
	}
	{
		p.SetState(73)
		p.Match(TomlParserEQUALS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(74)
		p.Value()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IKeyContext is an interface to support dynamic dispatch.
type IKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Simple_key() ISimple_keyContext
	Dotted_key() IDotted_keyContext

	// IsKeyContext differentiates from other interfaces.
	IsKeyContext()
}

type KeyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyContext() *KeyContext {
	var p = new(KeyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TomlParserRULE_key
	return p
}

func InitEmptyKeyContext(p *KeyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TomlParserRULE_key
}

func (*KeyContext) IsKeyContext() {}

func NewKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyContext {
	var p = new(KeyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TomlParserRULE_key

	return p
}

func (s *KeyContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyContext) Simple_key() ISimple_keyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimple_keyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimple_keyContext)
}

func (s *KeyContext) Dotted_key() IDotted_keyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDotted_keyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDotted_keyContext)
}

func (s *KeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlParserListener); ok {
		listenerT.EnterKey(s)
	}
}

func (s *KeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlParserListener); ok {
		listenerT.ExitKey(s)
	}
}

func (p *TomlParser) Key() (localctx IKeyContext) {
	localctx = NewKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, TomlParserRULE_key)
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(76)
			p.Simple_key()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(77)
			p.Dotted_key()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISimple_keyContext is an interface to support dynamic dispatch.
type ISimple_keyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Quoted_key() IQuoted_keyContext
	Unquoted_key() IUnquoted_keyContext

	// IsSimple_keyContext differentiates from other interfaces.
	IsSimple_keyContext()
}

type Simple_keyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimple_keyContext() *Simple_keyContext {
	var p = new(Simple_keyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TomlParserRULE_simple_key
	return p
}

func InitEmptySimple_keyContext(p *Simple_keyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TomlParserRULE_simple_key
}

func (*Simple_keyContext) IsSimple_keyContext() {}

func NewSimple_keyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Simple_keyContext {
	var p = new(Simple_keyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TomlParserRULE_simple_key

	return p
}

func (s *Simple_keyContext) GetParser() antlr.Parser { return s.parser }

func (s *Simple_keyContext) Quoted_key() IQuoted_keyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQuoted_keyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQuoted_keyContext)
}

func (s *Simple_keyContext) Unquoted_key() IUnquoted_keyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnquoted_keyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnquoted_keyContext)
}

func (s *Simple_keyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Simple_keyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Simple_keyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlParserListener); ok {
		listenerT.EnterSimple_key(s)
	}
}

func (s *Simple_keyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlParserListener); ok {
		listenerT.ExitSimple_key(s)
	}
}

func (p *TomlParser) Simple_key() (localctx ISimple_keyContext) {
	localctx = NewSimple_keyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, TomlParserRULE_simple_key)
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TomlParserBASIC_STRING, TomlParserLITERAL_STRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(80)
			p.Quoted_key()
		}

	case TomlParserUNQUOTED_KEY:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(81)
			p.Unquoted_key()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnquoted_keyContext is an interface to support dynamic dispatch.
type IUnquoted_keyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UNQUOTED_KEY() antlr.TerminalNode

	// IsUnquoted_keyContext differentiates from other interfaces.
	IsUnquoted_keyContext()
}

type Unquoted_keyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnquoted_keyContext() *Unquoted_keyContext {
	var p = new(Unquoted_keyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TomlParserRULE_unquoted_key
	return p
}

func InitEmptyUnquoted_keyContext(p *Unquoted_keyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TomlParserRULE_unquoted_key
}

func (*Unquoted_keyContext) IsUnquoted_keyContext() {}

func NewUnquoted_keyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Unquoted_keyContext {
	var p = new(Unquoted_keyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TomlParserRULE_unquoted_key

	return p
}

func (s *Unquoted_keyContext) GetParser() antlr.Parser { return s.parser }

func (s *Unquoted_keyContext) UNQUOTED_KEY() antlr.TerminalNode {
	return s.GetToken(TomlParserUNQUOTED_KEY, 0)
}

func (s *Unquoted_keyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Unquoted_keyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Unquoted_keyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlParserListener); ok {
		listenerT.EnterUnquoted_key(s)
	}
}

func (s *Unquoted_keyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlParserListener); ok {
		listenerT.ExitUnquoted_key(s)
	}
}

func (p *TomlParser) Unquoted_key() (localctx IUnquoted_keyContext) {
	localctx = NewUnquoted_keyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, TomlParserRULE_unquoted_key)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(84)
		p.Match(TomlParserUNQUOTED_KEY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IQuoted_keyContext is an interface to support dynamic dispatch.
type IQuoted_keyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BASIC_STRING() antlr.TerminalNode
	LITERAL_STRING() antlr.TerminalNode

	// IsQuoted_keyContext differentiates from other interfaces.
	IsQuoted_keyContext()
}

type Quoted_keyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuoted_keyContext() *Quoted_keyContext {
	var p = new(Quoted_keyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TomlParserRULE_quoted_key
	return p
}

func InitEmptyQuoted_keyContext(p *Quoted_keyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TomlParserRULE_quoted_key
}

func (*Quoted_keyContext) IsQuoted_keyContext() {}

func NewQuoted_keyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Quoted_keyContext {
	var p = new(Quoted_keyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TomlParserRULE_quoted_key

	return p
}

func (s *Quoted_keyContext) GetParser() antlr.Parser { return s.parser }

func (s *Quoted_keyContext) BASIC_STRING() antlr.TerminalNode {
	return s.GetToken(TomlParserBASIC_STRING, 0)
}

func (s *Quoted_keyContext) LITERAL_STRING() antlr.TerminalNode {
	return s.GetToken(TomlParserLITERAL_STRING, 0)
}

func (s *Quoted_keyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Quoted_keyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Quoted_keyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlParserListener); ok {
		listenerT.EnterQuoted_key(s)
	}
}

func (s *Quoted_keyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlParserListener); ok {
		listenerT.ExitQuoted_key(s)
	}
}

func (p *TomlParser) Quoted_key() (localctx IQuoted_keyContext) {
	localctx = NewQuoted_keyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, TomlParserRULE_quoted_key)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(86)
		_la = p.GetTokenStream().LA(1)

		if !(_la == TomlParserBASIC_STRING || _la == TomlParserLITERAL_STRING) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDotted_keyContext is an interface to support dynamic dispatch.
type IDotted_keyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSimple_key() []ISimple_keyContext
	Simple_key(i int) ISimple_keyContext
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode

	// IsDotted_keyContext differentiates from other interfaces.
	IsDotted_keyContext()
}

type Dotted_keyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDotted_keyContext() *Dotted_keyContext {
	var p = new(Dotted_keyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TomlParserRULE_dotted_key
	return p
}

func InitEmptyDotted_keyContext(p *Dotted_keyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TomlParserRULE_dotted_key
}

func (*Dotted_keyContext) IsDotted_keyContext() {}

func NewDotted_keyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Dotted_keyContext {
	var p = new(Dotted_keyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TomlParserRULE_dotted_key

	return p
}

func (s *Dotted_keyContext) GetParser() antlr.Parser { return s.parser }

func (s *Dotted_keyContext) AllSimple_key() []ISimple_keyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISimple_keyContext); ok {
			len++
		}
	}

	tst := make([]ISimple_keyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISimple_keyContext); ok {
			tst[i] = t.(ISimple_keyContext)
			i++
		}
	}

	return tst
}

func (s *Dotted_keyContext) Simple_key(i int) ISimple_keyContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimple_keyContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimple_keyContext)
}

func (s *Dotted_keyContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(TomlParserDOT)
}

func (s *Dotted_keyContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(TomlParserDOT, i)
}

func (s *Dotted_keyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Dotted_keyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Dotted_keyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlParserListener); ok {
		listenerT.EnterDotted_key(s)
	}
}

func (s *Dotted_keyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlParserListener); ok {
		listenerT.ExitDotted_key(s)
	}
}

func (p *TomlParser) Dotted_key() (localctx IDotted_keyContext) {
	localctx = NewDotted_keyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, TomlParserRULE_dotted_key)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(88)
		p.Simple_key()
	}
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == TomlParserDOT {
		{
			p.SetState(89)
			p.Match(TomlParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(90)
			p.Simple_key()
		}

		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	String_() IStringContext
	Integer() IIntegerContext
	Floating_point() IFloating_pointContext
	Bool_() IBool_Context
	Date_time() IDate_timeContext
	Array_() IArray_Context
	Inline_table() IInline_tableContext

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TomlParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TomlParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TomlParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) String_() IStringContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringContext)
}

func (s *ValueContext) Integer() IIntegerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntegerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *ValueContext) Floating_point() IFloating_pointContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFloating_pointContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFloating_pointContext)
}

func (s *ValueContext) Bool_() IBool_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBool_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBool_Context)
}

func (s *ValueContext) Date_time() IDate_timeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDate_timeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDate_timeContext)
}

func (s *ValueContext) Array_() IArray_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArray_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArray_Context)
}

func (s *ValueContext) Inline_table() IInline_tableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInline_tableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInline_tableContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlParserListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlParserListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *TomlParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, TomlParserRULE_value)
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TomlParserBASIC_STRING, TomlParserLITERAL_STRING, TomlParserML_BASIC_STRING, TomlParserML_LITERAL_STRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(95)
			p.String_()
		}

	case TomlParserDEC_INT, TomlParserHEX_INT, TomlParserOCT_INT, TomlParserBIN_INT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(96)
			p.Integer()
		}

	case TomlParserFLOAT, TomlParserINF, TomlParserNAN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(97)
			p.Floating_point()
		}

	case TomlParserBOOLEAN:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(98)
			p.Bool_()
		}

	case TomlParserOFFSET_DATE_TIME, TomlParserLOCAL_DATE_TIME, TomlParserLOCAL_DATE, TomlParserLOCAL_TIME:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(99)
			p.Date_time()
		}

	case TomlParserL_BRACKET:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(100)
			p.Array_()
		}

	case TomlParserL_BRACE:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(101)
			p.Inline_table()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStringContext is an interface to support dynamic dispatch.
type IStringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BASIC_STRING() antlr.TerminalNode
	ML_BASIC_STRING() antlr.TerminalNode
	LITERAL_STRING() antlr.TerminalNode
	ML_LITERAL_STRING() antlr.TerminalNode

	// IsStringContext differentiates from other interfaces.
	IsStringContext()
}

type StringContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringContext() *StringContext {
	var p = new(StringContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TomlParserRULE_string
	return p
}

func InitEmptyStringContext(p *StringContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TomlParserRULE_string
}

func (*StringContext) IsStringContext() {}

func NewStringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringContext {
	var p = new(StringContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TomlParserRULE_string

	return p
}

func (s *StringContext) GetParser() antlr.Parser { return s.parser }

func (s *StringContext) BASIC_STRING() antlr.TerminalNode {
	return s.GetToken(TomlParserBASIC_STRING, 0)
}

func (s *StringContext) ML_BASIC_STRING() antlr.TerminalNode {
	return s.GetToken(TomlParserML_BASIC_STRING, 0)
}

func (s *StringContext) LITERAL_STRING() antlr.TerminalNode {
	return s.GetToken(TomlParserLITERAL_STRING, 0)
}

func (s *StringContext) ML_LITERAL_STRING() antlr.TerminalNode {
	return s.GetToken(TomlParserML_LITERAL_STRING, 0)
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlParserListener); ok {
		listenerT.EnterString(s)
	}
}

func (s *StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlParserListener); ok {
		listenerT.ExitString(s)
	}
}

func (p *TomlParser) String_() (localctx IStringContext) {
	localctx = NewStringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, TomlParserRULE_string)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(104)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&399360) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIntegerContext is an interface to support dynamic dispatch.
type IIntegerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DEC_INT() antlr.TerminalNode
	HEX_INT() antlr.TerminalNode
	OCT_INT() antlr.TerminalNode
	BIN_INT() antlr.TerminalNode

	// IsIntegerContext differentiates from other interfaces.
	IsIntegerContext()
}

type IntegerContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerContext() *IntegerContext {
	var p = new(IntegerContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TomlParserRULE_integer
	return p
}

func InitEmptyIntegerContext(p *IntegerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TomlParserRULE_integer
}

func (*IntegerContext) IsIntegerContext() {}

func NewIntegerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerContext {
	var p = new(IntegerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TomlParserRULE_integer

	return p
}

func (s *IntegerContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerContext) DEC_INT() antlr.TerminalNode {
	return s.GetToken(TomlParserDEC_INT, 0)
}

func (s *IntegerContext) HEX_INT() antlr.TerminalNode {
	return s.GetToken(TomlParserHEX_INT, 0)
}

func (s *IntegerContext) OCT_INT() antlr.TerminalNode {
	return s.GetToken(TomlParserOCT_INT, 0)
}

func (s *IntegerContext) BIN_INT() antlr.TerminalNode {
	return s.GetToken(TomlParserBIN_INT, 0)
}

func (s *IntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlParserListener); ok {
		listenerT.EnterInteger(s)
	}
}

func (s *IntegerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlParserListener); ok {
		listenerT.ExitInteger(s)
	}
}

func (p *TomlParser) Integer() (localctx IIntegerContext) {
	localctx = NewIntegerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, TomlParserRULE_integer)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(106)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&62914560) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFloating_pointContext is an interface to support dynamic dispatch.
type IFloating_pointContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FLOAT() antlr.TerminalNode
	INF() antlr.TerminalNode
	NAN() antlr.TerminalNode

	// IsFloating_pointContext differentiates from other interfaces.
	IsFloating_pointContext()
}

type Floating_pointContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFloating_pointContext() *Floating_pointContext {
	var p = new(Floating_pointContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TomlParserRULE_floating_point
	return p
}

func InitEmptyFloating_pointContext(p *Floating_pointContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TomlParserRULE_floating_point
}

func (*Floating_pointContext) IsFloating_pointContext() {}

func NewFloating_pointContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Floating_pointContext {
	var p = new(Floating_pointContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TomlParserRULE_floating_point

	return p
}

func (s *Floating_pointContext) GetParser() antlr.Parser { return s.parser }

func (s *Floating_pointContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(TomlParserFLOAT, 0)
}

func (s *Floating_pointContext) INF() antlr.TerminalNode {
	return s.GetToken(TomlParserINF, 0)
}

func (s *Floating_pointContext) NAN() antlr.TerminalNode {
	return s.GetToken(TomlParserNAN, 0)
}

func (s *Floating_pointContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Floating_pointContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Floating_pointContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlParserListener); ok {
		listenerT.EnterFloating_point(s)
	}
}

func (s *Floating_pointContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlParserListener); ok {
		listenerT.ExitFloating_point(s)
	}
}

func (p *TomlParser) Floating_point() (localctx IFloating_pointContext) {
	localctx = NewFloating_pointContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, TomlParserRULE_floating_point)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(108)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3670016) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBool_Context is an interface to support dynamic dispatch.
type IBool_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BOOLEAN() antlr.TerminalNode

	// IsBool_Context differentiates from other interfaces.
	IsBool_Context()
}

type Bool_Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBool_Context() *Bool_Context {
	var p = new(Bool_Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TomlParserRULE_bool_
	return p
}

func InitEmptyBool_Context(p *Bool_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TomlParserRULE_bool_
}

func (*Bool_Context) IsBool_Context() {}

func NewBool_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Bool_Context {
	var p = new(Bool_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TomlParserRULE_bool_

	return p
}

func (s *Bool_Context) GetParser() antlr.Parser { return s.parser }

func (s *Bool_Context) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(TomlParserBOOLEAN, 0)
}

func (s *Bool_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Bool_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Bool_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlParserListener); ok {
		listenerT.EnterBool_(s)
	}
}

func (s *Bool_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlParserListener); ok {
		listenerT.ExitBool_(s)
	}
}

func (p *TomlParser) Bool_() (localctx IBool_Context) {
	localctx = NewBool_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, TomlParserRULE_bool_)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(110)
		p.Match(TomlParserBOOLEAN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDate_timeContext is an interface to support dynamic dispatch.
type IDate_timeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OFFSET_DATE_TIME() antlr.TerminalNode
	LOCAL_DATE_TIME() antlr.TerminalNode
	LOCAL_DATE() antlr.TerminalNode
	LOCAL_TIME() antlr.TerminalNode

	// IsDate_timeContext differentiates from other interfaces.
	IsDate_timeContext()
}

type Date_timeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDate_timeContext() *Date_timeContext {
	var p = new(Date_timeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TomlParserRULE_date_time
	return p
}

func InitEmptyDate_timeContext(p *Date_timeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TomlParserRULE_date_time
}

func (*Date_timeContext) IsDate_timeContext() {}

func NewDate_timeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Date_timeContext {
	var p = new(Date_timeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TomlParserRULE_date_time

	return p
}

func (s *Date_timeContext) GetParser() antlr.Parser { return s.parser }

func (s *Date_timeContext) OFFSET_DATE_TIME() antlr.TerminalNode {
	return s.GetToken(TomlParserOFFSET_DATE_TIME, 0)
}

func (s *Date_timeContext) LOCAL_DATE_TIME() antlr.TerminalNode {
	return s.GetToken(TomlParserLOCAL_DATE_TIME, 0)
}

func (s *Date_timeContext) LOCAL_DATE() antlr.TerminalNode {
	return s.GetToken(TomlParserLOCAL_DATE, 0)
}

func (s *Date_timeContext) LOCAL_TIME() antlr.TerminalNode {
	return s.GetToken(TomlParserLOCAL_TIME, 0)
}

func (s *Date_timeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Date_timeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Date_timeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlParserListener); ok {
		listenerT.EnterDate_time(s)
	}
}

func (s *Date_timeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlParserListener); ok {
		listenerT.ExitDate_time(s)
	}
}

func (p *TomlParser) Date_time() (localctx IDate_timeContext) {
	localctx = NewDate_timeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, TomlParserRULE_date_time)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(112)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1006632960) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArray_Context is an interface to support dynamic dispatch.
type IArray_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	L_BRACKET() antlr.TerminalNode
	Comment_or_nl() IComment_or_nlContext
	R_BRACKET() antlr.TerminalNode
	Array_values() IArray_valuesContext

	// IsArray_Context differentiates from other interfaces.
	IsArray_Context()
}

type Array_Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArray_Context() *Array_Context {
	var p = new(Array_Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TomlParserRULE_array_
	return p
}

func InitEmptyArray_Context(p *Array_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TomlParserRULE_array_
}

func (*Array_Context) IsArray_Context() {}

func NewArray_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Array_Context {
	var p = new(Array_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TomlParserRULE_array_

	return p
}

func (s *Array_Context) GetParser() antlr.Parser { return s.parser }

func (s *Array_Context) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(TomlParserL_BRACKET, 0)
}

func (s *Array_Context) Comment_or_nl() IComment_or_nlContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComment_or_nlContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComment_or_nlContext)
}

func (s *Array_Context) R_BRACKET() antlr.TerminalNode {
	return s.GetToken(TomlParserR_BRACKET, 0)
}

func (s *Array_Context) Array_values() IArray_valuesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArray_valuesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArray_valuesContext)
}

func (s *Array_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Array_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Array_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlParserListener); ok {
		listenerT.EnterArray_(s)
	}
}

func (s *Array_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlParserListener); ok {
		listenerT.ExitArray_(s)
	}
}

func (p *TomlParser) Array_() (localctx IArray_Context) {
	localctx = NewArray_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, TomlParserRULE_array_)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(114)
		p.Match(TomlParserL_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(116)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(115)
			p.Array_values()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(118)
		p.Comment_or_nl()
	}
	{
		p.SetState(119)
		p.Match(TomlParserR_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArray_valuesContext is an interface to support dynamic dispatch.
type IArray_valuesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllComment_or_nl() []IComment_or_nlContext
	Comment_or_nl(i int) IComment_or_nlContext
	Value() IValueContext
	Nl_or_comment() INl_or_commentContext
	COMMA() antlr.TerminalNode
	Array_values() IArray_valuesContext

	// IsArray_valuesContext differentiates from other interfaces.
	IsArray_valuesContext()
}

type Array_valuesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArray_valuesContext() *Array_valuesContext {
	var p = new(Array_valuesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TomlParserRULE_array_values
	return p
}

func InitEmptyArray_valuesContext(p *Array_valuesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TomlParserRULE_array_values
}

func (*Array_valuesContext) IsArray_valuesContext() {}

func NewArray_valuesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Array_valuesContext {
	var p = new(Array_valuesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TomlParserRULE_array_values

	return p
}

func (s *Array_valuesContext) GetParser() antlr.Parser { return s.parser }

func (s *Array_valuesContext) AllComment_or_nl() []IComment_or_nlContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IComment_or_nlContext); ok {
			len++
		}
	}

	tst := make([]IComment_or_nlContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IComment_or_nlContext); ok {
			tst[i] = t.(IComment_or_nlContext)
			i++
		}
	}

	return tst
}

func (s *Array_valuesContext) Comment_or_nl(i int) IComment_or_nlContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComment_or_nlContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComment_or_nlContext)
}

func (s *Array_valuesContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *Array_valuesContext) Nl_or_comment() INl_or_commentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INl_or_commentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INl_or_commentContext)
}

func (s *Array_valuesContext) COMMA() antlr.TerminalNode {
	return s.GetToken(TomlParserCOMMA, 0)
}

func (s *Array_valuesContext) Array_values() IArray_valuesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArray_valuesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArray_valuesContext)
}

func (s *Array_valuesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Array_valuesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Array_valuesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlParserListener); ok {
		listenerT.EnterArray_values(s)
	}
}

func (s *Array_valuesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlParserListener); ok {
		listenerT.ExitArray_values(s)
	}
}

func (p *TomlParser) Array_values() (localctx IArray_valuesContext) {
	localctx = NewArray_valuesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, TomlParserRULE_array_values)
	var _la int

	p.SetState(134)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(121)
			p.Comment_or_nl()
		}
		{
			p.SetState(122)
			p.Value()
		}
		{
			p.SetState(123)
			p.Nl_or_comment()
		}
		{
			p.SetState(124)
			p.Match(TomlParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(125)
			p.Array_values()
		}
		{
			p.SetState(126)
			p.Comment_or_nl()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(128)
			p.Comment_or_nl()
		}
		{
			p.SetState(129)
			p.Value()
		}
		{
			p.SetState(130)
			p.Nl_or_comment()
		}
		p.SetState(132)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == TomlParserCOMMA {
			{
				p.SetState(131)
				p.Match(TomlParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IComment_or_nlContext is an interface to support dynamic dispatch.
type IComment_or_nlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode
	AllCOMMENT() []antlr.TerminalNode
	COMMENT(i int) antlr.TerminalNode

	// IsComment_or_nlContext differentiates from other interfaces.
	IsComment_or_nlContext()
}

type Comment_or_nlContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComment_or_nlContext() *Comment_or_nlContext {
	var p = new(Comment_or_nlContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TomlParserRULE_comment_or_nl
	return p
}

func InitEmptyComment_or_nlContext(p *Comment_or_nlContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TomlParserRULE_comment_or_nl
}

func (*Comment_or_nlContext) IsComment_or_nlContext() {}

func NewComment_or_nlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Comment_or_nlContext {
	var p = new(Comment_or_nlContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TomlParserRULE_comment_or_nl

	return p
}

func (s *Comment_or_nlContext) GetParser() antlr.Parser { return s.parser }

func (s *Comment_or_nlContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(TomlParserNL)
}

func (s *Comment_or_nlContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(TomlParserNL, i)
}

func (s *Comment_or_nlContext) AllCOMMENT() []antlr.TerminalNode {
	return s.GetTokens(TomlParserCOMMENT)
}

func (s *Comment_or_nlContext) COMMENT(i int) antlr.TerminalNode {
	return s.GetToken(TomlParserCOMMENT, i)
}

func (s *Comment_or_nlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Comment_or_nlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Comment_or_nlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlParserListener); ok {
		listenerT.EnterComment_or_nl(s)
	}
}

func (s *Comment_or_nlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlParserListener); ok {
		listenerT.ExitComment_or_nl(s)
	}
}

func (p *TomlParser) Comment_or_nl() (localctx IComment_or_nlContext) {
	localctx = NewComment_or_nlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, TomlParserRULE_comment_or_nl)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(137)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == TomlParserCOMMENT {
				{
					p.SetState(136)
					p.Match(TomlParserCOMMENT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(139)
				p.Match(TomlParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(144)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INl_or_commentContext is an interface to support dynamic dispatch.
type INl_or_commentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode
	AllCOMMENT() []antlr.TerminalNode
	COMMENT(i int) antlr.TerminalNode

	// IsNl_or_commentContext differentiates from other interfaces.
	IsNl_or_commentContext()
}

type Nl_or_commentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNl_or_commentContext() *Nl_or_commentContext {
	var p = new(Nl_or_commentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TomlParserRULE_nl_or_comment
	return p
}

func InitEmptyNl_or_commentContext(p *Nl_or_commentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TomlParserRULE_nl_or_comment
}

func (*Nl_or_commentContext) IsNl_or_commentContext() {}

func NewNl_or_commentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Nl_or_commentContext {
	var p = new(Nl_or_commentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TomlParserRULE_nl_or_comment

	return p
}

func (s *Nl_or_commentContext) GetParser() antlr.Parser { return s.parser }

func (s *Nl_or_commentContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(TomlParserNL)
}

func (s *Nl_or_commentContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(TomlParserNL, i)
}

func (s *Nl_or_commentContext) AllCOMMENT() []antlr.TerminalNode {
	return s.GetTokens(TomlParserCOMMENT)
}

func (s *Nl_or_commentContext) COMMENT(i int) antlr.TerminalNode {
	return s.GetToken(TomlParserCOMMENT, i)
}

func (s *Nl_or_commentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Nl_or_commentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Nl_or_commentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlParserListener); ok {
		listenerT.EnterNl_or_comment(s)
	}
}

func (s *Nl_or_commentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlParserListener); ok {
		listenerT.ExitNl_or_comment(s)
	}
}

func (p *TomlParser) Nl_or_comment() (localctx INl_or_commentContext) {
	localctx = NewNl_or_commentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, TomlParserRULE_nl_or_comment)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(145)
				p.Match(TomlParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(147)
			p.GetErrorHandler().Sync(p)

			if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) == 1 {
				{
					p.SetState(146)
					p.Match(TomlParserCOMMENT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			} else if p.HasError() { // JIM
				goto errorExit
			}

		}
		p.SetState(153)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITableContext is an interface to support dynamic dispatch.
type ITableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Standard_table() IStandard_tableContext
	Array_table() IArray_tableContext

	// IsTableContext differentiates from other interfaces.
	IsTableContext()
}

type TableContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableContext() *TableContext {
	var p = new(TableContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TomlParserRULE_table
	return p
}

func InitEmptyTableContext(p *TableContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TomlParserRULE_table
}

func (*TableContext) IsTableContext() {}

func NewTableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableContext {
	var p = new(TableContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TomlParserRULE_table

	return p
}

func (s *TableContext) GetParser() antlr.Parser { return s.parser }

func (s *TableContext) Standard_table() IStandard_tableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStandard_tableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStandard_tableContext)
}

func (s *TableContext) Array_table() IArray_tableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArray_tableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArray_tableContext)
}

func (s *TableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlParserListener); ok {
		listenerT.EnterTable(s)
	}
}

func (s *TableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlParserListener); ok {
		listenerT.ExitTable(s)
	}
}

func (p *TomlParser) Table() (localctx ITableContext) {
	localctx = NewTableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, TomlParserRULE_table)
	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TomlParserL_BRACKET:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(154)
			p.Standard_table()
		}

	case TomlParserDOUBLE_L_BRACKET:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(155)
			p.Array_table()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStandard_tableContext is an interface to support dynamic dispatch.
type IStandard_tableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	L_BRACKET() antlr.TerminalNode
	Key() IKeyContext
	R_BRACKET() antlr.TerminalNode

	// IsStandard_tableContext differentiates from other interfaces.
	IsStandard_tableContext()
}

type Standard_tableContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStandard_tableContext() *Standard_tableContext {
	var p = new(Standard_tableContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TomlParserRULE_standard_table
	return p
}

func InitEmptyStandard_tableContext(p *Standard_tableContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TomlParserRULE_standard_table
}

func (*Standard_tableContext) IsStandard_tableContext() {}

func NewStandard_tableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Standard_tableContext {
	var p = new(Standard_tableContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TomlParserRULE_standard_table

	return p
}

func (s *Standard_tableContext) GetParser() antlr.Parser { return s.parser }

func (s *Standard_tableContext) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(TomlParserL_BRACKET, 0)
}

func (s *Standard_tableContext) Key() IKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeyContext)
}

func (s *Standard_tableContext) R_BRACKET() antlr.TerminalNode {
	return s.GetToken(TomlParserR_BRACKET, 0)
}

func (s *Standard_tableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Standard_tableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Standard_tableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlParserListener); ok {
		listenerT.EnterStandard_table(s)
	}
}

func (s *Standard_tableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlParserListener); ok {
		listenerT.ExitStandard_table(s)
	}
}

func (p *TomlParser) Standard_table() (localctx IStandard_tableContext) {
	localctx = NewStandard_tableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, TomlParserRULE_standard_table)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(158)
		p.Match(TomlParserL_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(159)
		p.Key()
	}
	{
		p.SetState(160)
		p.Match(TomlParserR_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInline_tableContext is an interface to support dynamic dispatch.
type IInline_tableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	L_BRACE() antlr.TerminalNode
	Inline_table_keyvals() IInline_table_keyvalsContext
	R_BRACE() antlr.TerminalNode

	// IsInline_tableContext differentiates from other interfaces.
	IsInline_tableContext()
}

type Inline_tableContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInline_tableContext() *Inline_tableContext {
	var p = new(Inline_tableContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TomlParserRULE_inline_table
	return p
}

func InitEmptyInline_tableContext(p *Inline_tableContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TomlParserRULE_inline_table
}

func (*Inline_tableContext) IsInline_tableContext() {}

func NewInline_tableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inline_tableContext {
	var p = new(Inline_tableContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TomlParserRULE_inline_table

	return p
}

func (s *Inline_tableContext) GetParser() antlr.Parser { return s.parser }

func (s *Inline_tableContext) L_BRACE() antlr.TerminalNode {
	return s.GetToken(TomlParserL_BRACE, 0)
}

func (s *Inline_tableContext) Inline_table_keyvals() IInline_table_keyvalsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInline_table_keyvalsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInline_table_keyvalsContext)
}

func (s *Inline_tableContext) R_BRACE() antlr.TerminalNode {
	return s.GetToken(TomlParserR_BRACE, 0)
}

func (s *Inline_tableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inline_tableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Inline_tableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlParserListener); ok {
		listenerT.EnterInline_table(s)
	}
}

func (s *Inline_tableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlParserListener); ok {
		listenerT.ExitInline_table(s)
	}
}

func (p *TomlParser) Inline_table() (localctx IInline_tableContext) {
	localctx = NewInline_tableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, TomlParserRULE_inline_table)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(162)
		p.Match(TomlParserL_BRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(163)
		p.Inline_table_keyvals()
	}
	{
		p.SetState(164)
		p.Match(TomlParserR_BRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInline_table_keyvalsContext is an interface to support dynamic dispatch.
type IInline_table_keyvalsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Inline_table_keyvals_non_empty() IInline_table_keyvals_non_emptyContext

	// IsInline_table_keyvalsContext differentiates from other interfaces.
	IsInline_table_keyvalsContext()
}

type Inline_table_keyvalsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInline_table_keyvalsContext() *Inline_table_keyvalsContext {
	var p = new(Inline_table_keyvalsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TomlParserRULE_inline_table_keyvals
	return p
}

func InitEmptyInline_table_keyvalsContext(p *Inline_table_keyvalsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TomlParserRULE_inline_table_keyvals
}

func (*Inline_table_keyvalsContext) IsInline_table_keyvalsContext() {}

func NewInline_table_keyvalsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inline_table_keyvalsContext {
	var p = new(Inline_table_keyvalsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TomlParserRULE_inline_table_keyvals

	return p
}

func (s *Inline_table_keyvalsContext) GetParser() antlr.Parser { return s.parser }

func (s *Inline_table_keyvalsContext) Inline_table_keyvals_non_empty() IInline_table_keyvals_non_emptyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInline_table_keyvals_non_emptyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInline_table_keyvals_non_emptyContext)
}

func (s *Inline_table_keyvalsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inline_table_keyvalsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Inline_table_keyvalsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlParserListener); ok {
		listenerT.EnterInline_table_keyvals(s)
	}
}

func (s *Inline_table_keyvalsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlParserListener); ok {
		listenerT.ExitInline_table_keyvals(s)
	}
}

func (p *TomlParser) Inline_table_keyvals() (localctx IInline_table_keyvalsContext) {
	localctx = NewInline_table_keyvalsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, TomlParserRULE_inline_table_keyvals)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&14336) != 0 {
		{
			p.SetState(166)
			p.Inline_table_keyvals_non_empty()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInline_table_keyvals_non_emptyContext is an interface to support dynamic dispatch.
type IInline_table_keyvals_non_emptyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Key() IKeyContext
	EQUALS() antlr.TerminalNode
	Value() IValueContext
	COMMA() antlr.TerminalNode
	Inline_table_keyvals_non_empty() IInline_table_keyvals_non_emptyContext

	// IsInline_table_keyvals_non_emptyContext differentiates from other interfaces.
	IsInline_table_keyvals_non_emptyContext()
}

type Inline_table_keyvals_non_emptyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInline_table_keyvals_non_emptyContext() *Inline_table_keyvals_non_emptyContext {
	var p = new(Inline_table_keyvals_non_emptyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TomlParserRULE_inline_table_keyvals_non_empty
	return p
}

func InitEmptyInline_table_keyvals_non_emptyContext(p *Inline_table_keyvals_non_emptyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TomlParserRULE_inline_table_keyvals_non_empty
}

func (*Inline_table_keyvals_non_emptyContext) IsInline_table_keyvals_non_emptyContext() {}

func NewInline_table_keyvals_non_emptyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inline_table_keyvals_non_emptyContext {
	var p = new(Inline_table_keyvals_non_emptyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TomlParserRULE_inline_table_keyvals_non_empty

	return p
}

func (s *Inline_table_keyvals_non_emptyContext) GetParser() antlr.Parser { return s.parser }

func (s *Inline_table_keyvals_non_emptyContext) Key() IKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeyContext)
}

func (s *Inline_table_keyvals_non_emptyContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(TomlParserEQUALS, 0)
}

func (s *Inline_table_keyvals_non_emptyContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *Inline_table_keyvals_non_emptyContext) COMMA() antlr.TerminalNode {
	return s.GetToken(TomlParserCOMMA, 0)
}

func (s *Inline_table_keyvals_non_emptyContext) Inline_table_keyvals_non_empty() IInline_table_keyvals_non_emptyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInline_table_keyvals_non_emptyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInline_table_keyvals_non_emptyContext)
}

func (s *Inline_table_keyvals_non_emptyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inline_table_keyvals_non_emptyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Inline_table_keyvals_non_emptyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlParserListener); ok {
		listenerT.EnterInline_table_keyvals_non_empty(s)
	}
}

func (s *Inline_table_keyvals_non_emptyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlParserListener); ok {
		listenerT.ExitInline_table_keyvals_non_empty(s)
	}
}

func (p *TomlParser) Inline_table_keyvals_non_empty() (localctx IInline_table_keyvals_non_emptyContext) {
	localctx = NewInline_table_keyvals_non_emptyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, TomlParserRULE_inline_table_keyvals_non_empty)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(169)
		p.Key()
	}
	{
		p.SetState(170)
		p.Match(TomlParserEQUALS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(171)
		p.Value()
	}
	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TomlParserCOMMA {
		{
			p.SetState(172)
			p.Match(TomlParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(173)
			p.Inline_table_keyvals_non_empty()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArray_tableContext is an interface to support dynamic dispatch.
type IArray_tableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DOUBLE_L_BRACKET() antlr.TerminalNode
	Key() IKeyContext
	DOUBLE_R_BRACKET() antlr.TerminalNode

	// IsArray_tableContext differentiates from other interfaces.
	IsArray_tableContext()
}

type Array_tableContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArray_tableContext() *Array_tableContext {
	var p = new(Array_tableContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TomlParserRULE_array_table
	return p
}

func InitEmptyArray_tableContext(p *Array_tableContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TomlParserRULE_array_table
}

func (*Array_tableContext) IsArray_tableContext() {}

func NewArray_tableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Array_tableContext {
	var p = new(Array_tableContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TomlParserRULE_array_table

	return p
}

func (s *Array_tableContext) GetParser() antlr.Parser { return s.parser }

func (s *Array_tableContext) DOUBLE_L_BRACKET() antlr.TerminalNode {
	return s.GetToken(TomlParserDOUBLE_L_BRACKET, 0)
}

func (s *Array_tableContext) Key() IKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeyContext)
}

func (s *Array_tableContext) DOUBLE_R_BRACKET() antlr.TerminalNode {
	return s.GetToken(TomlParserDOUBLE_R_BRACKET, 0)
}

func (s *Array_tableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Array_tableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Array_tableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlParserListener); ok {
		listenerT.EnterArray_table(s)
	}
}

func (s *Array_tableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TomlParserListener); ok {
		listenerT.ExitArray_table(s)
	}
}

func (p *TomlParser) Array_table() (localctx IArray_tableContext) {
	localctx = NewArray_tableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, TomlParserRULE_array_table)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(176)
		p.Match(TomlParserDOUBLE_L_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(177)
		p.Key()
	}
	{
		p.SetState(178)
		p.Match(TomlParserDOUBLE_R_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
