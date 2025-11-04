// Code generated from minia.g4 by ANTLR 4.13.2. DO NOT EDIT.

package minia // minia
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

type miniaParser struct {
	*antlr.BaseParser
}

var MiniaParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func miniaParserInit() {
	staticData := &MiniaParserStaticData
	staticData.LiteralNames = []string{
		"", "'('", "')'", "'&'", "'|'", "'!'", "'true'", "'false'", "'=='",
		"'!='", "'>'", "'>='", "'<'", "'<='", "'+'", "'-'", "'*'", "'/'", "'%'",
		"','", "';'", "'='", "'\"'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "T_AND", "T_OR", "T_NOT", "T_TRUE", "T_FALSE", "T_EQ", "T_NEQ",
		"T_GT", "T_GTE", "T_LT", "T_LTE", "T_ADD", "T_SUB", "T_MUL", "T_DIV",
		"T_MOD", "COMMA", "SEMI", "T_EQUAL", "QUOTA", "STRING_LIST", "INTEGER_LIST",
		"DECIMAL_LIST", "IDENTIFIER", "INTEGER", "DECIMAL", "STRING", "WS",
	}
	staticData.RuleNames = []string{
		"prog", "start", "expr", "logicalOrExpr", "logicalAndExpr", "equalityExpr",
		"relationalExpr", "additiveExpr", "multiplicativeExpr", "unaryExpr",
		"primaryExpr", "listLiteral", "funcCall", "exprList", "literal",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 30, 170, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 1, 0, 1, 0,
		1, 0, 5, 0, 34, 8, 0, 10, 0, 12, 0, 37, 9, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 54,
		8, 3, 10, 3, 12, 3, 57, 9, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4,
		65, 8, 4, 10, 4, 12, 4, 68, 9, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 3, 5, 79, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 98, 8,
		6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 109, 8,
		7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 3, 8, 124, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 131, 8, 9,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3,
		10, 143, 8, 10, 1, 11, 1, 11, 1, 11, 3, 11, 148, 8, 11, 1, 12, 1, 12, 1,
		12, 3, 12, 153, 8, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 5, 13, 160, 8,
		13, 10, 13, 12, 13, 163, 9, 13, 1, 14, 1, 14, 1, 14, 3, 14, 168, 8, 14,
		1, 14, 0, 2, 6, 8, 15, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26,
		28, 0, 0, 182, 0, 35, 1, 0, 0, 0, 2, 41, 1, 0, 0, 0, 4, 45, 1, 0, 0, 0,
		6, 47, 1, 0, 0, 0, 8, 58, 1, 0, 0, 0, 10, 78, 1, 0, 0, 0, 12, 97, 1, 0,
		0, 0, 14, 108, 1, 0, 0, 0, 16, 123, 1, 0, 0, 0, 18, 130, 1, 0, 0, 0, 20,
		142, 1, 0, 0, 0, 22, 147, 1, 0, 0, 0, 24, 149, 1, 0, 0, 0, 26, 156, 1,
		0, 0, 0, 28, 167, 1, 0, 0, 0, 30, 31, 3, 2, 1, 0, 31, 32, 5, 20, 0, 0,
		32, 34, 1, 0, 0, 0, 33, 30, 1, 0, 0, 0, 34, 37, 1, 0, 0, 0, 35, 33, 1,
		0, 0, 0, 35, 36, 1, 0, 0, 0, 36, 38, 1, 0, 0, 0, 37, 35, 1, 0, 0, 0, 38,
		39, 3, 2, 1, 0, 39, 40, 5, 0, 0, 1, 40, 1, 1, 0, 0, 0, 41, 42, 5, 26, 0,
		0, 42, 43, 5, 21, 0, 0, 43, 44, 3, 4, 2, 0, 44, 3, 1, 0, 0, 0, 45, 46,
		3, 6, 3, 0, 46, 5, 1, 0, 0, 0, 47, 48, 6, 3, -1, 0, 48, 49, 3, 8, 4, 0,
		49, 55, 1, 0, 0, 0, 50, 51, 10, 2, 0, 0, 51, 52, 5, 4, 0, 0, 52, 54, 3,
		6, 3, 3, 53, 50, 1, 0, 0, 0, 54, 57, 1, 0, 0, 0, 55, 53, 1, 0, 0, 0, 55,
		56, 1, 0, 0, 0, 56, 7, 1, 0, 0, 0, 57, 55, 1, 0, 0, 0, 58, 59, 6, 4, -1,
		0, 59, 60, 3, 10, 5, 0, 60, 66, 1, 0, 0, 0, 61, 62, 10, 2, 0, 0, 62, 63,
		5, 3, 0, 0, 63, 65, 3, 8, 4, 3, 64, 61, 1, 0, 0, 0, 65, 68, 1, 0, 0, 0,
		66, 64, 1, 0, 0, 0, 66, 67, 1, 0, 0, 0, 67, 9, 1, 0, 0, 0, 68, 66, 1, 0,
		0, 0, 69, 70, 3, 12, 6, 0, 70, 71, 5, 8, 0, 0, 71, 72, 3, 12, 6, 0, 72,
		79, 1, 0, 0, 0, 73, 74, 3, 12, 6, 0, 74, 75, 5, 9, 0, 0, 75, 76, 3, 12,
		6, 0, 76, 79, 1, 0, 0, 0, 77, 79, 3, 12, 6, 0, 78, 69, 1, 0, 0, 0, 78,
		73, 1, 0, 0, 0, 78, 77, 1, 0, 0, 0, 79, 11, 1, 0, 0, 0, 80, 81, 3, 14,
		7, 0, 81, 82, 5, 10, 0, 0, 82, 83, 3, 14, 7, 0, 83, 98, 1, 0, 0, 0, 84,
		85, 3, 14, 7, 0, 85, 86, 5, 11, 0, 0, 86, 87, 3, 14, 7, 0, 87, 98, 1, 0,
		0, 0, 88, 89, 3, 14, 7, 0, 89, 90, 5, 12, 0, 0, 90, 91, 3, 14, 7, 0, 91,
		98, 1, 0, 0, 0, 92, 93, 3, 14, 7, 0, 93, 94, 5, 13, 0, 0, 94, 95, 3, 14,
		7, 0, 95, 98, 1, 0, 0, 0, 96, 98, 3, 14, 7, 0, 97, 80, 1, 0, 0, 0, 97,
		84, 1, 0, 0, 0, 97, 88, 1, 0, 0, 0, 97, 92, 1, 0, 0, 0, 97, 96, 1, 0, 0,
		0, 98, 13, 1, 0, 0, 0, 99, 100, 3, 16, 8, 0, 100, 101, 5, 14, 0, 0, 101,
		102, 3, 16, 8, 0, 102, 109, 1, 0, 0, 0, 103, 104, 3, 16, 8, 0, 104, 105,
		5, 15, 0, 0, 105, 106, 3, 16, 8, 0, 106, 109, 1, 0, 0, 0, 107, 109, 3,
		16, 8, 0, 108, 99, 1, 0, 0, 0, 108, 103, 1, 0, 0, 0, 108, 107, 1, 0, 0,
		0, 109, 15, 1, 0, 0, 0, 110, 111, 3, 18, 9, 0, 111, 112, 5, 16, 0, 0, 112,
		113, 3, 18, 9, 0, 113, 124, 1, 0, 0, 0, 114, 115, 3, 18, 9, 0, 115, 116,
		5, 17, 0, 0, 116, 117, 3, 18, 9, 0, 117, 124, 1, 0, 0, 0, 118, 119, 3,
		18, 9, 0, 119, 120, 5, 18, 0, 0, 120, 121, 3, 18, 9, 0, 121, 124, 1, 0,
		0, 0, 122, 124, 3, 18, 9, 0, 123, 110, 1, 0, 0, 0, 123, 114, 1, 0, 0, 0,
		123, 118, 1, 0, 0, 0, 123, 122, 1, 0, 0, 0, 124, 17, 1, 0, 0, 0, 125, 126,
		5, 5, 0, 0, 126, 131, 3, 18, 9, 0, 127, 128, 5, 15, 0, 0, 128, 131, 3,
		18, 9, 0, 129, 131, 3, 20, 10, 0, 130, 125, 1, 0, 0, 0, 130, 127, 1, 0,
		0, 0, 130, 129, 1, 0, 0, 0, 131, 19, 1, 0, 0, 0, 132, 133, 5, 1, 0, 0,
		133, 134, 3, 4, 2, 0, 134, 135, 5, 2, 0, 0, 135, 143, 1, 0, 0, 0, 136,
		143, 3, 24, 12, 0, 137, 143, 5, 26, 0, 0, 138, 143, 3, 28, 14, 0, 139,
		143, 3, 22, 11, 0, 140, 143, 5, 6, 0, 0, 141, 143, 5, 7, 0, 0, 142, 132,
		1, 0, 0, 0, 142, 136, 1, 0, 0, 0, 142, 137, 1, 0, 0, 0, 142, 138, 1, 0,
		0, 0, 142, 139, 1, 0, 0, 0, 142, 140, 1, 0, 0, 0, 142, 141, 1, 0, 0, 0,
		143, 21, 1, 0, 0, 0, 144, 148, 5, 23, 0, 0, 145, 148, 5, 24, 0, 0, 146,
		148, 5, 25, 0, 0, 147, 144, 1, 0, 0, 0, 147, 145, 1, 0, 0, 0, 147, 146,
		1, 0, 0, 0, 148, 23, 1, 0, 0, 0, 149, 150, 5, 26, 0, 0, 150, 152, 5, 1,
		0, 0, 151, 153, 3, 26, 13, 0, 152, 151, 1, 0, 0, 0, 152, 153, 1, 0, 0,
		0, 153, 154, 1, 0, 0, 0, 154, 155, 5, 2, 0, 0, 155, 25, 1, 0, 0, 0, 156,
		161, 3, 4, 2, 0, 157, 158, 5, 19, 0, 0, 158, 160, 3, 4, 2, 0, 159, 157,
		1, 0, 0, 0, 160, 163, 1, 0, 0, 0, 161, 159, 1, 0, 0, 0, 161, 162, 1, 0,
		0, 0, 162, 27, 1, 0, 0, 0, 163, 161, 1, 0, 0, 0, 164, 168, 5, 29, 0, 0,
		165, 168, 5, 27, 0, 0, 166, 168, 5, 28, 0, 0, 167, 164, 1, 0, 0, 0, 167,
		165, 1, 0, 0, 0, 167, 166, 1, 0, 0, 0, 168, 29, 1, 0, 0, 0, 13, 35, 55,
		66, 78, 97, 108, 123, 130, 142, 147, 152, 161, 167,
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

// miniaParserInit initializes any static state used to implement miniaParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewminiaParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MiniaParserInit() {
	staticData := &MiniaParserStaticData
	staticData.once.Do(miniaParserInit)
}

// NewminiaParser produces a new parser instance for the optional input antlr.TokenStream.
func NewminiaParser(input antlr.TokenStream) *miniaParser {
	MiniaParserInit()
	this := new(miniaParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &MiniaParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "minia.g4"

	return this
}

// miniaParser tokens.
const (
	miniaParserEOF          = antlr.TokenEOF
	miniaParserT__0         = 1
	miniaParserT__1         = 2
	miniaParserT_AND        = 3
	miniaParserT_OR         = 4
	miniaParserT_NOT        = 5
	miniaParserT_TRUE       = 6
	miniaParserT_FALSE      = 7
	miniaParserT_EQ         = 8
	miniaParserT_NEQ        = 9
	miniaParserT_GT         = 10
	miniaParserT_GTE        = 11
	miniaParserT_LT         = 12
	miniaParserT_LTE        = 13
	miniaParserT_ADD        = 14
	miniaParserT_SUB        = 15
	miniaParserT_MUL        = 16
	miniaParserT_DIV        = 17
	miniaParserT_MOD        = 18
	miniaParserCOMMA        = 19
	miniaParserSEMI         = 20
	miniaParserT_EQUAL      = 21
	miniaParserQUOTA        = 22
	miniaParserSTRING_LIST  = 23
	miniaParserINTEGER_LIST = 24
	miniaParserDECIMAL_LIST = 25
	miniaParserIDENTIFIER   = 26
	miniaParserINTEGER      = 27
	miniaParserDECIMAL      = 28
	miniaParserSTRING       = 29
	miniaParserWS           = 30
)

// miniaParser rules.
const (
	miniaParserRULE_prog               = 0
	miniaParserRULE_start              = 1
	miniaParserRULE_expr               = 2
	miniaParserRULE_logicalOrExpr      = 3
	miniaParserRULE_logicalAndExpr     = 4
	miniaParserRULE_equalityExpr       = 5
	miniaParserRULE_relationalExpr     = 6
	miniaParserRULE_additiveExpr       = 7
	miniaParserRULE_multiplicativeExpr = 8
	miniaParserRULE_unaryExpr          = 9
	miniaParserRULE_primaryExpr        = 10
	miniaParserRULE_listLiteral        = 11
	miniaParserRULE_funcCall           = 12
	miniaParserRULE_exprList           = 13
	miniaParserRULE_literal            = 14
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStart_() []IStartContext
	Start_(i int) IStartContext
	EOF() antlr.TerminalNode
	AllSEMI() []antlr.TerminalNode
	SEMI(i int) antlr.TerminalNode

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniaParserRULE_prog
	return p
}

func InitEmptyProgContext(p *ProgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniaParserRULE_prog
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = miniaParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) AllStart_() []IStartContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStartContext); ok {
			len++
		}
	}

	tst := make([]IStartContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStartContext); ok {
			tst[i] = t.(IStartContext)
			i++
		}
	}

	return tst
}

func (s *ProgContext) Start_(i int) IStartContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStartContext); ok {
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

	return t.(IStartContext)
}

func (s *ProgContext) EOF() antlr.TerminalNode {
	return s.GetToken(miniaParserEOF, 0)
}

func (s *ProgContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(miniaParserSEMI)
}

func (s *ProgContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(miniaParserSEMI, i)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.ExitProg(s)
	}
}

func (p *miniaParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, miniaParserRULE_prog)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(30)
				p.Start_()
			}
			{
				p.SetState(31)
				p.Match(miniaParserSEMI)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(37)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(38)
		p.Start_()
	}
	{
		p.SetState(39)
		p.Match(miniaParserEOF)
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

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	T_EQUAL() antlr.TerminalNode
	Expr() IExprContext

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
	p.RuleIndex = miniaParserRULE_start
	return p
}

func InitEmptyStartContext(p *StartContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniaParserRULE_start
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = miniaParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(miniaParserIDENTIFIER, 0)
}

func (s *StartContext) T_EQUAL() antlr.TerminalNode {
	return s.GetToken(miniaParserT_EQUAL, 0)
}

func (s *StartContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *miniaParser) Start_() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, miniaParserRULE_start)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(41)
		p.Match(miniaParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(42)
		p.Match(miniaParserT_EQUAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(43)
		p.Expr()
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

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LogicalOrExpr() ILogicalOrExprContext

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniaParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniaParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = miniaParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) LogicalOrExpr() ILogicalOrExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogicalOrExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogicalOrExprContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *miniaParser) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, miniaParserRULE_expr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(45)
		p.logicalOrExpr(0)
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

// ILogicalOrExprContext is an interface to support dynamic dispatch.
type ILogicalOrExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLogicalOrExprContext differentiates from other interfaces.
	IsLogicalOrExprContext()
}

type LogicalOrExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicalOrExprContext() *LogicalOrExprContext {
	var p = new(LogicalOrExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniaParserRULE_logicalOrExpr
	return p
}

func InitEmptyLogicalOrExprContext(p *LogicalOrExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniaParserRULE_logicalOrExpr
}

func (*LogicalOrExprContext) IsLogicalOrExprContext() {}

func NewLogicalOrExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalOrExprContext {
	var p = new(LogicalOrExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = miniaParserRULE_logicalOrExpr

	return p
}

func (s *LogicalOrExprContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicalOrExprContext) CopyAll(ctx *LogicalOrExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *LogicalOrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalOrExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type OrExprContext struct {
	LogicalOrExprContext
}

func NewOrExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrExprContext {
	var p = new(OrExprContext)

	InitEmptyLogicalOrExprContext(&p.LogicalOrExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*LogicalOrExprContext))

	return p
}

func (s *OrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExprContext) AllLogicalOrExpr() []ILogicalOrExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILogicalOrExprContext); ok {
			len++
		}
	}

	tst := make([]ILogicalOrExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILogicalOrExprContext); ok {
			tst[i] = t.(ILogicalOrExprContext)
			i++
		}
	}

	return tst
}

func (s *OrExprContext) LogicalOrExpr(i int) ILogicalOrExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogicalOrExprContext); ok {
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

	return t.(ILogicalOrExprContext)
}

func (s *OrExprContext) T_OR() antlr.TerminalNode {
	return s.GetToken(miniaParserT_OR, 0)
}

func (s *OrExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.EnterOrExpr(s)
	}
}

func (s *OrExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.ExitOrExpr(s)
	}
}

type TrivialLogicalAndExprContext struct {
	LogicalOrExprContext
}

func NewTrivialLogicalAndExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TrivialLogicalAndExprContext {
	var p = new(TrivialLogicalAndExprContext)

	InitEmptyLogicalOrExprContext(&p.LogicalOrExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*LogicalOrExprContext))

	return p
}

func (s *TrivialLogicalAndExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TrivialLogicalAndExprContext) LogicalAndExpr() ILogicalAndExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogicalAndExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogicalAndExprContext)
}

func (s *TrivialLogicalAndExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.EnterTrivialLogicalAndExpr(s)
	}
}

func (s *TrivialLogicalAndExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.ExitTrivialLogicalAndExpr(s)
	}
}

func (p *miniaParser) LogicalOrExpr() (localctx ILogicalOrExprContext) {
	return p.logicalOrExpr(0)
}

func (p *miniaParser) logicalOrExpr(_p int) (localctx ILogicalOrExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewLogicalOrExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ILogicalOrExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 6
	p.EnterRecursionRule(localctx, 6, miniaParserRULE_logicalOrExpr, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewTrivialLogicalAndExprContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(48)
		p.logicalAndExpr(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewOrExprContext(p, NewLogicalOrExprContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, miniaParserRULE_logicalOrExpr)
			p.SetState(50)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(51)
				p.Match(miniaParserT_OR)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(52)
				p.logicalOrExpr(3)
			}

		}
		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext())
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

// ILogicalAndExprContext is an interface to support dynamic dispatch.
type ILogicalAndExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLogicalAndExprContext differentiates from other interfaces.
	IsLogicalAndExprContext()
}

type LogicalAndExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicalAndExprContext() *LogicalAndExprContext {
	var p = new(LogicalAndExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniaParserRULE_logicalAndExpr
	return p
}

func InitEmptyLogicalAndExprContext(p *LogicalAndExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniaParserRULE_logicalAndExpr
}

func (*LogicalAndExprContext) IsLogicalAndExprContext() {}

func NewLogicalAndExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalAndExprContext {
	var p = new(LogicalAndExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = miniaParserRULE_logicalAndExpr

	return p
}

func (s *LogicalAndExprContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicalAndExprContext) CopyAll(ctx *LogicalAndExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *LogicalAndExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalAndExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AndExprContext struct {
	LogicalAndExprContext
}

func NewAndExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndExprContext {
	var p = new(AndExprContext)

	InitEmptyLogicalAndExprContext(&p.LogicalAndExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*LogicalAndExprContext))

	return p
}

func (s *AndExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExprContext) AllLogicalAndExpr() []ILogicalAndExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILogicalAndExprContext); ok {
			len++
		}
	}

	tst := make([]ILogicalAndExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILogicalAndExprContext); ok {
			tst[i] = t.(ILogicalAndExprContext)
			i++
		}
	}

	return tst
}

func (s *AndExprContext) LogicalAndExpr(i int) ILogicalAndExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogicalAndExprContext); ok {
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

	return t.(ILogicalAndExprContext)
}

func (s *AndExprContext) T_AND() antlr.TerminalNode {
	return s.GetToken(miniaParserT_AND, 0)
}

func (s *AndExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.EnterAndExpr(s)
	}
}

func (s *AndExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.ExitAndExpr(s)
	}
}

type TrivialEqualityExprContext struct {
	LogicalAndExprContext
}

func NewTrivialEqualityExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TrivialEqualityExprContext {
	var p = new(TrivialEqualityExprContext)

	InitEmptyLogicalAndExprContext(&p.LogicalAndExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*LogicalAndExprContext))

	return p
}

func (s *TrivialEqualityExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TrivialEqualityExprContext) EqualityExpr() IEqualityExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEqualityExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEqualityExprContext)
}

func (s *TrivialEqualityExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.EnterTrivialEqualityExpr(s)
	}
}

func (s *TrivialEqualityExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.ExitTrivialEqualityExpr(s)
	}
}

func (p *miniaParser) LogicalAndExpr() (localctx ILogicalAndExprContext) {
	return p.logicalAndExpr(0)
}

func (p *miniaParser) logicalAndExpr(_p int) (localctx ILogicalAndExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewLogicalAndExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ILogicalAndExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 8
	p.EnterRecursionRule(localctx, 8, miniaParserRULE_logicalAndExpr, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewTrivialEqualityExprContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(59)
		p.EqualityExpr()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(66)
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
			localctx = NewAndExprContext(p, NewLogicalAndExprContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, miniaParserRULE_logicalAndExpr)
			p.SetState(61)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(62)
				p.Match(miniaParserT_AND)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(63)
				p.logicalAndExpr(3)
			}

		}
		p.SetState(68)
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

// IEqualityExprContext is an interface to support dynamic dispatch.
type IEqualityExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsEqualityExprContext differentiates from other interfaces.
	IsEqualityExprContext()
}

type EqualityExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqualityExprContext() *EqualityExprContext {
	var p = new(EqualityExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniaParserRULE_equalityExpr
	return p
}

func InitEmptyEqualityExprContext(p *EqualityExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniaParserRULE_equalityExpr
}

func (*EqualityExprContext) IsEqualityExprContext() {}

func NewEqualityExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqualityExprContext {
	var p = new(EqualityExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = miniaParserRULE_equalityExpr

	return p
}

func (s *EqualityExprContext) GetParser() antlr.Parser { return s.parser }

func (s *EqualityExprContext) CopyAll(ctx *EqualityExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *EqualityExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type EqualExprContext struct {
	EqualityExprContext
}

func NewEqualExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualExprContext {
	var p = new(EqualExprContext)

	InitEmptyEqualityExprContext(&p.EqualityExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*EqualityExprContext))

	return p
}

func (s *EqualExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualExprContext) AllRelationalExpr() []IRelationalExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRelationalExprContext); ok {
			len++
		}
	}

	tst := make([]IRelationalExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRelationalExprContext); ok {
			tst[i] = t.(IRelationalExprContext)
			i++
		}
	}

	return tst
}

func (s *EqualExprContext) RelationalExpr(i int) IRelationalExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationalExprContext); ok {
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

	return t.(IRelationalExprContext)
}

func (s *EqualExprContext) T_EQ() antlr.TerminalNode {
	return s.GetToken(miniaParserT_EQ, 0)
}

func (s *EqualExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.EnterEqualExpr(s)
	}
}

func (s *EqualExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.ExitEqualExpr(s)
	}
}

type NotEqualExprContext struct {
	EqualityExprContext
}

func NewNotEqualExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotEqualExprContext {
	var p = new(NotEqualExprContext)

	InitEmptyEqualityExprContext(&p.EqualityExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*EqualityExprContext))

	return p
}

func (s *NotEqualExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotEqualExprContext) AllRelationalExpr() []IRelationalExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRelationalExprContext); ok {
			len++
		}
	}

	tst := make([]IRelationalExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRelationalExprContext); ok {
			tst[i] = t.(IRelationalExprContext)
			i++
		}
	}

	return tst
}

func (s *NotEqualExprContext) RelationalExpr(i int) IRelationalExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationalExprContext); ok {
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

	return t.(IRelationalExprContext)
}

func (s *NotEqualExprContext) T_NEQ() antlr.TerminalNode {
	return s.GetToken(miniaParserT_NEQ, 0)
}

func (s *NotEqualExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.EnterNotEqualExpr(s)
	}
}

func (s *NotEqualExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.ExitNotEqualExpr(s)
	}
}

type TrivialRelationalExprContext struct {
	EqualityExprContext
}

func NewTrivialRelationalExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TrivialRelationalExprContext {
	var p = new(TrivialRelationalExprContext)

	InitEmptyEqualityExprContext(&p.EqualityExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*EqualityExprContext))

	return p
}

func (s *TrivialRelationalExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TrivialRelationalExprContext) RelationalExpr() IRelationalExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationalExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationalExprContext)
}

func (s *TrivialRelationalExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.EnterTrivialRelationalExpr(s)
	}
}

func (s *TrivialRelationalExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.ExitTrivialRelationalExpr(s)
	}
}

func (p *miniaParser) EqualityExpr() (localctx IEqualityExprContext) {
	localctx = NewEqualityExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, miniaParserRULE_equalityExpr)
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		localctx = NewEqualExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(69)
			p.RelationalExpr()
		}
		{
			p.SetState(70)
			p.Match(miniaParserT_EQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(71)
			p.RelationalExpr()
		}

	case 2:
		localctx = NewNotEqualExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(73)
			p.RelationalExpr()
		}
		{
			p.SetState(74)
			p.Match(miniaParserT_NEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(75)
			p.RelationalExpr()
		}

	case 3:
		localctx = NewTrivialRelationalExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(77)
			p.RelationalExpr()
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

// IRelationalExprContext is an interface to support dynamic dispatch.
type IRelationalExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsRelationalExprContext differentiates from other interfaces.
	IsRelationalExprContext()
}

type RelationalExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationalExprContext() *RelationalExprContext {
	var p = new(RelationalExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniaParserRULE_relationalExpr
	return p
}

func InitEmptyRelationalExprContext(p *RelationalExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniaParserRULE_relationalExpr
}

func (*RelationalExprContext) IsRelationalExprContext() {}

func NewRelationalExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationalExprContext {
	var p = new(RelationalExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = miniaParserRULE_relationalExpr

	return p
}

func (s *RelationalExprContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationalExprContext) CopyAll(ctx *RelationalExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *RelationalExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationalExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TrivialAdditiveExprContext struct {
	RelationalExprContext
}

func NewTrivialAdditiveExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TrivialAdditiveExprContext {
	var p = new(TrivialAdditiveExprContext)

	InitEmptyRelationalExprContext(&p.RelationalExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*RelationalExprContext))

	return p
}

func (s *TrivialAdditiveExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TrivialAdditiveExprContext) AdditiveExpr() IAdditiveExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAdditiveExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAdditiveExprContext)
}

func (s *TrivialAdditiveExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.EnterTrivialAdditiveExpr(s)
	}
}

func (s *TrivialAdditiveExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.ExitTrivialAdditiveExpr(s)
	}
}

type GreaterThanExprContext struct {
	RelationalExprContext
}

func NewGreaterThanExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GreaterThanExprContext {
	var p = new(GreaterThanExprContext)

	InitEmptyRelationalExprContext(&p.RelationalExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*RelationalExprContext))

	return p
}

func (s *GreaterThanExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GreaterThanExprContext) AllAdditiveExpr() []IAdditiveExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAdditiveExprContext); ok {
			len++
		}
	}

	tst := make([]IAdditiveExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAdditiveExprContext); ok {
			tst[i] = t.(IAdditiveExprContext)
			i++
		}
	}

	return tst
}

func (s *GreaterThanExprContext) AdditiveExpr(i int) IAdditiveExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAdditiveExprContext); ok {
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

	return t.(IAdditiveExprContext)
}

func (s *GreaterThanExprContext) T_GT() antlr.TerminalNode {
	return s.GetToken(miniaParserT_GT, 0)
}

func (s *GreaterThanExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.EnterGreaterThanExpr(s)
	}
}

func (s *GreaterThanExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.ExitGreaterThanExpr(s)
	}
}

type LessThanEqualExprContext struct {
	RelationalExprContext
}

func NewLessThanEqualExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LessThanEqualExprContext {
	var p = new(LessThanEqualExprContext)

	InitEmptyRelationalExprContext(&p.RelationalExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*RelationalExprContext))

	return p
}

func (s *LessThanEqualExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LessThanEqualExprContext) AllAdditiveExpr() []IAdditiveExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAdditiveExprContext); ok {
			len++
		}
	}

	tst := make([]IAdditiveExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAdditiveExprContext); ok {
			tst[i] = t.(IAdditiveExprContext)
			i++
		}
	}

	return tst
}

func (s *LessThanEqualExprContext) AdditiveExpr(i int) IAdditiveExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAdditiveExprContext); ok {
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

	return t.(IAdditiveExprContext)
}

func (s *LessThanEqualExprContext) T_LTE() antlr.TerminalNode {
	return s.GetToken(miniaParserT_LTE, 0)
}

func (s *LessThanEqualExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.EnterLessThanEqualExpr(s)
	}
}

func (s *LessThanEqualExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.ExitLessThanEqualExpr(s)
	}
}

type LessThanExprContext struct {
	RelationalExprContext
}

func NewLessThanExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LessThanExprContext {
	var p = new(LessThanExprContext)

	InitEmptyRelationalExprContext(&p.RelationalExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*RelationalExprContext))

	return p
}

func (s *LessThanExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LessThanExprContext) AllAdditiveExpr() []IAdditiveExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAdditiveExprContext); ok {
			len++
		}
	}

	tst := make([]IAdditiveExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAdditiveExprContext); ok {
			tst[i] = t.(IAdditiveExprContext)
			i++
		}
	}

	return tst
}

func (s *LessThanExprContext) AdditiveExpr(i int) IAdditiveExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAdditiveExprContext); ok {
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

	return t.(IAdditiveExprContext)
}

func (s *LessThanExprContext) T_LT() antlr.TerminalNode {
	return s.GetToken(miniaParserT_LT, 0)
}

func (s *LessThanExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.EnterLessThanExpr(s)
	}
}

func (s *LessThanExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.ExitLessThanExpr(s)
	}
}

type GreaterThanEqualExprContext struct {
	RelationalExprContext
}

func NewGreaterThanEqualExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GreaterThanEqualExprContext {
	var p = new(GreaterThanEqualExprContext)

	InitEmptyRelationalExprContext(&p.RelationalExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*RelationalExprContext))

	return p
}

func (s *GreaterThanEqualExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GreaterThanEqualExprContext) AllAdditiveExpr() []IAdditiveExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAdditiveExprContext); ok {
			len++
		}
	}

	tst := make([]IAdditiveExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAdditiveExprContext); ok {
			tst[i] = t.(IAdditiveExprContext)
			i++
		}
	}

	return tst
}

func (s *GreaterThanEqualExprContext) AdditiveExpr(i int) IAdditiveExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAdditiveExprContext); ok {
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

	return t.(IAdditiveExprContext)
}

func (s *GreaterThanEqualExprContext) T_GTE() antlr.TerminalNode {
	return s.GetToken(miniaParserT_GTE, 0)
}

func (s *GreaterThanEqualExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.EnterGreaterThanEqualExpr(s)
	}
}

func (s *GreaterThanEqualExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.ExitGreaterThanEqualExpr(s)
	}
}

func (p *miniaParser) RelationalExpr() (localctx IRelationalExprContext) {
	localctx = NewRelationalExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, miniaParserRULE_relationalExpr)
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		localctx = NewGreaterThanExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(80)
			p.AdditiveExpr()
		}
		{
			p.SetState(81)
			p.Match(miniaParserT_GT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(82)
			p.AdditiveExpr()
		}

	case 2:
		localctx = NewGreaterThanEqualExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(84)
			p.AdditiveExpr()
		}
		{
			p.SetState(85)
			p.Match(miniaParserT_GTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(86)
			p.AdditiveExpr()
		}

	case 3:
		localctx = NewLessThanExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(88)
			p.AdditiveExpr()
		}
		{
			p.SetState(89)
			p.Match(miniaParserT_LT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(90)
			p.AdditiveExpr()
		}

	case 4:
		localctx = NewLessThanEqualExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(92)
			p.AdditiveExpr()
		}
		{
			p.SetState(93)
			p.Match(miniaParserT_LTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(94)
			p.AdditiveExpr()
		}

	case 5:
		localctx = NewTrivialAdditiveExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(96)
			p.AdditiveExpr()
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

// IAdditiveExprContext is an interface to support dynamic dispatch.
type IAdditiveExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAdditiveExprContext differentiates from other interfaces.
	IsAdditiveExprContext()
}

type AdditiveExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdditiveExprContext() *AdditiveExprContext {
	var p = new(AdditiveExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniaParserRULE_additiveExpr
	return p
}

func InitEmptyAdditiveExprContext(p *AdditiveExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniaParserRULE_additiveExpr
}

func (*AdditiveExprContext) IsAdditiveExprContext() {}

func NewAdditiveExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdditiveExprContext {
	var p = new(AdditiveExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = miniaParserRULE_additiveExpr

	return p
}

func (s *AdditiveExprContext) GetParser() antlr.Parser { return s.parser }

func (s *AdditiveExprContext) CopyAll(ctx *AdditiveExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *AdditiveExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TrivialMultiplicativeExprContext struct {
	AdditiveExprContext
}

func NewTrivialMultiplicativeExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TrivialMultiplicativeExprContext {
	var p = new(TrivialMultiplicativeExprContext)

	InitEmptyAdditiveExprContext(&p.AdditiveExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*AdditiveExprContext))

	return p
}

func (s *TrivialMultiplicativeExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TrivialMultiplicativeExprContext) MultiplicativeExpr() IMultiplicativeExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiplicativeExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultiplicativeExprContext)
}

func (s *TrivialMultiplicativeExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.EnterTrivialMultiplicativeExpr(s)
	}
}

func (s *TrivialMultiplicativeExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.ExitTrivialMultiplicativeExpr(s)
	}
}

type SubExprContext struct {
	AdditiveExprContext
}

func NewSubExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SubExprContext {
	var p = new(SubExprContext)

	InitEmptyAdditiveExprContext(&p.AdditiveExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*AdditiveExprContext))

	return p
}

func (s *SubExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubExprContext) AllMultiplicativeExpr() []IMultiplicativeExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMultiplicativeExprContext); ok {
			len++
		}
	}

	tst := make([]IMultiplicativeExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMultiplicativeExprContext); ok {
			tst[i] = t.(IMultiplicativeExprContext)
			i++
		}
	}

	return tst
}

func (s *SubExprContext) MultiplicativeExpr(i int) IMultiplicativeExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiplicativeExprContext); ok {
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

	return t.(IMultiplicativeExprContext)
}

func (s *SubExprContext) T_SUB() antlr.TerminalNode {
	return s.GetToken(miniaParserT_SUB, 0)
}

func (s *SubExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.EnterSubExpr(s)
	}
}

func (s *SubExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.ExitSubExpr(s)
	}
}

type AddExprContext struct {
	AdditiveExprContext
}

func NewAddExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddExprContext {
	var p = new(AddExprContext)

	InitEmptyAdditiveExprContext(&p.AdditiveExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*AdditiveExprContext))

	return p
}

func (s *AddExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddExprContext) AllMultiplicativeExpr() []IMultiplicativeExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMultiplicativeExprContext); ok {
			len++
		}
	}

	tst := make([]IMultiplicativeExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMultiplicativeExprContext); ok {
			tst[i] = t.(IMultiplicativeExprContext)
			i++
		}
	}

	return tst
}

func (s *AddExprContext) MultiplicativeExpr(i int) IMultiplicativeExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiplicativeExprContext); ok {
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

	return t.(IMultiplicativeExprContext)
}

func (s *AddExprContext) T_ADD() antlr.TerminalNode {
	return s.GetToken(miniaParserT_ADD, 0)
}

func (s *AddExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.EnterAddExpr(s)
	}
}

func (s *AddExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.ExitAddExpr(s)
	}
}

func (p *miniaParser) AdditiveExpr() (localctx IAdditiveExprContext) {
	localctx = NewAdditiveExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, miniaParserRULE_additiveExpr)
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAddExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(99)
			p.MultiplicativeExpr()
		}
		{
			p.SetState(100)
			p.Match(miniaParserT_ADD)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(101)
			p.MultiplicativeExpr()
		}

	case 2:
		localctx = NewSubExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(103)
			p.MultiplicativeExpr()
		}
		{
			p.SetState(104)
			p.Match(miniaParserT_SUB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(105)
			p.MultiplicativeExpr()
		}

	case 3:
		localctx = NewTrivialMultiplicativeExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(107)
			p.MultiplicativeExpr()
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

// IMultiplicativeExprContext is an interface to support dynamic dispatch.
type IMultiplicativeExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsMultiplicativeExprContext differentiates from other interfaces.
	IsMultiplicativeExprContext()
}

type MultiplicativeExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplicativeExprContext() *MultiplicativeExprContext {
	var p = new(MultiplicativeExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniaParserRULE_multiplicativeExpr
	return p
}

func InitEmptyMultiplicativeExprContext(p *MultiplicativeExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniaParserRULE_multiplicativeExpr
}

func (*MultiplicativeExprContext) IsMultiplicativeExprContext() {}

func NewMultiplicativeExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplicativeExprContext {
	var p = new(MultiplicativeExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = miniaParserRULE_multiplicativeExpr

	return p
}

func (s *MultiplicativeExprContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplicativeExprContext) CopyAll(ctx *MultiplicativeExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *MultiplicativeExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicativeExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MulExprContext struct {
	MultiplicativeExprContext
}

func NewMulExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulExprContext {
	var p = new(MulExprContext)

	InitEmptyMultiplicativeExprContext(&p.MultiplicativeExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*MultiplicativeExprContext))

	return p
}

func (s *MulExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulExprContext) AllUnaryExpr() []IUnaryExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IUnaryExprContext); ok {
			len++
		}
	}

	tst := make([]IUnaryExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IUnaryExprContext); ok {
			tst[i] = t.(IUnaryExprContext)
			i++
		}
	}

	return tst
}

func (s *MulExprContext) UnaryExpr(i int) IUnaryExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryExprContext); ok {
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

	return t.(IUnaryExprContext)
}

func (s *MulExprContext) T_MUL() antlr.TerminalNode {
	return s.GetToken(miniaParserT_MUL, 0)
}

func (s *MulExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.EnterMulExpr(s)
	}
}

func (s *MulExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.ExitMulExpr(s)
	}
}

type DivExprContext struct {
	MultiplicativeExprContext
}

func NewDivExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DivExprContext {
	var p = new(DivExprContext)

	InitEmptyMultiplicativeExprContext(&p.MultiplicativeExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*MultiplicativeExprContext))

	return p
}

func (s *DivExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DivExprContext) AllUnaryExpr() []IUnaryExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IUnaryExprContext); ok {
			len++
		}
	}

	tst := make([]IUnaryExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IUnaryExprContext); ok {
			tst[i] = t.(IUnaryExprContext)
			i++
		}
	}

	return tst
}

func (s *DivExprContext) UnaryExpr(i int) IUnaryExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryExprContext); ok {
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

	return t.(IUnaryExprContext)
}

func (s *DivExprContext) T_DIV() antlr.TerminalNode {
	return s.GetToken(miniaParserT_DIV, 0)
}

func (s *DivExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.EnterDivExpr(s)
	}
}

func (s *DivExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.ExitDivExpr(s)
	}
}

type TrivialUnaryExprContext struct {
	MultiplicativeExprContext
}

func NewTrivialUnaryExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TrivialUnaryExprContext {
	var p = new(TrivialUnaryExprContext)

	InitEmptyMultiplicativeExprContext(&p.MultiplicativeExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*MultiplicativeExprContext))

	return p
}

func (s *TrivialUnaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TrivialUnaryExprContext) UnaryExpr() IUnaryExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryExprContext)
}

func (s *TrivialUnaryExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.EnterTrivialUnaryExpr(s)
	}
}

func (s *TrivialUnaryExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.ExitTrivialUnaryExpr(s)
	}
}

type ModExprContext struct {
	MultiplicativeExprContext
}

func NewModExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ModExprContext {
	var p = new(ModExprContext)

	InitEmptyMultiplicativeExprContext(&p.MultiplicativeExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*MultiplicativeExprContext))

	return p
}

func (s *ModExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModExprContext) AllUnaryExpr() []IUnaryExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IUnaryExprContext); ok {
			len++
		}
	}

	tst := make([]IUnaryExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IUnaryExprContext); ok {
			tst[i] = t.(IUnaryExprContext)
			i++
		}
	}

	return tst
}

func (s *ModExprContext) UnaryExpr(i int) IUnaryExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryExprContext); ok {
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

	return t.(IUnaryExprContext)
}

func (s *ModExprContext) T_MOD() antlr.TerminalNode {
	return s.GetToken(miniaParserT_MOD, 0)
}

func (s *ModExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.EnterModExpr(s)
	}
}

func (s *ModExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.ExitModExpr(s)
	}
}

func (p *miniaParser) MultiplicativeExpr() (localctx IMultiplicativeExprContext) {
	localctx = NewMultiplicativeExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, miniaParserRULE_multiplicativeExpr)
	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		localctx = NewMulExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(110)
			p.UnaryExpr()
		}
		{
			p.SetState(111)
			p.Match(miniaParserT_MUL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(112)
			p.UnaryExpr()
		}

	case 2:
		localctx = NewDivExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(114)
			p.UnaryExpr()
		}
		{
			p.SetState(115)
			p.Match(miniaParserT_DIV)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(116)
			p.UnaryExpr()
		}

	case 3:
		localctx = NewModExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(118)
			p.UnaryExpr()
		}
		{
			p.SetState(119)
			p.Match(miniaParserT_MOD)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(120)
			p.UnaryExpr()
		}

	case 4:
		localctx = NewTrivialUnaryExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(122)
			p.UnaryExpr()
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

// IUnaryExprContext is an interface to support dynamic dispatch.
type IUnaryExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsUnaryExprContext differentiates from other interfaces.
	IsUnaryExprContext()
}

type UnaryExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryExprContext() *UnaryExprContext {
	var p = new(UnaryExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniaParserRULE_unaryExpr
	return p
}

func InitEmptyUnaryExprContext(p *UnaryExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniaParserRULE_unaryExpr
}

func (*UnaryExprContext) IsUnaryExprContext() {}

func NewUnaryExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryExprContext {
	var p = new(UnaryExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = miniaParserRULE_unaryExpr

	return p
}

func (s *UnaryExprContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryExprContext) CopyAll(ctx *UnaryExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *UnaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TrivialPrimaryExprContext struct {
	UnaryExprContext
}

func NewTrivialPrimaryExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TrivialPrimaryExprContext {
	var p = new(TrivialPrimaryExprContext)

	InitEmptyUnaryExprContext(&p.UnaryExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*UnaryExprContext))

	return p
}

func (s *TrivialPrimaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TrivialPrimaryExprContext) PrimaryExpr() IPrimaryExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExprContext)
}

func (s *TrivialPrimaryExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.EnterTrivialPrimaryExpr(s)
	}
}

func (s *TrivialPrimaryExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.ExitTrivialPrimaryExpr(s)
	}
}

type NegExprContext struct {
	UnaryExprContext
}

func NewNegExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NegExprContext {
	var p = new(NegExprContext)

	InitEmptyUnaryExprContext(&p.UnaryExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*UnaryExprContext))

	return p
}

func (s *NegExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NegExprContext) T_SUB() antlr.TerminalNode {
	return s.GetToken(miniaParserT_SUB, 0)
}

func (s *NegExprContext) UnaryExpr() IUnaryExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryExprContext)
}

func (s *NegExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.EnterNegExpr(s)
	}
}

func (s *NegExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.ExitNegExpr(s)
	}
}

type NotExprContext struct {
	UnaryExprContext
}

func NewNotExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotExprContext {
	var p = new(NotExprContext)

	InitEmptyUnaryExprContext(&p.UnaryExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*UnaryExprContext))

	return p
}

func (s *NotExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExprContext) T_NOT() antlr.TerminalNode {
	return s.GetToken(miniaParserT_NOT, 0)
}

func (s *NotExprContext) UnaryExpr() IUnaryExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryExprContext)
}

func (s *NotExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.EnterNotExpr(s)
	}
}

func (s *NotExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.ExitNotExpr(s)
	}
}

func (p *miniaParser) UnaryExpr() (localctx IUnaryExprContext) {
	localctx = NewUnaryExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, miniaParserRULE_unaryExpr)
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case miniaParserT_NOT:
		localctx = NewNotExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(125)
			p.Match(miniaParserT_NOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(126)
			p.UnaryExpr()
		}

	case miniaParserT_SUB:
		localctx = NewNegExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(127)
			p.Match(miniaParserT_SUB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(128)
			p.UnaryExpr()
		}

	case miniaParserT__0, miniaParserT_TRUE, miniaParserT_FALSE, miniaParserSTRING_LIST, miniaParserINTEGER_LIST, miniaParserDECIMAL_LIST, miniaParserIDENTIFIER, miniaParserINTEGER, miniaParserDECIMAL, miniaParserSTRING:
		localctx = NewTrivialPrimaryExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(129)
			p.PrimaryExpr()
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

// IPrimaryExprContext is an interface to support dynamic dispatch.
type IPrimaryExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsPrimaryExprContext differentiates from other interfaces.
	IsPrimaryExprContext()
}

type PrimaryExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExprContext() *PrimaryExprContext {
	var p = new(PrimaryExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniaParserRULE_primaryExpr
	return p
}

func InitEmptyPrimaryExprContext(p *PrimaryExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniaParserRULE_primaryExpr
}

func (*PrimaryExprContext) IsPrimaryExprContext() {}

func NewPrimaryExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExprContext {
	var p = new(PrimaryExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = miniaParserRULE_primaryExpr

	return p
}

func (s *PrimaryExprContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExprContext) CopyAll(ctx *PrimaryExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *PrimaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FunctionCallExprContext struct {
	PrimaryExprContext
}

func NewFunctionCallExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionCallExprContext {
	var p = new(FunctionCallExprContext)

	InitEmptyPrimaryExprContext(&p.PrimaryExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExprContext))

	return p
}

func (s *FunctionCallExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallExprContext) FuncCall() IFuncCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncCallContext)
}

func (s *FunctionCallExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.EnterFunctionCallExpr(s)
	}
}

func (s *FunctionCallExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.ExitFunctionCallExpr(s)
	}
}

type TrueExprContext struct {
	PrimaryExprContext
}

func NewTrueExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TrueExprContext {
	var p = new(TrueExprContext)

	InitEmptyPrimaryExprContext(&p.PrimaryExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExprContext))

	return p
}

func (s *TrueExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TrueExprContext) T_TRUE() antlr.TerminalNode {
	return s.GetToken(miniaParserT_TRUE, 0)
}

func (s *TrueExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.EnterTrueExpr(s)
	}
}

func (s *TrueExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.ExitTrueExpr(s)
	}
}

type ColumnExprContext struct {
	PrimaryExprContext
}

func NewColumnExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ColumnExprContext {
	var p = new(ColumnExprContext)

	InitEmptyPrimaryExprContext(&p.PrimaryExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExprContext))

	return p
}

func (s *ColumnExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnExprContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(miniaParserIDENTIFIER, 0)
}

func (s *ColumnExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.EnterColumnExpr(s)
	}
}

func (s *ColumnExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.ExitColumnExpr(s)
	}
}

type LiteralExprContext struct {
	PrimaryExprContext
}

func NewLiteralExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralExprContext {
	var p = new(LiteralExprContext)

	InitEmptyPrimaryExprContext(&p.PrimaryExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExprContext))

	return p
}

func (s *LiteralExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralExprContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *LiteralExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.EnterLiteralExpr(s)
	}
}

func (s *LiteralExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.ExitLiteralExpr(s)
	}
}

type ParenthesizedExprContext struct {
	PrimaryExprContext
}

func NewParenthesizedExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenthesizedExprContext {
	var p = new(ParenthesizedExprContext)

	InitEmptyPrimaryExprContext(&p.PrimaryExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExprContext))

	return p
}

func (s *ParenthesizedExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenthesizedExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParenthesizedExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.EnterParenthesizedExpr(s)
	}
}

func (s *ParenthesizedExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.ExitParenthesizedExpr(s)
	}
}

type ListExprContext struct {
	PrimaryExprContext
}

func NewListExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListExprContext {
	var p = new(ListExprContext)

	InitEmptyPrimaryExprContext(&p.PrimaryExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExprContext))

	return p
}

func (s *ListExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListExprContext) ListLiteral() IListLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListLiteralContext)
}

func (s *ListExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.EnterListExpr(s)
	}
}

func (s *ListExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.ExitListExpr(s)
	}
}

type FalseExprContext struct {
	PrimaryExprContext
}

func NewFalseExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FalseExprContext {
	var p = new(FalseExprContext)

	InitEmptyPrimaryExprContext(&p.PrimaryExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExprContext))

	return p
}

func (s *FalseExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FalseExprContext) T_FALSE() antlr.TerminalNode {
	return s.GetToken(miniaParserT_FALSE, 0)
}

func (s *FalseExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.EnterFalseExpr(s)
	}
}

func (s *FalseExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.ExitFalseExpr(s)
	}
}

func (p *miniaParser) PrimaryExpr() (localctx IPrimaryExprContext) {
	localctx = NewPrimaryExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, miniaParserRULE_primaryExpr)
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		localctx = NewParenthesizedExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(132)
			p.Match(miniaParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(133)
			p.Expr()
		}
		{
			p.SetState(134)
			p.Match(miniaParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewFunctionCallExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(136)
			p.FuncCall()
		}

	case 3:
		localctx = NewColumnExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(137)
			p.Match(miniaParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewLiteralExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(138)
			p.Literal()
		}

	case 5:
		localctx = NewListExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(139)
			p.ListLiteral()
		}

	case 6:
		localctx = NewTrueExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(140)
			p.Match(miniaParserT_TRUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		localctx = NewFalseExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(141)
			p.Match(miniaParserT_FALSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
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

// IListLiteralContext is an interface to support dynamic dispatch.
type IListLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsListLiteralContext differentiates from other interfaces.
	IsListLiteralContext()
}

type ListLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListLiteralContext() *ListLiteralContext {
	var p = new(ListLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniaParserRULE_listLiteral
	return p
}

func InitEmptyListLiteralContext(p *ListLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniaParserRULE_listLiteral
}

func (*ListLiteralContext) IsListLiteralContext() {}

func NewListLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListLiteralContext {
	var p = new(ListLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = miniaParserRULE_listLiteral

	return p
}

func (s *ListLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *ListLiteralContext) CopyAll(ctx *ListLiteralContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ListLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IntegerListExprContext struct {
	ListLiteralContext
}

func NewIntegerListExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntegerListExprContext {
	var p = new(IntegerListExprContext)

	InitEmptyListLiteralContext(&p.ListLiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*ListLiteralContext))

	return p
}

func (s *IntegerListExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerListExprContext) INTEGER_LIST() antlr.TerminalNode {
	return s.GetToken(miniaParserINTEGER_LIST, 0)
}

func (s *IntegerListExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.EnterIntegerListExpr(s)
	}
}

func (s *IntegerListExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.ExitIntegerListExpr(s)
	}
}

type StringListExprContext struct {
	ListLiteralContext
}

func NewStringListExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringListExprContext {
	var p = new(StringListExprContext)

	InitEmptyListLiteralContext(&p.ListLiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*ListLiteralContext))

	return p
}

func (s *StringListExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringListExprContext) STRING_LIST() antlr.TerminalNode {
	return s.GetToken(miniaParserSTRING_LIST, 0)
}

func (s *StringListExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.EnterStringListExpr(s)
	}
}

func (s *StringListExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.ExitStringListExpr(s)
	}
}

type DecimalListExprContext struct {
	ListLiteralContext
}

func NewDecimalListExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DecimalListExprContext {
	var p = new(DecimalListExprContext)

	InitEmptyListLiteralContext(&p.ListLiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*ListLiteralContext))

	return p
}

func (s *DecimalListExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecimalListExprContext) DECIMAL_LIST() antlr.TerminalNode {
	return s.GetToken(miniaParserDECIMAL_LIST, 0)
}

func (s *DecimalListExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.EnterDecimalListExpr(s)
	}
}

func (s *DecimalListExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.ExitDecimalListExpr(s)
	}
}

func (p *miniaParser) ListLiteral() (localctx IListLiteralContext) {
	localctx = NewListLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, miniaParserRULE_listLiteral)
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case miniaParserSTRING_LIST:
		localctx = NewStringListExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(144)
			p.Match(miniaParserSTRING_LIST)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case miniaParserINTEGER_LIST:
		localctx = NewIntegerListExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(145)
			p.Match(miniaParserINTEGER_LIST)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case miniaParserDECIMAL_LIST:
		localctx = NewDecimalListExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(146)
			p.Match(miniaParserDECIMAL_LIST)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// IFuncCallContext is an interface to support dynamic dispatch.
type IFuncCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	ExprList() IExprListContext

	// IsFuncCallContext differentiates from other interfaces.
	IsFuncCallContext()
}

type FuncCallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncCallContext() *FuncCallContext {
	var p = new(FuncCallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniaParserRULE_funcCall
	return p
}

func InitEmptyFuncCallContext(p *FuncCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniaParserRULE_funcCall
}

func (*FuncCallContext) IsFuncCallContext() {}

func NewFuncCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncCallContext {
	var p = new(FuncCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = miniaParserRULE_funcCall

	return p
}

func (s *FuncCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncCallContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(miniaParserIDENTIFIER, 0)
}

func (s *FuncCallContext) ExprList() IExprListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprListContext)
}

func (s *FuncCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.EnterFuncCall(s)
	}
}

func (s *FuncCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.ExitFuncCall(s)
	}
}

func (p *miniaParser) FuncCall() (localctx IFuncCallContext) {
	localctx = NewFuncCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, miniaParserRULE_funcCall)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(149)
		p.Match(miniaParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(150)
		p.Match(miniaParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1065386210) != 0 {
		{
			p.SetState(151)
			p.ExprList()
		}

	}
	{
		p.SetState(154)
		p.Match(miniaParserT__1)
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

// IExprListContext is an interface to support dynamic dispatch.
type IExprListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsExprListContext differentiates from other interfaces.
	IsExprListContext()
}

type ExprListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprListContext() *ExprListContext {
	var p = new(ExprListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniaParserRULE_exprList
	return p
}

func InitEmptyExprListContext(p *ExprListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniaParserRULE_exprList
}

func (*ExprListContext) IsExprListContext() {}

func NewExprListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprListContext {
	var p = new(ExprListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = miniaParserRULE_exprList

	return p
}

func (s *ExprListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprListContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprListContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *ExprListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(miniaParserCOMMA)
}

func (s *ExprListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(miniaParserCOMMA, i)
}

func (s *ExprListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.EnterExprList(s)
	}
}

func (s *ExprListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.ExitExprList(s)
	}
}

func (p *miniaParser) ExprList() (localctx IExprListContext) {
	localctx = NewExprListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, miniaParserRULE_exprList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(156)
		p.Expr()
	}
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == miniaParserCOMMA {
		{
			p.SetState(157)
			p.Match(miniaParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(158)
			p.Expr()
		}

		p.SetState(163)
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

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniaParserRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = miniaParserRULE_literal
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = miniaParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) CopyAll(ctx *LiteralContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StringExprContext struct {
	LiteralContext
}

func NewStringExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringExprContext {
	var p = new(StringExprContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *StringExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringExprContext) STRING() antlr.TerminalNode {
	return s.GetToken(miniaParserSTRING, 0)
}

func (s *StringExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.EnterStringExpr(s)
	}
}

func (s *StringExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.ExitStringExpr(s)
	}
}

type IntegerExprContext struct {
	LiteralContext
}

func NewIntegerExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntegerExprContext {
	var p = new(IntegerExprContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *IntegerExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerExprContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(miniaParserINTEGER, 0)
}

func (s *IntegerExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.EnterIntegerExpr(s)
	}
}

func (s *IntegerExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.ExitIntegerExpr(s)
	}
}

type DecimalExprContext struct {
	LiteralContext
}

func NewDecimalExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DecimalExprContext {
	var p = new(DecimalExprContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *DecimalExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecimalExprContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(miniaParserDECIMAL, 0)
}

func (s *DecimalExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.EnterDecimalExpr(s)
	}
}

func (s *DecimalExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(miniaListener); ok {
		listenerT.ExitDecimalExpr(s)
	}
}

func (p *miniaParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, miniaParserRULE_literal)
	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case miniaParserSTRING:
		localctx = NewStringExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(164)
			p.Match(miniaParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case miniaParserINTEGER:
		localctx = NewIntegerExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(165)
			p.Match(miniaParserINTEGER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case miniaParserDECIMAL:
		localctx = NewDecimalExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(166)
			p.Match(miniaParserDECIMAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

func (p *miniaParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 3:
		var t *LogicalOrExprContext = nil
		if localctx != nil {
			t = localctx.(*LogicalOrExprContext)
		}
		return p.LogicalOrExpr_Sempred(t, predIndex)

	case 4:
		var t *LogicalAndExprContext = nil
		if localctx != nil {
			t = localctx.(*LogicalAndExprContext)
		}
		return p.LogicalAndExpr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *miniaParser) LogicalOrExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *miniaParser) LogicalAndExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
