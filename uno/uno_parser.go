// Code generated from uno.g4 by ANTLR 4.13.1. DO NOT EDIT.

package uno // uno
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

type unoParser struct {
	*antlr.BaseParser
}

var UnoParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func unoParserInit() {
	staticData := &UnoParserStaticData
	staticData.LiteralNames = []string{
		"", "'('", "')'", "','", "'.'", "'['", "']'", "'+'", "'-'", "'*'", "'/'",
		"'%'", "", "", "", "", "", "", "", "", "", "", "", "'='", "'=='", "'<>'",
		"'!='", "'>'", "'>='", "'<'", "'<='",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "T_ADD", "T_SUB", "T_MUL", "T_DIV", "T_MOD",
		"T_INT", "T_FLOAT", "T_STRING", "T_ON", "T_AND", "T_OR", "T_NOT", "T_IN",
		"T_TRUE", "T_FALSE", "T_COMPARE", "T_EQUAL", "T_EQUAL2", "T_NOTEQUAL",
		"T_NOTEQUAL2", "T_GREATER", "T_GREATEREQUAL", "T_LESS", "T_LESSEQUAL",
		"IDENTIFIER", "INTEGER_LIST", "INTEGER", "DECIMAL_LIST", "DECIMAL",
		"STRING_LIST", "STRING", "WS",
	}
	staticData.RuleNames = []string{
		"start", "boolean_expression", "arithmetic_expression", "type_marker",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 38, 102, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1, 0, 1,
		0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3,
		1, 34, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 42, 8, 1, 10, 1,
		12, 1, 45, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		5, 2, 56, 8, 2, 10, 2, 12, 2, 59, 9, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 76, 8, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 5, 2, 93, 8, 2, 10, 2, 12, 2, 96, 9, 2, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 0, 2, 2, 4, 4, 0, 2, 4, 6, 0, 2, 3, 0, 32, 32, 34, 34, 36,
		36, 1, 0, 12, 14, 118, 0, 8, 1, 0, 0, 0, 2, 33, 1, 0, 0, 0, 4, 75, 1, 0,
		0, 0, 6, 97, 1, 0, 0, 0, 8, 9, 3, 2, 1, 0, 9, 10, 5, 0, 0, 1, 10, 1, 1,
		0, 0, 0, 11, 12, 6, 1, -1, 0, 12, 13, 3, 4, 2, 0, 13, 14, 5, 22, 0, 0,
		14, 15, 3, 4, 2, 0, 15, 34, 1, 0, 0, 0, 16, 17, 5, 18, 0, 0, 17, 34, 3,
		2, 1, 6, 18, 19, 3, 4, 2, 0, 19, 20, 5, 19, 0, 0, 20, 21, 7, 0, 0, 0, 21,
		34, 1, 0, 0, 0, 22, 23, 3, 4, 2, 0, 23, 24, 5, 18, 0, 0, 24, 25, 5, 19,
		0, 0, 25, 26, 7, 0, 0, 0, 26, 34, 1, 0, 0, 0, 27, 28, 5, 1, 0, 0, 28, 29,
		3, 2, 1, 0, 29, 30, 5, 2, 0, 0, 30, 34, 1, 0, 0, 0, 31, 34, 5, 20, 0, 0,
		32, 34, 5, 21, 0, 0, 33, 11, 1, 0, 0, 0, 33, 16, 1, 0, 0, 0, 33, 18, 1,
		0, 0, 0, 33, 22, 1, 0, 0, 0, 33, 27, 1, 0, 0, 0, 33, 31, 1, 0, 0, 0, 33,
		32, 1, 0, 0, 0, 34, 43, 1, 0, 0, 0, 35, 36, 10, 9, 0, 0, 36, 37, 5, 16,
		0, 0, 37, 42, 3, 2, 1, 10, 38, 39, 10, 8, 0, 0, 39, 40, 5, 17, 0, 0, 40,
		42, 3, 2, 1, 9, 41, 35, 1, 0, 0, 0, 41, 38, 1, 0, 0, 0, 42, 45, 1, 0, 0,
		0, 43, 41, 1, 0, 0, 0, 43, 44, 1, 0, 0, 0, 44, 3, 1, 0, 0, 0, 45, 43, 1,
		0, 0, 0, 46, 47, 6, 2, -1, 0, 47, 48, 5, 31, 0, 0, 48, 49, 5, 1, 0, 0,
		49, 76, 5, 2, 0, 0, 50, 51, 5, 31, 0, 0, 51, 52, 5, 1, 0, 0, 52, 57, 3,
		4, 2, 0, 53, 54, 5, 3, 0, 0, 54, 56, 3, 4, 2, 0, 55, 53, 1, 0, 0, 0, 56,
		59, 1, 0, 0, 0, 57, 55, 1, 0, 0, 0, 57, 58, 1, 0, 0, 0, 58, 60, 1, 0, 0,
		0, 59, 57, 1, 0, 0, 0, 60, 61, 5, 2, 0, 0, 61, 76, 1, 0, 0, 0, 62, 63,
		5, 31, 0, 0, 63, 76, 3, 6, 3, 0, 64, 65, 5, 31, 0, 0, 65, 66, 5, 4, 0,
		0, 66, 67, 5, 31, 0, 0, 67, 76, 3, 6, 3, 0, 68, 76, 5, 37, 0, 0, 69, 76,
		5, 33, 0, 0, 70, 76, 5, 35, 0, 0, 71, 72, 5, 1, 0, 0, 72, 73, 3, 4, 2,
		0, 73, 74, 5, 2, 0, 0, 74, 76, 1, 0, 0, 0, 75, 46, 1, 0, 0, 0, 75, 50,
		1, 0, 0, 0, 75, 62, 1, 0, 0, 0, 75, 64, 1, 0, 0, 0, 75, 68, 1, 0, 0, 0,
		75, 69, 1, 0, 0, 0, 75, 70, 1, 0, 0, 0, 75, 71, 1, 0, 0, 0, 76, 94, 1,
		0, 0, 0, 77, 78, 10, 13, 0, 0, 78, 79, 5, 11, 0, 0, 79, 93, 3, 4, 2, 14,
		80, 81, 10, 12, 0, 0, 81, 82, 5, 9, 0, 0, 82, 93, 3, 4, 2, 13, 83, 84,
		10, 11, 0, 0, 84, 85, 5, 10, 0, 0, 85, 93, 3, 4, 2, 12, 86, 87, 10, 10,
		0, 0, 87, 88, 5, 7, 0, 0, 88, 93, 3, 4, 2, 11, 89, 90, 10, 9, 0, 0, 90,
		91, 5, 8, 0, 0, 91, 93, 3, 4, 2, 10, 92, 77, 1, 0, 0, 0, 92, 80, 1, 0,
		0, 0, 92, 83, 1, 0, 0, 0, 92, 86, 1, 0, 0, 0, 92, 89, 1, 0, 0, 0, 93, 96,
		1, 0, 0, 0, 94, 92, 1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95, 5, 1, 0, 0, 0,
		96, 94, 1, 0, 0, 0, 97, 98, 5, 5, 0, 0, 98, 99, 7, 1, 0, 0, 99, 100, 5,
		6, 0, 0, 100, 7, 1, 0, 0, 0, 7, 33, 41, 43, 57, 75, 92, 94,
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

// unoParserInit initializes any static state used to implement unoParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewunoParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func UnoParserInit() {
	staticData := &UnoParserStaticData
	staticData.once.Do(unoParserInit)
}

// NewunoParser produces a new parser instance for the optional input antlr.TokenStream.
func NewunoParser(input antlr.TokenStream) *unoParser {
	UnoParserInit()
	this := new(unoParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &UnoParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "uno.g4"

	return this
}

// unoParser tokens.
const (
	unoParserEOF            = antlr.TokenEOF
	unoParserT__0           = 1
	unoParserT__1           = 2
	unoParserT__2           = 3
	unoParserT__3           = 4
	unoParserT__4           = 5
	unoParserT__5           = 6
	unoParserT_ADD          = 7
	unoParserT_SUB          = 8
	unoParserT_MUL          = 9
	unoParserT_DIV          = 10
	unoParserT_MOD          = 11
	unoParserT_INT          = 12
	unoParserT_FLOAT        = 13
	unoParserT_STRING       = 14
	unoParserT_ON           = 15
	unoParserT_AND          = 16
	unoParserT_OR           = 17
	unoParserT_NOT          = 18
	unoParserT_IN           = 19
	unoParserT_TRUE         = 20
	unoParserT_FALSE        = 21
	unoParserT_COMPARE      = 22
	unoParserT_EQUAL        = 23
	unoParserT_EQUAL2       = 24
	unoParserT_NOTEQUAL     = 25
	unoParserT_NOTEQUAL2    = 26
	unoParserT_GREATER      = 27
	unoParserT_GREATEREQUAL = 28
	unoParserT_LESS         = 29
	unoParserT_LESSEQUAL    = 30
	unoParserIDENTIFIER     = 31
	unoParserINTEGER_LIST   = 32
	unoParserINTEGER        = 33
	unoParserDECIMAL_LIST   = 34
	unoParserDECIMAL        = 35
	unoParserSTRING_LIST    = 36
	unoParserSTRING         = 37
	unoParserWS             = 38
)

// unoParser rules.
const (
	unoParserRULE_start                 = 0
	unoParserRULE_boolean_expression    = 1
	unoParserRULE_arithmetic_expression = 2
	unoParserRULE_type_marker           = 3
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Boolean_expression() IBoolean_expressionContext
	EOF() antlr.TerminalNode

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = unoParserRULE_start
	return p
}

func InitEmptyStartContext(p *StartContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = unoParserRULE_start
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = unoParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Boolean_expression() IBoolean_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolean_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolean_expressionContext)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(unoParserEOF, 0)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(unoListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(unoListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *unoParser) Start_() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, unoParserRULE_start)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(8)
		p.boolean_expression(0)
	}
	{
		p.SetState(9)
		p.Match(unoParserEOF)
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

// IBoolean_expressionContext is an interface to support dynamic dispatch.
type IBoolean_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsBoolean_expressionContext differentiates from other interfaces.
	IsBoolean_expressionContext()
}

type Boolean_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolean_expressionContext() *Boolean_expressionContext {
	var p = new(Boolean_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = unoParserRULE_boolean_expression
	return p
}

func InitEmptyBoolean_expressionContext(p *Boolean_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = unoParserRULE_boolean_expression
}

func (*Boolean_expressionContext) IsBoolean_expressionContext() {}

func NewBoolean_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Boolean_expressionContext {
	var p = new(Boolean_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = unoParserRULE_boolean_expression

	return p
}

func (s *Boolean_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Boolean_expressionContext) CopyAll(ctx *Boolean_expressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Boolean_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Boolean_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type CmpBooleanExpressionContext struct {
	Boolean_expressionContext
}

func NewCmpBooleanExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CmpBooleanExpressionContext {
	var p = new(CmpBooleanExpressionContext)

	InitEmptyBoolean_expressionContext(&p.Boolean_expressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Boolean_expressionContext))

	return p
}

func (s *CmpBooleanExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CmpBooleanExpressionContext) AllArithmetic_expression() []IArithmetic_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArithmetic_expressionContext); ok {
			len++
		}
	}

	tst := make([]IArithmetic_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArithmetic_expressionContext); ok {
			tst[i] = t.(IArithmetic_expressionContext)
			i++
		}
	}

	return tst
}

func (s *CmpBooleanExpressionContext) Arithmetic_expression(i int) IArithmetic_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArithmetic_expressionContext); ok {
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

	return t.(IArithmetic_expressionContext)
}

func (s *CmpBooleanExpressionContext) T_COMPARE() antlr.TerminalNode {
	return s.GetToken(unoParserT_COMPARE, 0)
}

func (s *CmpBooleanExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(unoListener); ok {
		listenerT.EnterCmpBooleanExpression(s)
	}
}

func (s *CmpBooleanExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(unoListener); ok {
		listenerT.ExitCmpBooleanExpression(s)
	}
}

type NotBooleanExpressionContext struct {
	Boolean_expressionContext
}

func NewNotBooleanExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotBooleanExpressionContext {
	var p = new(NotBooleanExpressionContext)

	InitEmptyBoolean_expressionContext(&p.Boolean_expressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Boolean_expressionContext))

	return p
}

func (s *NotBooleanExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotBooleanExpressionContext) T_NOT() antlr.TerminalNode {
	return s.GetToken(unoParserT_NOT, 0)
}

func (s *NotBooleanExpressionContext) Boolean_expression() IBoolean_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolean_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolean_expressionContext)
}

func (s *NotBooleanExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(unoListener); ok {
		listenerT.EnterNotBooleanExpression(s)
	}
}

func (s *NotBooleanExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(unoListener); ok {
		listenerT.ExitNotBooleanExpression(s)
	}
}

type PlainBooleanExpressionContext struct {
	Boolean_expressionContext
}

func NewPlainBooleanExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PlainBooleanExpressionContext {
	var p = new(PlainBooleanExpressionContext)

	InitEmptyBoolean_expressionContext(&p.Boolean_expressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Boolean_expressionContext))

	return p
}

func (s *PlainBooleanExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PlainBooleanExpressionContext) Boolean_expression() IBoolean_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolean_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolean_expressionContext)
}

func (s *PlainBooleanExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(unoListener); ok {
		listenerT.EnterPlainBooleanExpression(s)
	}
}

func (s *PlainBooleanExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(unoListener); ok {
		listenerT.ExitPlainBooleanExpression(s)
	}
}

type OrBooleanExpressionContext struct {
	Boolean_expressionContext
}

func NewOrBooleanExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrBooleanExpressionContext {
	var p = new(OrBooleanExpressionContext)

	InitEmptyBoolean_expressionContext(&p.Boolean_expressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Boolean_expressionContext))

	return p
}

func (s *OrBooleanExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrBooleanExpressionContext) AllBoolean_expression() []IBoolean_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBoolean_expressionContext); ok {
			len++
		}
	}

	tst := make([]IBoolean_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBoolean_expressionContext); ok {
			tst[i] = t.(IBoolean_expressionContext)
			i++
		}
	}

	return tst
}

func (s *OrBooleanExpressionContext) Boolean_expression(i int) IBoolean_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolean_expressionContext); ok {
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

	return t.(IBoolean_expressionContext)
}

func (s *OrBooleanExpressionContext) T_OR() antlr.TerminalNode {
	return s.GetToken(unoParserT_OR, 0)
}

func (s *OrBooleanExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(unoListener); ok {
		listenerT.EnterOrBooleanExpression(s)
	}
}

func (s *OrBooleanExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(unoListener); ok {
		listenerT.ExitOrBooleanExpression(s)
	}
}

type TrueBooleanExpressionContext struct {
	Boolean_expressionContext
}

func NewTrueBooleanExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TrueBooleanExpressionContext {
	var p = new(TrueBooleanExpressionContext)

	InitEmptyBoolean_expressionContext(&p.Boolean_expressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Boolean_expressionContext))

	return p
}

func (s *TrueBooleanExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TrueBooleanExpressionContext) T_TRUE() antlr.TerminalNode {
	return s.GetToken(unoParserT_TRUE, 0)
}

func (s *TrueBooleanExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(unoListener); ok {
		listenerT.EnterTrueBooleanExpression(s)
	}
}

func (s *TrueBooleanExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(unoListener); ok {
		listenerT.ExitTrueBooleanExpression(s)
	}
}

type AndBooleanExpressionContext struct {
	Boolean_expressionContext
}

func NewAndBooleanExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndBooleanExpressionContext {
	var p = new(AndBooleanExpressionContext)

	InitEmptyBoolean_expressionContext(&p.Boolean_expressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Boolean_expressionContext))

	return p
}

func (s *AndBooleanExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndBooleanExpressionContext) AllBoolean_expression() []IBoolean_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBoolean_expressionContext); ok {
			len++
		}
	}

	tst := make([]IBoolean_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBoolean_expressionContext); ok {
			tst[i] = t.(IBoolean_expressionContext)
			i++
		}
	}

	return tst
}

func (s *AndBooleanExpressionContext) Boolean_expression(i int) IBoolean_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolean_expressionContext); ok {
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

	return t.(IBoolean_expressionContext)
}

func (s *AndBooleanExpressionContext) T_AND() antlr.TerminalNode {
	return s.GetToken(unoParserT_AND, 0)
}

func (s *AndBooleanExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(unoListener); ok {
		listenerT.EnterAndBooleanExpression(s)
	}
}

func (s *AndBooleanExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(unoListener); ok {
		listenerT.ExitAndBooleanExpression(s)
	}
}

type NotInBooleanExpressionContext struct {
	Boolean_expressionContext
}

func NewNotInBooleanExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotInBooleanExpressionContext {
	var p = new(NotInBooleanExpressionContext)

	InitEmptyBoolean_expressionContext(&p.Boolean_expressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Boolean_expressionContext))

	return p
}

func (s *NotInBooleanExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotInBooleanExpressionContext) Arithmetic_expression() IArithmetic_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArithmetic_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArithmetic_expressionContext)
}

func (s *NotInBooleanExpressionContext) T_NOT() antlr.TerminalNode {
	return s.GetToken(unoParserT_NOT, 0)
}

func (s *NotInBooleanExpressionContext) T_IN() antlr.TerminalNode {
	return s.GetToken(unoParserT_IN, 0)
}

func (s *NotInBooleanExpressionContext) INTEGER_LIST() antlr.TerminalNode {
	return s.GetToken(unoParserINTEGER_LIST, 0)
}

func (s *NotInBooleanExpressionContext) STRING_LIST() antlr.TerminalNode {
	return s.GetToken(unoParserSTRING_LIST, 0)
}

func (s *NotInBooleanExpressionContext) DECIMAL_LIST() antlr.TerminalNode {
	return s.GetToken(unoParserDECIMAL_LIST, 0)
}

func (s *NotInBooleanExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(unoListener); ok {
		listenerT.EnterNotInBooleanExpression(s)
	}
}

func (s *NotInBooleanExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(unoListener); ok {
		listenerT.ExitNotInBooleanExpression(s)
	}
}

type FalseBooleanExpressionContext struct {
	Boolean_expressionContext
}

func NewFalseBooleanExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FalseBooleanExpressionContext {
	var p = new(FalseBooleanExpressionContext)

	InitEmptyBoolean_expressionContext(&p.Boolean_expressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Boolean_expressionContext))

	return p
}

func (s *FalseBooleanExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FalseBooleanExpressionContext) T_FALSE() antlr.TerminalNode {
	return s.GetToken(unoParserT_FALSE, 0)
}

func (s *FalseBooleanExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(unoListener); ok {
		listenerT.EnterFalseBooleanExpression(s)
	}
}

func (s *FalseBooleanExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(unoListener); ok {
		listenerT.ExitFalseBooleanExpression(s)
	}
}

type InBooleanExpressionContext struct {
	Boolean_expressionContext
}

func NewInBooleanExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InBooleanExpressionContext {
	var p = new(InBooleanExpressionContext)

	InitEmptyBoolean_expressionContext(&p.Boolean_expressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Boolean_expressionContext))

	return p
}

func (s *InBooleanExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InBooleanExpressionContext) Arithmetic_expression() IArithmetic_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArithmetic_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArithmetic_expressionContext)
}

func (s *InBooleanExpressionContext) T_IN() antlr.TerminalNode {
	return s.GetToken(unoParserT_IN, 0)
}

func (s *InBooleanExpressionContext) INTEGER_LIST() antlr.TerminalNode {
	return s.GetToken(unoParserINTEGER_LIST, 0)
}

func (s *InBooleanExpressionContext) STRING_LIST() antlr.TerminalNode {
	return s.GetToken(unoParserSTRING_LIST, 0)
}

func (s *InBooleanExpressionContext) DECIMAL_LIST() antlr.TerminalNode {
	return s.GetToken(unoParserDECIMAL_LIST, 0)
}

func (s *InBooleanExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(unoListener); ok {
		listenerT.EnterInBooleanExpression(s)
	}
}

func (s *InBooleanExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(unoListener); ok {
		listenerT.ExitInBooleanExpression(s)
	}
}

func (p *unoParser) Boolean_expression() (localctx IBoolean_expressionContext) {
	return p.boolean_expression(0)
}

func (p *unoParser) boolean_expression(_p int) (localctx IBoolean_expressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewBoolean_expressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IBoolean_expressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, unoParserRULE_boolean_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		localctx = NewCmpBooleanExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(12)
			p.arithmetic_expression(0)
		}
		{
			p.SetState(13)
			p.Match(unoParserT_COMPARE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(14)
			p.arithmetic_expression(0)
		}

	case 2:
		localctx = NewNotBooleanExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(16)
			p.Match(unoParserT_NOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(17)
			p.boolean_expression(6)
		}

	case 3:
		localctx = NewInBooleanExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(18)
			p.arithmetic_expression(0)
		}
		{
			p.SetState(19)
			p.Match(unoParserT_IN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(20)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&90194313216) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 4:
		localctx = NewNotInBooleanExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(22)
			p.arithmetic_expression(0)
		}
		{
			p.SetState(23)
			p.Match(unoParserT_NOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(24)
			p.Match(unoParserT_IN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(25)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&90194313216) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 5:
		localctx = NewPlainBooleanExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(27)
			p.Match(unoParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(28)
			p.boolean_expression(0)
		}
		{
			p.SetState(29)
			p.Match(unoParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewTrueBooleanExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(31)
			p.Match(unoParserT_TRUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		localctx = NewFalseBooleanExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(32)
			p.Match(unoParserT_FALSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(41)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAndBooleanExpressionContext(p, NewBoolean_expressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, unoParserRULE_boolean_expression)
				p.SetState(35)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(36)
					p.Match(unoParserT_AND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(37)
					p.boolean_expression(10)
				}

			case 2:
				localctx = NewOrBooleanExpressionContext(p, NewBoolean_expressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, unoParserRULE_boolean_expression)
				p.SetState(38)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(39)
					p.Match(unoParserT_OR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(40)
					p.boolean_expression(9)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(45)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
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
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArithmetic_expressionContext is an interface to support dynamic dispatch.
type IArithmetic_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsArithmetic_expressionContext differentiates from other interfaces.
	IsArithmetic_expressionContext()
}

type Arithmetic_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArithmetic_expressionContext() *Arithmetic_expressionContext {
	var p = new(Arithmetic_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = unoParserRULE_arithmetic_expression
	return p
}

func InitEmptyArithmetic_expressionContext(p *Arithmetic_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = unoParserRULE_arithmetic_expression
}

func (*Arithmetic_expressionContext) IsArithmetic_expressionContext() {}

func NewArithmetic_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Arithmetic_expressionContext {
	var p = new(Arithmetic_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = unoParserRULE_arithmetic_expression

	return p
}

func (s *Arithmetic_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Arithmetic_expressionContext) CopyAll(ctx *Arithmetic_expressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Arithmetic_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Arithmetic_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PlainArithmeticExpressionContext struct {
	Arithmetic_expressionContext
}

func NewPlainArithmeticExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PlainArithmeticExpressionContext {
	var p = new(PlainArithmeticExpressionContext)

	InitEmptyArithmetic_expressionContext(&p.Arithmetic_expressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Arithmetic_expressionContext))

	return p
}

func (s *PlainArithmeticExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PlainArithmeticExpressionContext) Arithmetic_expression() IArithmetic_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArithmetic_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArithmetic_expressionContext)
}

func (s *PlainArithmeticExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(unoListener); ok {
		listenerT.EnterPlainArithmeticExpression(s)
	}
}

func (s *PlainArithmeticExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(unoListener); ok {
		listenerT.ExitPlainArithmeticExpression(s)
	}
}

type AddArithmeticExpressionContext struct {
	Arithmetic_expressionContext
}

func NewAddArithmeticExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddArithmeticExpressionContext {
	var p = new(AddArithmeticExpressionContext)

	InitEmptyArithmetic_expressionContext(&p.Arithmetic_expressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Arithmetic_expressionContext))

	return p
}

func (s *AddArithmeticExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddArithmeticExpressionContext) AllArithmetic_expression() []IArithmetic_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArithmetic_expressionContext); ok {
			len++
		}
	}

	tst := make([]IArithmetic_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArithmetic_expressionContext); ok {
			tst[i] = t.(IArithmetic_expressionContext)
			i++
		}
	}

	return tst
}

func (s *AddArithmeticExpressionContext) Arithmetic_expression(i int) IArithmetic_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArithmetic_expressionContext); ok {
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

	return t.(IArithmetic_expressionContext)
}

func (s *AddArithmeticExpressionContext) T_ADD() antlr.TerminalNode {
	return s.GetToken(unoParserT_ADD, 0)
}

func (s *AddArithmeticExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(unoListener); ok {
		listenerT.EnterAddArithmeticExpression(s)
	}
}

func (s *AddArithmeticExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(unoListener); ok {
		listenerT.ExitAddArithmeticExpression(s)
	}
}

type StringArithmeticExpressionContext struct {
	Arithmetic_expressionContext
}

func NewStringArithmeticExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringArithmeticExpressionContext {
	var p = new(StringArithmeticExpressionContext)

	InitEmptyArithmetic_expressionContext(&p.Arithmetic_expressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Arithmetic_expressionContext))

	return p
}

func (s *StringArithmeticExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringArithmeticExpressionContext) STRING() antlr.TerminalNode {
	return s.GetToken(unoParserSTRING, 0)
}

func (s *StringArithmeticExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(unoListener); ok {
		listenerT.EnterStringArithmeticExpression(s)
	}
}

func (s *StringArithmeticExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(unoListener); ok {
		listenerT.ExitStringArithmeticExpression(s)
	}
}

type IntegerArithmeticExpressionContext struct {
	Arithmetic_expressionContext
}

func NewIntegerArithmeticExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntegerArithmeticExpressionContext {
	var p = new(IntegerArithmeticExpressionContext)

	InitEmptyArithmetic_expressionContext(&p.Arithmetic_expressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Arithmetic_expressionContext))

	return p
}

func (s *IntegerArithmeticExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerArithmeticExpressionContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(unoParserINTEGER, 0)
}

func (s *IntegerArithmeticExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(unoListener); ok {
		listenerT.EnterIntegerArithmeticExpression(s)
	}
}

func (s *IntegerArithmeticExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(unoListener); ok {
		listenerT.ExitIntegerArithmeticExpression(s)
	}
}

type DecimalArithmeticExpressionContext struct {
	Arithmetic_expressionContext
}

func NewDecimalArithmeticExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DecimalArithmeticExpressionContext {
	var p = new(DecimalArithmeticExpressionContext)

	InitEmptyArithmetic_expressionContext(&p.Arithmetic_expressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Arithmetic_expressionContext))

	return p
}

func (s *DecimalArithmeticExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecimalArithmeticExpressionContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(unoParserDECIMAL, 0)
}

func (s *DecimalArithmeticExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(unoListener); ok {
		listenerT.EnterDecimalArithmeticExpression(s)
	}
}

func (s *DecimalArithmeticExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(unoListener); ok {
		listenerT.ExitDecimalArithmeticExpression(s)
	}
}

type FuncArithmeticExpressionContext struct {
	Arithmetic_expressionContext
}

func NewFuncArithmeticExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncArithmeticExpressionContext {
	var p = new(FuncArithmeticExpressionContext)

	InitEmptyArithmetic_expressionContext(&p.Arithmetic_expressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Arithmetic_expressionContext))

	return p
}

func (s *FuncArithmeticExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncArithmeticExpressionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(unoParserIDENTIFIER, 0)
}

func (s *FuncArithmeticExpressionContext) AllArithmetic_expression() []IArithmetic_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArithmetic_expressionContext); ok {
			len++
		}
	}

	tst := make([]IArithmetic_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArithmetic_expressionContext); ok {
			tst[i] = t.(IArithmetic_expressionContext)
			i++
		}
	}

	return tst
}

func (s *FuncArithmeticExpressionContext) Arithmetic_expression(i int) IArithmetic_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArithmetic_expressionContext); ok {
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

	return t.(IArithmetic_expressionContext)
}

func (s *FuncArithmeticExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(unoListener); ok {
		listenerT.EnterFuncArithmeticExpression(s)
	}
}

func (s *FuncArithmeticExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(unoListener); ok {
		listenerT.ExitFuncArithmeticExpression(s)
	}
}

type ColumnArithmeticExpressionContext struct {
	Arithmetic_expressionContext
}

func NewColumnArithmeticExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ColumnArithmeticExpressionContext {
	var p = new(ColumnArithmeticExpressionContext)

	InitEmptyArithmetic_expressionContext(&p.Arithmetic_expressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Arithmetic_expressionContext))

	return p
}

func (s *ColumnArithmeticExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnArithmeticExpressionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(unoParserIDENTIFIER, 0)
}

func (s *ColumnArithmeticExpressionContext) Type_marker() IType_markerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_markerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_markerContext)
}

func (s *ColumnArithmeticExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(unoListener); ok {
		listenerT.EnterColumnArithmeticExpression(s)
	}
}

func (s *ColumnArithmeticExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(unoListener); ok {
		listenerT.ExitColumnArithmeticExpression(s)
	}
}

type DivArithmeticExpressionContext struct {
	Arithmetic_expressionContext
}

func NewDivArithmeticExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DivArithmeticExpressionContext {
	var p = new(DivArithmeticExpressionContext)

	InitEmptyArithmetic_expressionContext(&p.Arithmetic_expressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Arithmetic_expressionContext))

	return p
}

func (s *DivArithmeticExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DivArithmeticExpressionContext) AllArithmetic_expression() []IArithmetic_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArithmetic_expressionContext); ok {
			len++
		}
	}

	tst := make([]IArithmetic_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArithmetic_expressionContext); ok {
			tst[i] = t.(IArithmetic_expressionContext)
			i++
		}
	}

	return tst
}

func (s *DivArithmeticExpressionContext) Arithmetic_expression(i int) IArithmetic_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArithmetic_expressionContext); ok {
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

	return t.(IArithmetic_expressionContext)
}

func (s *DivArithmeticExpressionContext) T_DIV() antlr.TerminalNode {
	return s.GetToken(unoParserT_DIV, 0)
}

func (s *DivArithmeticExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(unoListener); ok {
		listenerT.EnterDivArithmeticExpression(s)
	}
}

func (s *DivArithmeticExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(unoListener); ok {
		listenerT.ExitDivArithmeticExpression(s)
	}
}

type FieldColumnArithmeticExpressionContext struct {
	Arithmetic_expressionContext
}

func NewFieldColumnArithmeticExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FieldColumnArithmeticExpressionContext {
	var p = new(FieldColumnArithmeticExpressionContext)

	InitEmptyArithmetic_expressionContext(&p.Arithmetic_expressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Arithmetic_expressionContext))

	return p
}

func (s *FieldColumnArithmeticExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldColumnArithmeticExpressionContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(unoParserIDENTIFIER)
}

func (s *FieldColumnArithmeticExpressionContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(unoParserIDENTIFIER, i)
}

func (s *FieldColumnArithmeticExpressionContext) Type_marker() IType_markerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_markerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_markerContext)
}

func (s *FieldColumnArithmeticExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(unoListener); ok {
		listenerT.EnterFieldColumnArithmeticExpression(s)
	}
}

func (s *FieldColumnArithmeticExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(unoListener); ok {
		listenerT.ExitFieldColumnArithmeticExpression(s)
	}
}

type SubArithmeticExpressionContext struct {
	Arithmetic_expressionContext
}

func NewSubArithmeticExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SubArithmeticExpressionContext {
	var p = new(SubArithmeticExpressionContext)

	InitEmptyArithmetic_expressionContext(&p.Arithmetic_expressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Arithmetic_expressionContext))

	return p
}

func (s *SubArithmeticExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubArithmeticExpressionContext) AllArithmetic_expression() []IArithmetic_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArithmetic_expressionContext); ok {
			len++
		}
	}

	tst := make([]IArithmetic_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArithmetic_expressionContext); ok {
			tst[i] = t.(IArithmetic_expressionContext)
			i++
		}
	}

	return tst
}

func (s *SubArithmeticExpressionContext) Arithmetic_expression(i int) IArithmetic_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArithmetic_expressionContext); ok {
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

	return t.(IArithmetic_expressionContext)
}

func (s *SubArithmeticExpressionContext) T_SUB() antlr.TerminalNode {
	return s.GetToken(unoParserT_SUB, 0)
}

func (s *SubArithmeticExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(unoListener); ok {
		listenerT.EnterSubArithmeticExpression(s)
	}
}

func (s *SubArithmeticExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(unoListener); ok {
		listenerT.ExitSubArithmeticExpression(s)
	}
}

type ModArithmeticExpressionContext struct {
	Arithmetic_expressionContext
}

func NewModArithmeticExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ModArithmeticExpressionContext {
	var p = new(ModArithmeticExpressionContext)

	InitEmptyArithmetic_expressionContext(&p.Arithmetic_expressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Arithmetic_expressionContext))

	return p
}

func (s *ModArithmeticExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModArithmeticExpressionContext) AllArithmetic_expression() []IArithmetic_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArithmetic_expressionContext); ok {
			len++
		}
	}

	tst := make([]IArithmetic_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArithmetic_expressionContext); ok {
			tst[i] = t.(IArithmetic_expressionContext)
			i++
		}
	}

	return tst
}

func (s *ModArithmeticExpressionContext) Arithmetic_expression(i int) IArithmetic_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArithmetic_expressionContext); ok {
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

	return t.(IArithmetic_expressionContext)
}

func (s *ModArithmeticExpressionContext) T_MOD() antlr.TerminalNode {
	return s.GetToken(unoParserT_MOD, 0)
}

func (s *ModArithmeticExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(unoListener); ok {
		listenerT.EnterModArithmeticExpression(s)
	}
}

func (s *ModArithmeticExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(unoListener); ok {
		listenerT.ExitModArithmeticExpression(s)
	}
}

type RuntTimeFuncArithmeticExpressionContext struct {
	Arithmetic_expressionContext
}

func NewRuntTimeFuncArithmeticExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RuntTimeFuncArithmeticExpressionContext {
	var p = new(RuntTimeFuncArithmeticExpressionContext)

	InitEmptyArithmetic_expressionContext(&p.Arithmetic_expressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Arithmetic_expressionContext))

	return p
}

func (s *RuntTimeFuncArithmeticExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuntTimeFuncArithmeticExpressionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(unoParserIDENTIFIER, 0)
}

func (s *RuntTimeFuncArithmeticExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(unoListener); ok {
		listenerT.EnterRuntTimeFuncArithmeticExpression(s)
	}
}

func (s *RuntTimeFuncArithmeticExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(unoListener); ok {
		listenerT.ExitRuntTimeFuncArithmeticExpression(s)
	}
}

type MulArithmeticExpressionContext struct {
	Arithmetic_expressionContext
}

func NewMulArithmeticExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulArithmeticExpressionContext {
	var p = new(MulArithmeticExpressionContext)

	InitEmptyArithmetic_expressionContext(&p.Arithmetic_expressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Arithmetic_expressionContext))

	return p
}

func (s *MulArithmeticExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulArithmeticExpressionContext) AllArithmetic_expression() []IArithmetic_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArithmetic_expressionContext); ok {
			len++
		}
	}

	tst := make([]IArithmetic_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArithmetic_expressionContext); ok {
			tst[i] = t.(IArithmetic_expressionContext)
			i++
		}
	}

	return tst
}

func (s *MulArithmeticExpressionContext) Arithmetic_expression(i int) IArithmetic_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArithmetic_expressionContext); ok {
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

	return t.(IArithmetic_expressionContext)
}

func (s *MulArithmeticExpressionContext) T_MUL() antlr.TerminalNode {
	return s.GetToken(unoParserT_MUL, 0)
}

func (s *MulArithmeticExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(unoListener); ok {
		listenerT.EnterMulArithmeticExpression(s)
	}
}

func (s *MulArithmeticExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(unoListener); ok {
		listenerT.ExitMulArithmeticExpression(s)
	}
}

func (p *unoParser) Arithmetic_expression() (localctx IArithmetic_expressionContext) {
	return p.arithmetic_expression(0)
}

func (p *unoParser) arithmetic_expression(_p int) (localctx IArithmetic_expressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewArithmetic_expressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IArithmetic_expressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, unoParserRULE_arithmetic_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		localctx = NewRuntTimeFuncArithmeticExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(47)
			p.Match(unoParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(48)
			p.Match(unoParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(49)
			p.Match(unoParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewFuncArithmeticExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(50)
			p.Match(unoParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(51)
			p.Match(unoParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(52)
			p.arithmetic_expression(0)
		}
		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == unoParserT__2 {
			{
				p.SetState(53)
				p.Match(unoParserT__2)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(54)
				p.arithmetic_expression(0)
			}

			p.SetState(59)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(60)
			p.Match(unoParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewColumnArithmeticExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(62)
			p.Match(unoParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(63)
			p.Type_marker()
		}

	case 4:
		localctx = NewFieldColumnArithmeticExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(64)
			p.Match(unoParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(65)
			p.Match(unoParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(66)
			p.Match(unoParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(67)
			p.Type_marker()
		}

	case 5:
		localctx = NewStringArithmeticExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(68)
			p.Match(unoParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewIntegerArithmeticExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(69)
			p.Match(unoParserINTEGER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		localctx = NewDecimalArithmeticExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(70)
			p.Match(unoParserDECIMAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		localctx = NewPlainArithmeticExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(71)
			p.Match(unoParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(72)
			p.arithmetic_expression(0)
		}
		{
			p.SetState(73)
			p.Match(unoParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(92)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
			case 1:
				localctx = NewModArithmeticExpressionContext(p, NewArithmetic_expressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, unoParserRULE_arithmetic_expression)
				p.SetState(77)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(78)
					p.Match(unoParserT_MOD)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(79)
					p.arithmetic_expression(14)
				}

			case 2:
				localctx = NewMulArithmeticExpressionContext(p, NewArithmetic_expressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, unoParserRULE_arithmetic_expression)
				p.SetState(80)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(81)
					p.Match(unoParserT_MUL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(82)
					p.arithmetic_expression(13)
				}

			case 3:
				localctx = NewDivArithmeticExpressionContext(p, NewArithmetic_expressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, unoParserRULE_arithmetic_expression)
				p.SetState(83)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(84)
					p.Match(unoParserT_DIV)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(85)
					p.arithmetic_expression(12)
				}

			case 4:
				localctx = NewAddArithmeticExpressionContext(p, NewArithmetic_expressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, unoParserRULE_arithmetic_expression)
				p.SetState(86)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(87)
					p.Match(unoParserT_ADD)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(88)
					p.arithmetic_expression(11)
				}

			case 5:
				localctx = NewSubArithmeticExpressionContext(p, NewArithmetic_expressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, unoParserRULE_arithmetic_expression)
				p.SetState(89)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(90)
					p.Match(unoParserT_SUB)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(91)
					p.arithmetic_expression(10)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(96)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
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
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IType_markerContext is an interface to support dynamic dispatch.
type IType_markerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	T_INT() antlr.TerminalNode
	T_FLOAT() antlr.TerminalNode
	T_STRING() antlr.TerminalNode

	// IsType_markerContext differentiates from other interfaces.
	IsType_markerContext()
}

type Type_markerContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_markerContext() *Type_markerContext {
	var p = new(Type_markerContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = unoParserRULE_type_marker
	return p
}

func InitEmptyType_markerContext(p *Type_markerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = unoParserRULE_type_marker
}

func (*Type_markerContext) IsType_markerContext() {}

func NewType_markerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_markerContext {
	var p = new(Type_markerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = unoParserRULE_type_marker

	return p
}

func (s *Type_markerContext) GetParser() antlr.Parser { return s.parser }

func (s *Type_markerContext) T_INT() antlr.TerminalNode {
	return s.GetToken(unoParserT_INT, 0)
}

func (s *Type_markerContext) T_FLOAT() antlr.TerminalNode {
	return s.GetToken(unoParserT_FLOAT, 0)
}

func (s *Type_markerContext) T_STRING() antlr.TerminalNode {
	return s.GetToken(unoParserT_STRING, 0)
}

func (s *Type_markerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_markerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_markerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(unoListener); ok {
		listenerT.EnterType_marker(s)
	}
}

func (s *Type_markerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(unoListener); ok {
		listenerT.ExitType_marker(s)
	}
}

func (p *unoParser) Type_marker() (localctx IType_markerContext) {
	localctx = NewType_markerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, unoParserRULE_type_marker)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(97)
		p.Match(unoParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(98)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&28672) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(99)
		p.Match(unoParserT__5)
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

func (p *unoParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *Boolean_expressionContext = nil
		if localctx != nil {
			t = localctx.(*Boolean_expressionContext)
		}
		return p.Boolean_expression_Sempred(t, predIndex)

	case 2:
		var t *Arithmetic_expressionContext = nil
		if localctx != nil {
			t = localctx.(*Arithmetic_expressionContext)
		}
		return p.Arithmetic_expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *unoParser) Boolean_expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 8)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *unoParser) Arithmetic_expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 9)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
