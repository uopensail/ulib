// Code generated from minia.g4 by ANTLR 4.13.2. DO NOT EDIT.

package minia // minia
import "github.com/antlr4-go/antlr/v4"

// BaseminiaListener is a complete listener for a parse tree produced by miniaParser.
type BaseminiaListener struct{}

var _ miniaListener = &BaseminiaListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseminiaListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseminiaListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseminiaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseminiaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseminiaListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseminiaListener) ExitProg(ctx *ProgContext) {}

// EnterStart is called when production start is entered.
func (s *BaseminiaListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseminiaListener) ExitStart(ctx *StartContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseminiaListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseminiaListener) ExitExpr(ctx *ExprContext) {}

// EnterOrExpr is called when production OrExpr is entered.
func (s *BaseminiaListener) EnterOrExpr(ctx *OrExprContext) {}

// ExitOrExpr is called when production OrExpr is exited.
func (s *BaseminiaListener) ExitOrExpr(ctx *OrExprContext) {}

// EnterTrivialLogicalAndExpr is called when production TrivialLogicalAndExpr is entered.
func (s *BaseminiaListener) EnterTrivialLogicalAndExpr(ctx *TrivialLogicalAndExprContext) {}

// ExitTrivialLogicalAndExpr is called when production TrivialLogicalAndExpr is exited.
func (s *BaseminiaListener) ExitTrivialLogicalAndExpr(ctx *TrivialLogicalAndExprContext) {}

// EnterAndExpr is called when production AndExpr is entered.
func (s *BaseminiaListener) EnterAndExpr(ctx *AndExprContext) {}

// ExitAndExpr is called when production AndExpr is exited.
func (s *BaseminiaListener) ExitAndExpr(ctx *AndExprContext) {}

// EnterTrivialEqualityExpr is called when production TrivialEqualityExpr is entered.
func (s *BaseminiaListener) EnterTrivialEqualityExpr(ctx *TrivialEqualityExprContext) {}

// ExitTrivialEqualityExpr is called when production TrivialEqualityExpr is exited.
func (s *BaseminiaListener) ExitTrivialEqualityExpr(ctx *TrivialEqualityExprContext) {}

// EnterEqualExpr is called when production EqualExpr is entered.
func (s *BaseminiaListener) EnterEqualExpr(ctx *EqualExprContext) {}

// ExitEqualExpr is called when production EqualExpr is exited.
func (s *BaseminiaListener) ExitEqualExpr(ctx *EqualExprContext) {}

// EnterNotEqualExpr is called when production NotEqualExpr is entered.
func (s *BaseminiaListener) EnterNotEqualExpr(ctx *NotEqualExprContext) {}

// ExitNotEqualExpr is called when production NotEqualExpr is exited.
func (s *BaseminiaListener) ExitNotEqualExpr(ctx *NotEqualExprContext) {}

// EnterTrivialRelationalExpr is called when production TrivialRelationalExpr is entered.
func (s *BaseminiaListener) EnterTrivialRelationalExpr(ctx *TrivialRelationalExprContext) {}

// ExitTrivialRelationalExpr is called when production TrivialRelationalExpr is exited.
func (s *BaseminiaListener) ExitTrivialRelationalExpr(ctx *TrivialRelationalExprContext) {}

// EnterGreaterThanExpr is called when production GreaterThanExpr is entered.
func (s *BaseminiaListener) EnterGreaterThanExpr(ctx *GreaterThanExprContext) {}

// ExitGreaterThanExpr is called when production GreaterThanExpr is exited.
func (s *BaseminiaListener) ExitGreaterThanExpr(ctx *GreaterThanExprContext) {}

// EnterGreaterThanEqualExpr is called when production GreaterThanEqualExpr is entered.
func (s *BaseminiaListener) EnterGreaterThanEqualExpr(ctx *GreaterThanEqualExprContext) {}

// ExitGreaterThanEqualExpr is called when production GreaterThanEqualExpr is exited.
func (s *BaseminiaListener) ExitGreaterThanEqualExpr(ctx *GreaterThanEqualExprContext) {}

// EnterLessThanExpr is called when production LessThanExpr is entered.
func (s *BaseminiaListener) EnterLessThanExpr(ctx *LessThanExprContext) {}

// ExitLessThanExpr is called when production LessThanExpr is exited.
func (s *BaseminiaListener) ExitLessThanExpr(ctx *LessThanExprContext) {}

// EnterLessThanEqualExpr is called when production LessThanEqualExpr is entered.
func (s *BaseminiaListener) EnterLessThanEqualExpr(ctx *LessThanEqualExprContext) {}

// ExitLessThanEqualExpr is called when production LessThanEqualExpr is exited.
func (s *BaseminiaListener) ExitLessThanEqualExpr(ctx *LessThanEqualExprContext) {}

// EnterTrivialAdditiveExpr is called when production TrivialAdditiveExpr is entered.
func (s *BaseminiaListener) EnterTrivialAdditiveExpr(ctx *TrivialAdditiveExprContext) {}

// ExitTrivialAdditiveExpr is called when production TrivialAdditiveExpr is exited.
func (s *BaseminiaListener) ExitTrivialAdditiveExpr(ctx *TrivialAdditiveExprContext) {}

// EnterAddExpr is called when production AddExpr is entered.
func (s *BaseminiaListener) EnterAddExpr(ctx *AddExprContext) {}

// ExitAddExpr is called when production AddExpr is exited.
func (s *BaseminiaListener) ExitAddExpr(ctx *AddExprContext) {}

// EnterSubExpr is called when production SubExpr is entered.
func (s *BaseminiaListener) EnterSubExpr(ctx *SubExprContext) {}

// ExitSubExpr is called when production SubExpr is exited.
func (s *BaseminiaListener) ExitSubExpr(ctx *SubExprContext) {}

// EnterTrivialMultiplicativeExpr is called when production TrivialMultiplicativeExpr is entered.
func (s *BaseminiaListener) EnterTrivialMultiplicativeExpr(ctx *TrivialMultiplicativeExprContext) {}

// ExitTrivialMultiplicativeExpr is called when production TrivialMultiplicativeExpr is exited.
func (s *BaseminiaListener) ExitTrivialMultiplicativeExpr(ctx *TrivialMultiplicativeExprContext) {}

// EnterMulExpr is called when production MulExpr is entered.
func (s *BaseminiaListener) EnterMulExpr(ctx *MulExprContext) {}

// ExitMulExpr is called when production MulExpr is exited.
func (s *BaseminiaListener) ExitMulExpr(ctx *MulExprContext) {}

// EnterDivExpr is called when production DivExpr is entered.
func (s *BaseminiaListener) EnterDivExpr(ctx *DivExprContext) {}

// ExitDivExpr is called when production DivExpr is exited.
func (s *BaseminiaListener) ExitDivExpr(ctx *DivExprContext) {}

// EnterModExpr is called when production ModExpr is entered.
func (s *BaseminiaListener) EnterModExpr(ctx *ModExprContext) {}

// ExitModExpr is called when production ModExpr is exited.
func (s *BaseminiaListener) ExitModExpr(ctx *ModExprContext) {}

// EnterTrivialUnaryExpr is called when production TrivialUnaryExpr is entered.
func (s *BaseminiaListener) EnterTrivialUnaryExpr(ctx *TrivialUnaryExprContext) {}

// ExitTrivialUnaryExpr is called when production TrivialUnaryExpr is exited.
func (s *BaseminiaListener) ExitTrivialUnaryExpr(ctx *TrivialUnaryExprContext) {}

// EnterNotExpr is called when production NotExpr is entered.
func (s *BaseminiaListener) EnterNotExpr(ctx *NotExprContext) {}

// ExitNotExpr is called when production NotExpr is exited.
func (s *BaseminiaListener) ExitNotExpr(ctx *NotExprContext) {}

// EnterNegExpr is called when production NegExpr is entered.
func (s *BaseminiaListener) EnterNegExpr(ctx *NegExprContext) {}

// ExitNegExpr is called when production NegExpr is exited.
func (s *BaseminiaListener) ExitNegExpr(ctx *NegExprContext) {}

// EnterTrivialPrimaryExpr is called when production TrivialPrimaryExpr is entered.
func (s *BaseminiaListener) EnterTrivialPrimaryExpr(ctx *TrivialPrimaryExprContext) {}

// ExitTrivialPrimaryExpr is called when production TrivialPrimaryExpr is exited.
func (s *BaseminiaListener) ExitTrivialPrimaryExpr(ctx *TrivialPrimaryExprContext) {}

// EnterParenthesizedExpr is called when production ParenthesizedExpr is entered.
func (s *BaseminiaListener) EnterParenthesizedExpr(ctx *ParenthesizedExprContext) {}

// ExitParenthesizedExpr is called when production ParenthesizedExpr is exited.
func (s *BaseminiaListener) ExitParenthesizedExpr(ctx *ParenthesizedExprContext) {}

// EnterFunctionCallExpr is called when production FunctionCallExpr is entered.
func (s *BaseminiaListener) EnterFunctionCallExpr(ctx *FunctionCallExprContext) {}

// ExitFunctionCallExpr is called when production FunctionCallExpr is exited.
func (s *BaseminiaListener) ExitFunctionCallExpr(ctx *FunctionCallExprContext) {}

// EnterColumnExpr is called when production ColumnExpr is entered.
func (s *BaseminiaListener) EnterColumnExpr(ctx *ColumnExprContext) {}

// ExitColumnExpr is called when production ColumnExpr is exited.
func (s *BaseminiaListener) ExitColumnExpr(ctx *ColumnExprContext) {}

// EnterLiteralExpr is called when production LiteralExpr is entered.
func (s *BaseminiaListener) EnterLiteralExpr(ctx *LiteralExprContext) {}

// ExitLiteralExpr is called when production LiteralExpr is exited.
func (s *BaseminiaListener) ExitLiteralExpr(ctx *LiteralExprContext) {}

// EnterListExpr is called when production ListExpr is entered.
func (s *BaseminiaListener) EnterListExpr(ctx *ListExprContext) {}

// ExitListExpr is called when production ListExpr is exited.
func (s *BaseminiaListener) ExitListExpr(ctx *ListExprContext) {}

// EnterTrueExpr is called when production TrueExpr is entered.
func (s *BaseminiaListener) EnterTrueExpr(ctx *TrueExprContext) {}

// ExitTrueExpr is called when production TrueExpr is exited.
func (s *BaseminiaListener) ExitTrueExpr(ctx *TrueExprContext) {}

// EnterFalseExpr is called when production FalseExpr is entered.
func (s *BaseminiaListener) EnterFalseExpr(ctx *FalseExprContext) {}

// ExitFalseExpr is called when production FalseExpr is exited.
func (s *BaseminiaListener) ExitFalseExpr(ctx *FalseExprContext) {}

// EnterStringListExpr is called when production StringListExpr is entered.
func (s *BaseminiaListener) EnterStringListExpr(ctx *StringListExprContext) {}

// ExitStringListExpr is called when production StringListExpr is exited.
func (s *BaseminiaListener) ExitStringListExpr(ctx *StringListExprContext) {}

// EnterIntegerListExpr is called when production IntegerListExpr is entered.
func (s *BaseminiaListener) EnterIntegerListExpr(ctx *IntegerListExprContext) {}

// ExitIntegerListExpr is called when production IntegerListExpr is exited.
func (s *BaseminiaListener) ExitIntegerListExpr(ctx *IntegerListExprContext) {}

// EnterDecimalListExpr is called when production DecimalListExpr is entered.
func (s *BaseminiaListener) EnterDecimalListExpr(ctx *DecimalListExprContext) {}

// ExitDecimalListExpr is called when production DecimalListExpr is exited.
func (s *BaseminiaListener) ExitDecimalListExpr(ctx *DecimalListExprContext) {}

// EnterFuncCall is called when production funcCall is entered.
func (s *BaseminiaListener) EnterFuncCall(ctx *FuncCallContext) {}

// ExitFuncCall is called when production funcCall is exited.
func (s *BaseminiaListener) ExitFuncCall(ctx *FuncCallContext) {}

// EnterExprList is called when production exprList is entered.
func (s *BaseminiaListener) EnterExprList(ctx *ExprListContext) {}

// ExitExprList is called when production exprList is exited.
func (s *BaseminiaListener) ExitExprList(ctx *ExprListContext) {}

// EnterStringExpr is called when production StringExpr is entered.
func (s *BaseminiaListener) EnterStringExpr(ctx *StringExprContext) {}

// ExitStringExpr is called when production StringExpr is exited.
func (s *BaseminiaListener) ExitStringExpr(ctx *StringExprContext) {}

// EnterIntegerExpr is called when production IntegerExpr is entered.
func (s *BaseminiaListener) EnterIntegerExpr(ctx *IntegerExprContext) {}

// ExitIntegerExpr is called when production IntegerExpr is exited.
func (s *BaseminiaListener) ExitIntegerExpr(ctx *IntegerExprContext) {}

// EnterDecimalExpr is called when production DecimalExpr is entered.
func (s *BaseminiaListener) EnterDecimalExpr(ctx *DecimalExprContext) {}

// ExitDecimalExpr is called when production DecimalExpr is exited.
func (s *BaseminiaListener) ExitDecimalExpr(ctx *DecimalExprContext) {}
