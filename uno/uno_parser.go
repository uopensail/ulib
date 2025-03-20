// Code generated from uno.g4 by ANTLR 4.12.0. DO NOT EDIT.

package uno // uno
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type unoParser struct {
	*antlr.BaseParser
}

var unoParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func unoParserInit() {
	staticData := &unoParserStaticData
	staticData.literalNames = []string{
		"", "'('", "')'", "'['", "']'", "'.'", "','", "'\"'", "'+'", "'-'",
		"'*'", "'/'", "'%'", "'on'", "'and'", "'or'", "'not'", "'in'", "'true'",
		"'false'", "", "'='", "'=='", "'<>'", "'!='", "'>'", "'>='", "'<'",
		"'<='",
	}
	staticData.symbolicNames = []string{
		"", "BRACKET_OPEN", "BRACKET_CLOSE", "SQUARE_OPEN", "SQUARE_CLOSE",
		"DOT", "COMMA", "QUOTA", "T_ADD", "T_SUB", "T_MUL", "T_DIV", "T_MOD",
		"T_ON", "T_AND", "T_OR", "T_NOT", "T_IN", "T_TRUE", "T_FALSE", "T_COMPARE",
		"T_EQUAL", "T_EQUAL2", "T_NOTEQUAL", "T_NOTEQUAL2", "T_GREATER", "T_GREATEREQUAL",
		"T_LESS", "T_LESSEQUAL", "IDENTIFIER", "INTEGER_LIST", "INTEGER", "DECIMAL_LIST",
		"DECIMAL", "STRING_LIST", "STRING", "WS",
	}
	staticData.ruleNames = []string{
		"start", "boolean_expression", "arithmetic_expression",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 36, 91, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 1, 0, 1, 0, 1, 0, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 32, 8, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 40, 8, 1, 10, 1, 12, 1, 43, 9,
		1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 54, 8, 2,
		10, 2, 12, 2, 57, 9, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 3, 2, 69, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 86, 8, 2, 10, 2, 12,
		2, 89, 9, 2, 1, 2, 0, 2, 2, 4, 3, 0, 2, 4, 0, 1, 2, 0, 30, 30, 34, 34,
		107, 0, 6, 1, 0, 0, 0, 2, 31, 1, 0, 0, 0, 4, 68, 1, 0, 0, 0, 6, 7, 3, 2,
		1, 0, 7, 8, 5, 0, 0, 1, 8, 1, 1, 0, 0, 0, 9, 10, 6, 1, -1, 0, 10, 11, 3,
		4, 2, 0, 11, 12, 5, 20, 0, 0, 12, 13, 3, 4, 2, 0, 13, 32, 1, 0, 0, 0, 14,
		15, 5, 16, 0, 0, 15, 32, 3, 2, 1, 6, 16, 17, 3, 4, 2, 0, 17, 18, 5, 17,
		0, 0, 18, 19, 7, 0, 0, 0, 19, 32, 1, 0, 0, 0, 20, 21, 3, 4, 2, 0, 21, 22,
		5, 16, 0, 0, 22, 23, 5, 17, 0, 0, 23, 24, 7, 0, 0, 0, 24, 32, 1, 0, 0,
		0, 25, 26, 5, 1, 0, 0, 26, 27, 3, 2, 1, 0, 27, 28, 5, 2, 0, 0, 28, 32,
		1, 0, 0, 0, 29, 32, 5, 18, 0, 0, 30, 32, 5, 19, 0, 0, 31, 9, 1, 0, 0, 0,
		31, 14, 1, 0, 0, 0, 31, 16, 1, 0, 0, 0, 31, 20, 1, 0, 0, 0, 31, 25, 1,
		0, 0, 0, 31, 29, 1, 0, 0, 0, 31, 30, 1, 0, 0, 0, 32, 41, 1, 0, 0, 0, 33,
		34, 10, 9, 0, 0, 34, 35, 5, 14, 0, 0, 35, 40, 3, 2, 1, 10, 36, 37, 10,
		8, 0, 0, 37, 38, 5, 15, 0, 0, 38, 40, 3, 2, 1, 9, 39, 33, 1, 0, 0, 0, 39,
		36, 1, 0, 0, 0, 40, 43, 1, 0, 0, 0, 41, 39, 1, 0, 0, 0, 41, 42, 1, 0, 0,
		0, 42, 3, 1, 0, 0, 0, 43, 41, 1, 0, 0, 0, 44, 45, 6, 2, -1, 0, 45, 46,
		5, 29, 0, 0, 46, 47, 5, 1, 0, 0, 47, 69, 5, 2, 0, 0, 48, 49, 5, 29, 0,
		0, 49, 50, 5, 1, 0, 0, 50, 55, 3, 4, 2, 0, 51, 52, 5, 6, 0, 0, 52, 54,
		3, 4, 2, 0, 53, 51, 1, 0, 0, 0, 54, 57, 1, 0, 0, 0, 55, 53, 1, 0, 0, 0,
		55, 56, 1, 0, 0, 0, 56, 58, 1, 0, 0, 0, 57, 55, 1, 0, 0, 0, 58, 59, 5,
		2, 0, 0, 59, 69, 1, 0, 0, 0, 60, 69, 5, 29, 0, 0, 61, 69, 5, 35, 0, 0,
		62, 69, 5, 31, 0, 0, 63, 69, 5, 33, 0, 0, 64, 65, 5, 1, 0, 0, 65, 66, 3,
		4, 2, 0, 66, 67, 5, 2, 0, 0, 67, 69, 1, 0, 0, 0, 68, 44, 1, 0, 0, 0, 68,
		48, 1, 0, 0, 0, 68, 60, 1, 0, 0, 0, 68, 61, 1, 0, 0, 0, 68, 62, 1, 0, 0,
		0, 68, 63, 1, 0, 0, 0, 68, 64, 1, 0, 0, 0, 69, 87, 1, 0, 0, 0, 70, 71,
		10, 12, 0, 0, 71, 72, 5, 12, 0, 0, 72, 86, 3, 4, 2, 13, 73, 74, 10, 11,
		0, 0, 74, 75, 5, 10, 0, 0, 75, 86, 3, 4, 2, 12, 76, 77, 10, 10, 0, 0, 77,
		78, 5, 11, 0, 0, 78, 86, 3, 4, 2, 11, 79, 80, 10, 9, 0, 0, 80, 81, 5, 8,
		0, 0, 81, 86, 3, 4, 2, 10, 82, 83, 10, 8, 0, 0, 83, 84, 5, 9, 0, 0, 84,
		86, 3, 4, 2, 9, 85, 70, 1, 0, 0, 0, 85, 73, 1, 0, 0, 0, 85, 76, 1, 0, 0,
		0, 85, 79, 1, 0, 0, 0, 85, 82, 1, 0, 0, 0, 86, 89, 1, 0, 0, 0, 87, 85,
		1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 5, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0,
		7, 31, 39, 41, 55, 68, 85, 87,
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
	staticData := &unoParserStaticData
	staticData.once.Do(unoParserInit)
}

// NewunoParser produces a new parser instance for the optional input antlr.TokenStream.
func NewunoParser(input antlr.TokenStream) *unoParser {
	UnoParserInit()
	this := new(unoParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &unoParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "uno.g4"

	return this
}

// unoParser tokens.
const (
	unoParserEOF            = antlr.TokenEOF
	unoParserBRACKET_OPEN   = 1
	unoParserBRACKET_CLOSE  = 2
	unoParserSQUARE_OPEN    = 3
	unoParserSQUARE_CLOSE   = 4
	unoParserDOT            = 5
	unoParserCOMMA          = 6
	unoParserQUOTA          = 7
	unoParserT_ADD          = 8
	unoParserT_SUB          = 9
	unoParserT_MUL          = 10
	unoParserT_DIV          = 11
	unoParserT_MOD          = 12
	unoParserT_ON           = 13
	unoParserT_AND          = 14
	unoParserT_OR           = 15
	unoParserT_NOT          = 16
	unoParserT_IN           = 17
	unoParserT_TRUE         = 18
	unoParserT_FALSE        = 19
	unoParserT_COMPARE      = 20
	unoParserT_EQUAL        = 21
	unoParserT_EQUAL2       = 22
	unoParserT_NOTEQUAL     = 23
	unoParserT_NOTEQUAL2    = 24
	unoParserT_GREATER      = 25
	unoParserT_GREATEREQUAL = 26
	unoParserT_LESS         = 27
	unoParserT_LESSEQUAL    = 28
	unoParserIDENTIFIER     = 29
	unoParserINTEGER_LIST   = 30
	unoParserINTEGER        = 31
	unoParserDECIMAL_LIST   = 32
	unoParserDECIMAL        = 33
	unoParserSTRING_LIST    = 34
	unoParserSTRING         = 35
	unoParserWS             = 36
)

// unoParser rules.
const (
	unoParserRULE_start                 = 0
	unoParserRULE_boolean_expression    = 1
	unoParserRULE_arithmetic_expression = 2
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
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = unoParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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

func (p *unoParser) Start() (localctx IStartContext) {
	this := p
	_ = this

	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, unoParserRULE_start)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(6)
		p.boolean_expression(0)
	}
	{
		p.SetState(7)
		p.Match(unoParserEOF)
	}

	return localctx
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
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolean_expressionContext() *Boolean_expressionContext {
	var p = new(Boolean_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = unoParserRULE_boolean_expression
	return p
}

func (*Boolean_expressionContext) IsBoolean_expressionContext() {}

func NewBoolean_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Boolean_expressionContext {
	var p = new(Boolean_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = unoParserRULE_boolean_expression

	return p
}

func (s *Boolean_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Boolean_expressionContext) CopyFrom(ctx *Boolean_expressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Boolean_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Boolean_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type CmpBooleanExpressionContext struct {
	*Boolean_expressionContext
}

func NewCmpBooleanExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CmpBooleanExpressionContext {
	var p = new(CmpBooleanExpressionContext)

	p.Boolean_expressionContext = NewEmptyBoolean_expressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Boolean_expressionContext))

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
	*Boolean_expressionContext
}

func NewNotBooleanExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotBooleanExpressionContext {
	var p = new(NotBooleanExpressionContext)

	p.Boolean_expressionContext = NewEmptyBoolean_expressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Boolean_expressionContext))

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
	*Boolean_expressionContext
}

func NewPlainBooleanExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PlainBooleanExpressionContext {
	var p = new(PlainBooleanExpressionContext)

	p.Boolean_expressionContext = NewEmptyBoolean_expressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Boolean_expressionContext))

	return p
}

func (s *PlainBooleanExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PlainBooleanExpressionContext) BRACKET_OPEN() antlr.TerminalNode {
	return s.GetToken(unoParserBRACKET_OPEN, 0)
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

func (s *PlainBooleanExpressionContext) BRACKET_CLOSE() antlr.TerminalNode {
	return s.GetToken(unoParserBRACKET_CLOSE, 0)
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
	*Boolean_expressionContext
}

func NewOrBooleanExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrBooleanExpressionContext {
	var p = new(OrBooleanExpressionContext)

	p.Boolean_expressionContext = NewEmptyBoolean_expressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Boolean_expressionContext))

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
	*Boolean_expressionContext
}

func NewTrueBooleanExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TrueBooleanExpressionContext {
	var p = new(TrueBooleanExpressionContext)

	p.Boolean_expressionContext = NewEmptyBoolean_expressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Boolean_expressionContext))

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
	*Boolean_expressionContext
}

func NewAndBooleanExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndBooleanExpressionContext {
	var p = new(AndBooleanExpressionContext)

	p.Boolean_expressionContext = NewEmptyBoolean_expressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Boolean_expressionContext))

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
	*Boolean_expressionContext
}

func NewNotInBooleanExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotInBooleanExpressionContext {
	var p = new(NotInBooleanExpressionContext)

	p.Boolean_expressionContext = NewEmptyBoolean_expressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Boolean_expressionContext))

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
	*Boolean_expressionContext
}

func NewFalseBooleanExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FalseBooleanExpressionContext {
	var p = new(FalseBooleanExpressionContext)

	p.Boolean_expressionContext = NewEmptyBoolean_expressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Boolean_expressionContext))

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
	*Boolean_expressionContext
}

func NewInBooleanExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InBooleanExpressionContext {
	var p = new(InBooleanExpressionContext)

	p.Boolean_expressionContext = NewEmptyBoolean_expressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Boolean_expressionContext))

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
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewBoolean_expressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IBoolean_expressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, unoParserRULE_boolean_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		localctx = NewCmpBooleanExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(10)
			p.arithmetic_expression(0)
		}
		{
			p.SetState(11)
			p.Match(unoParserT_COMPARE)
		}
		{
			p.SetState(12)
			p.arithmetic_expression(0)
		}

	case 2:
		localctx = NewNotBooleanExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(14)
			p.Match(unoParserT_NOT)
		}
		{
			p.SetState(15)
			p.boolean_expression(6)
		}

	case 3:
		localctx = NewInBooleanExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(16)
			p.arithmetic_expression(0)
		}
		{
			p.SetState(17)
			p.Match(unoParserT_IN)
		}
		{
			p.SetState(18)
			_la = p.GetTokenStream().LA(1)

			if !(_la == unoParserINTEGER_LIST || _la == unoParserSTRING_LIST) {
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
			p.SetState(20)
			p.arithmetic_expression(0)
		}
		{
			p.SetState(21)
			p.Match(unoParserT_NOT)
		}
		{
			p.SetState(22)
			p.Match(unoParserT_IN)
		}
		{
			p.SetState(23)
			_la = p.GetTokenStream().LA(1)

			if !(_la == unoParserINTEGER_LIST || _la == unoParserSTRING_LIST) {
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
			p.SetState(25)
			p.Match(unoParserBRACKET_OPEN)
		}
		{
			p.SetState(26)
			p.boolean_expression(0)
		}
		{
			p.SetState(27)
			p.Match(unoParserBRACKET_CLOSE)
		}

	case 6:
		localctx = NewTrueBooleanExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(29)
			p.Match(unoParserT_TRUE)
		}

	case 7:
		localctx = NewFalseBooleanExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(30)
			p.Match(unoParserT_FALSE)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(39)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAndBooleanExpressionContext(p, NewBoolean_expressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, unoParserRULE_boolean_expression)
				p.SetState(33)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(34)
					p.Match(unoParserT_AND)
				}
				{
					p.SetState(35)
					p.boolean_expression(10)
				}

			case 2:
				localctx = NewOrBooleanExpressionContext(p, NewBoolean_expressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, unoParserRULE_boolean_expression)
				p.SetState(36)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(37)
					p.Match(unoParserT_OR)
				}
				{
					p.SetState(38)
					p.boolean_expression(9)
				}

			}

		}
		p.SetState(43)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
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
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArithmetic_expressionContext() *Arithmetic_expressionContext {
	var p = new(Arithmetic_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = unoParserRULE_arithmetic_expression
	return p
}

func (*Arithmetic_expressionContext) IsArithmetic_expressionContext() {}

func NewArithmetic_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Arithmetic_expressionContext {
	var p = new(Arithmetic_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = unoParserRULE_arithmetic_expression

	return p
}

func (s *Arithmetic_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Arithmetic_expressionContext) CopyFrom(ctx *Arithmetic_expressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Arithmetic_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Arithmetic_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PlainArithmeticExpressionContext struct {
	*Arithmetic_expressionContext
}

func NewPlainArithmeticExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PlainArithmeticExpressionContext {
	var p = new(PlainArithmeticExpressionContext)

	p.Arithmetic_expressionContext = NewEmptyArithmetic_expressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Arithmetic_expressionContext))

	return p
}

func (s *PlainArithmeticExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PlainArithmeticExpressionContext) BRACKET_OPEN() antlr.TerminalNode {
	return s.GetToken(unoParserBRACKET_OPEN, 0)
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

func (s *PlainArithmeticExpressionContext) BRACKET_CLOSE() antlr.TerminalNode {
	return s.GetToken(unoParserBRACKET_CLOSE, 0)
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
	*Arithmetic_expressionContext
}

func NewAddArithmeticExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddArithmeticExpressionContext {
	var p = new(AddArithmeticExpressionContext)

	p.Arithmetic_expressionContext = NewEmptyArithmetic_expressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Arithmetic_expressionContext))

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

type ModArithmeticExpressionContext struct {
	*Arithmetic_expressionContext
}

func NewModArithmeticExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ModArithmeticExpressionContext {
	var p = new(ModArithmeticExpressionContext)

	p.Arithmetic_expressionContext = NewEmptyArithmetic_expressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Arithmetic_expressionContext))

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
	*Arithmetic_expressionContext
}

func NewRuntTimeFuncArithmeticExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RuntTimeFuncArithmeticExpressionContext {
	var p = new(RuntTimeFuncArithmeticExpressionContext)

	p.Arithmetic_expressionContext = NewEmptyArithmetic_expressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Arithmetic_expressionContext))

	return p
}

func (s *RuntTimeFuncArithmeticExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuntTimeFuncArithmeticExpressionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(unoParserIDENTIFIER, 0)
}

func (s *RuntTimeFuncArithmeticExpressionContext) BRACKET_OPEN() antlr.TerminalNode {
	return s.GetToken(unoParserBRACKET_OPEN, 0)
}

func (s *RuntTimeFuncArithmeticExpressionContext) BRACKET_CLOSE() antlr.TerminalNode {
	return s.GetToken(unoParserBRACKET_CLOSE, 0)
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

type StringArithmeticExpressionContext struct {
	*Arithmetic_expressionContext
}

func NewStringArithmeticExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringArithmeticExpressionContext {
	var p = new(StringArithmeticExpressionContext)

	p.Arithmetic_expressionContext = NewEmptyArithmetic_expressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Arithmetic_expressionContext))

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
	*Arithmetic_expressionContext
}

func NewIntegerArithmeticExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntegerArithmeticExpressionContext {
	var p = new(IntegerArithmeticExpressionContext)

	p.Arithmetic_expressionContext = NewEmptyArithmetic_expressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Arithmetic_expressionContext))

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
	*Arithmetic_expressionContext
}

func NewDecimalArithmeticExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DecimalArithmeticExpressionContext {
	var p = new(DecimalArithmeticExpressionContext)

	p.Arithmetic_expressionContext = NewEmptyArithmetic_expressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Arithmetic_expressionContext))

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
	*Arithmetic_expressionContext
}

func NewFuncArithmeticExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncArithmeticExpressionContext {
	var p = new(FuncArithmeticExpressionContext)

	p.Arithmetic_expressionContext = NewEmptyArithmetic_expressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Arithmetic_expressionContext))

	return p
}

func (s *FuncArithmeticExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncArithmeticExpressionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(unoParserIDENTIFIER, 0)
}

func (s *FuncArithmeticExpressionContext) BRACKET_OPEN() antlr.TerminalNode {
	return s.GetToken(unoParserBRACKET_OPEN, 0)
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

func (s *FuncArithmeticExpressionContext) BRACKET_CLOSE() antlr.TerminalNode {
	return s.GetToken(unoParserBRACKET_CLOSE, 0)
}

func (s *FuncArithmeticExpressionContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(unoParserCOMMA)
}

func (s *FuncArithmeticExpressionContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(unoParserCOMMA, i)
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
	*Arithmetic_expressionContext
}

func NewColumnArithmeticExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ColumnArithmeticExpressionContext {
	var p = new(ColumnArithmeticExpressionContext)

	p.Arithmetic_expressionContext = NewEmptyArithmetic_expressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Arithmetic_expressionContext))

	return p
}

func (s *ColumnArithmeticExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnArithmeticExpressionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(unoParserIDENTIFIER, 0)
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
	*Arithmetic_expressionContext
}

func NewDivArithmeticExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DivArithmeticExpressionContext {
	var p = new(DivArithmeticExpressionContext)

	p.Arithmetic_expressionContext = NewEmptyArithmetic_expressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Arithmetic_expressionContext))

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

type MulArithmeticExpressionContext struct {
	*Arithmetic_expressionContext
}

func NewMulArithmeticExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulArithmeticExpressionContext {
	var p = new(MulArithmeticExpressionContext)

	p.Arithmetic_expressionContext = NewEmptyArithmetic_expressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Arithmetic_expressionContext))

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

type SubArithmeticExpressionContext struct {
	*Arithmetic_expressionContext
}

func NewSubArithmeticExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SubArithmeticExpressionContext {
	var p = new(SubArithmeticExpressionContext)

	p.Arithmetic_expressionContext = NewEmptyArithmetic_expressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Arithmetic_expressionContext))

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

func (p *unoParser) Arithmetic_expression() (localctx IArithmetic_expressionContext) {
	return p.arithmetic_expression(0)
}

func (p *unoParser) arithmetic_expression(_p int) (localctx IArithmetic_expressionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewArithmetic_expressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IArithmetic_expressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, unoParserRULE_arithmetic_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		localctx = NewRuntTimeFuncArithmeticExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(45)
			p.Match(unoParserIDENTIFIER)
		}
		{
			p.SetState(46)
			p.Match(unoParserBRACKET_OPEN)
		}
		{
			p.SetState(47)
			p.Match(unoParserBRACKET_CLOSE)
		}

	case 2:
		localctx = NewFuncArithmeticExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(48)
			p.Match(unoParserIDENTIFIER)
		}
		{
			p.SetState(49)
			p.Match(unoParserBRACKET_OPEN)
		}
		{
			p.SetState(50)
			p.arithmetic_expression(0)
		}
		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == unoParserCOMMA {
			{
				p.SetState(51)
				p.Match(unoParserCOMMA)
			}
			{
				p.SetState(52)
				p.arithmetic_expression(0)
			}

			p.SetState(57)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(58)
			p.Match(unoParserBRACKET_CLOSE)
		}

	case 3:
		localctx = NewColumnArithmeticExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(60)
			p.Match(unoParserIDENTIFIER)
		}

	case 4:
		localctx = NewStringArithmeticExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(61)
			p.Match(unoParserSTRING)
		}

	case 5:
		localctx = NewIntegerArithmeticExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(62)
			p.Match(unoParserINTEGER)
		}

	case 6:
		localctx = NewDecimalArithmeticExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(63)
			p.Match(unoParserDECIMAL)
		}

	case 7:
		localctx = NewPlainArithmeticExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(64)
			p.Match(unoParserBRACKET_OPEN)
		}
		{
			p.SetState(65)
			p.arithmetic_expression(0)
		}
		{
			p.SetState(66)
			p.Match(unoParserBRACKET_CLOSE)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(85)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
			case 1:
				localctx = NewModArithmeticExpressionContext(p, NewArithmetic_expressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, unoParserRULE_arithmetic_expression)
				p.SetState(70)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(71)
					p.Match(unoParserT_MOD)
				}
				{
					p.SetState(72)
					p.arithmetic_expression(13)
				}

			case 2:
				localctx = NewMulArithmeticExpressionContext(p, NewArithmetic_expressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, unoParserRULE_arithmetic_expression)
				p.SetState(73)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(74)
					p.Match(unoParserT_MUL)
				}
				{
					p.SetState(75)
					p.arithmetic_expression(12)
				}

			case 3:
				localctx = NewDivArithmeticExpressionContext(p, NewArithmetic_expressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, unoParserRULE_arithmetic_expression)
				p.SetState(76)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(77)
					p.Match(unoParserT_DIV)
				}
				{
					p.SetState(78)
					p.arithmetic_expression(11)
				}

			case 4:
				localctx = NewAddArithmeticExpressionContext(p, NewArithmetic_expressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, unoParserRULE_arithmetic_expression)
				p.SetState(79)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(80)
					p.Match(unoParserT_ADD)
				}
				{
					p.SetState(81)
					p.arithmetic_expression(10)
				}

			case 5:
				localctx = NewSubArithmeticExpressionContext(p, NewArithmetic_expressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, unoParserRULE_arithmetic_expression)
				p.SetState(82)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(83)
					p.Match(unoParserT_SUB)
				}
				{
					p.SetState(84)
					p.arithmetic_expression(9)
				}

			}

		}
		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}

	return localctx
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
	this := p
	_ = this

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
	this := p
	_ = this

	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 8)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
