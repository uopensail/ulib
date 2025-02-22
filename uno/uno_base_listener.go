// Code generated from uno.g4 by ANTLR 4.13.2. DO NOT EDIT.

package uno // uno
import "github.com/antlr4-go/antlr/v4"

// BaseunoListener is a complete listener for a parse tree produced by unoParser.
type BaseunoListener struct{}

var _ unoListener = &BaseunoListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseunoListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseunoListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseunoListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseunoListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseunoListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseunoListener) ExitStart(ctx *StartContext) {}

// EnterCmpBooleanExpression is called when production CmpBooleanExpression is entered.
func (s *BaseunoListener) EnterCmpBooleanExpression(ctx *CmpBooleanExpressionContext) {}

// ExitCmpBooleanExpression is called when production CmpBooleanExpression is exited.
func (s *BaseunoListener) ExitCmpBooleanExpression(ctx *CmpBooleanExpressionContext) {}

// EnterNotBooleanExpression is called when production NotBooleanExpression is entered.
func (s *BaseunoListener) EnterNotBooleanExpression(ctx *NotBooleanExpressionContext) {}

// ExitNotBooleanExpression is called when production NotBooleanExpression is exited.
func (s *BaseunoListener) ExitNotBooleanExpression(ctx *NotBooleanExpressionContext) {}

// EnterPlainBooleanExpression is called when production PlainBooleanExpression is entered.
func (s *BaseunoListener) EnterPlainBooleanExpression(ctx *PlainBooleanExpressionContext) {}

// ExitPlainBooleanExpression is called when production PlainBooleanExpression is exited.
func (s *BaseunoListener) ExitPlainBooleanExpression(ctx *PlainBooleanExpressionContext) {}

// EnterOrBooleanExpression is called when production OrBooleanExpression is entered.
func (s *BaseunoListener) EnterOrBooleanExpression(ctx *OrBooleanExpressionContext) {}

// ExitOrBooleanExpression is called when production OrBooleanExpression is exited.
func (s *BaseunoListener) ExitOrBooleanExpression(ctx *OrBooleanExpressionContext) {}

// EnterTrueBooleanExpression is called when production TrueBooleanExpression is entered.
func (s *BaseunoListener) EnterTrueBooleanExpression(ctx *TrueBooleanExpressionContext) {}

// ExitTrueBooleanExpression is called when production TrueBooleanExpression is exited.
func (s *BaseunoListener) ExitTrueBooleanExpression(ctx *TrueBooleanExpressionContext) {}

// EnterAndBooleanExpression is called when production AndBooleanExpression is entered.
func (s *BaseunoListener) EnterAndBooleanExpression(ctx *AndBooleanExpressionContext) {}

// ExitAndBooleanExpression is called when production AndBooleanExpression is exited.
func (s *BaseunoListener) ExitAndBooleanExpression(ctx *AndBooleanExpressionContext) {}

// EnterNotInBooleanExpression is called when production NotInBooleanExpression is entered.
func (s *BaseunoListener) EnterNotInBooleanExpression(ctx *NotInBooleanExpressionContext) {}

// ExitNotInBooleanExpression is called when production NotInBooleanExpression is exited.
func (s *BaseunoListener) ExitNotInBooleanExpression(ctx *NotInBooleanExpressionContext) {}

// EnterFalseBooleanExpression is called when production FalseBooleanExpression is entered.
func (s *BaseunoListener) EnterFalseBooleanExpression(ctx *FalseBooleanExpressionContext) {}

// ExitFalseBooleanExpression is called when production FalseBooleanExpression is exited.
func (s *BaseunoListener) ExitFalseBooleanExpression(ctx *FalseBooleanExpressionContext) {}

// EnterInBooleanExpression is called when production InBooleanExpression is entered.
func (s *BaseunoListener) EnterInBooleanExpression(ctx *InBooleanExpressionContext) {}

// ExitInBooleanExpression is called when production InBooleanExpression is exited.
func (s *BaseunoListener) ExitInBooleanExpression(ctx *InBooleanExpressionContext) {}

// EnterPlainArithmeticExpression is called when production PlainArithmeticExpression is entered.
func (s *BaseunoListener) EnterPlainArithmeticExpression(ctx *PlainArithmeticExpressionContext) {}

// ExitPlainArithmeticExpression is called when production PlainArithmeticExpression is exited.
func (s *BaseunoListener) ExitPlainArithmeticExpression(ctx *PlainArithmeticExpressionContext) {}

// EnterAddArithmeticExpression is called when production AddArithmeticExpression is entered.
func (s *BaseunoListener) EnterAddArithmeticExpression(ctx *AddArithmeticExpressionContext) {}

// ExitAddArithmeticExpression is called when production AddArithmeticExpression is exited.
func (s *BaseunoListener) ExitAddArithmeticExpression(ctx *AddArithmeticExpressionContext) {}

// EnterModArithmeticExpression is called when production ModArithmeticExpression is entered.
func (s *BaseunoListener) EnterModArithmeticExpression(ctx *ModArithmeticExpressionContext) {}

// ExitModArithmeticExpression is called when production ModArithmeticExpression is exited.
func (s *BaseunoListener) ExitModArithmeticExpression(ctx *ModArithmeticExpressionContext) {}

// EnterRuntTimeFuncArithmeticExpression is called when production RuntTimeFuncArithmeticExpression is entered.
func (s *BaseunoListener) EnterRuntTimeFuncArithmeticExpression(ctx *RuntTimeFuncArithmeticExpressionContext) {
}

// ExitRuntTimeFuncArithmeticExpression is called when production RuntTimeFuncArithmeticExpression is exited.
func (s *BaseunoListener) ExitRuntTimeFuncArithmeticExpression(ctx *RuntTimeFuncArithmeticExpressionContext) {
}

// EnterStringArithmeticExpression is called when production StringArithmeticExpression is entered.
func (s *BaseunoListener) EnterStringArithmeticExpression(ctx *StringArithmeticExpressionContext) {}

// ExitStringArithmeticExpression is called when production StringArithmeticExpression is exited.
func (s *BaseunoListener) ExitStringArithmeticExpression(ctx *StringArithmeticExpressionContext) {}

// EnterIntegerArithmeticExpression is called when production IntegerArithmeticExpression is entered.
func (s *BaseunoListener) EnterIntegerArithmeticExpression(ctx *IntegerArithmeticExpressionContext) {}

// ExitIntegerArithmeticExpression is called when production IntegerArithmeticExpression is exited.
func (s *BaseunoListener) ExitIntegerArithmeticExpression(ctx *IntegerArithmeticExpressionContext) {}

// EnterDecimalArithmeticExpression is called when production DecimalArithmeticExpression is entered.
func (s *BaseunoListener) EnterDecimalArithmeticExpression(ctx *DecimalArithmeticExpressionContext) {}

// ExitDecimalArithmeticExpression is called when production DecimalArithmeticExpression is exited.
func (s *BaseunoListener) ExitDecimalArithmeticExpression(ctx *DecimalArithmeticExpressionContext) {}

// EnterFuncArithmeticExpression is called when production FuncArithmeticExpression is entered.
func (s *BaseunoListener) EnterFuncArithmeticExpression(ctx *FuncArithmeticExpressionContext) {}

// ExitFuncArithmeticExpression is called when production FuncArithmeticExpression is exited.
func (s *BaseunoListener) ExitFuncArithmeticExpression(ctx *FuncArithmeticExpressionContext) {}

// EnterColumnArithmeticExpression is called when production ColumnArithmeticExpression is entered.
func (s *BaseunoListener) EnterColumnArithmeticExpression(ctx *ColumnArithmeticExpressionContext) {}

// ExitColumnArithmeticExpression is called when production ColumnArithmeticExpression is exited.
func (s *BaseunoListener) ExitColumnArithmeticExpression(ctx *ColumnArithmeticExpressionContext) {}

// EnterDivArithmeticExpression is called when production DivArithmeticExpression is entered.
func (s *BaseunoListener) EnterDivArithmeticExpression(ctx *DivArithmeticExpressionContext) {}

// ExitDivArithmeticExpression is called when production DivArithmeticExpression is exited.
func (s *BaseunoListener) ExitDivArithmeticExpression(ctx *DivArithmeticExpressionContext) {}

// EnterMulArithmeticExpression is called when production MulArithmeticExpression is entered.
func (s *BaseunoListener) EnterMulArithmeticExpression(ctx *MulArithmeticExpressionContext) {}

// ExitMulArithmeticExpression is called when production MulArithmeticExpression is exited.
func (s *BaseunoListener) ExitMulArithmeticExpression(ctx *MulArithmeticExpressionContext) {}

// EnterSubArithmeticExpression is called when production SubArithmeticExpression is entered.
func (s *BaseunoListener) EnterSubArithmeticExpression(ctx *SubArithmeticExpressionContext) {}

// ExitSubArithmeticExpression is called when production SubArithmeticExpression is exited.
func (s *BaseunoListener) ExitSubArithmeticExpression(ctx *SubArithmeticExpressionContext) {}
