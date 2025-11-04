// Code generated from minia.g4 by ANTLR 4.13.2. DO NOT EDIT.

package minia // minia
import "github.com/antlr4-go/antlr/v4"

// miniaListener is a complete listener for a parse tree produced by miniaParser.
type miniaListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterOrExpr is called when entering the OrExpr production.
	EnterOrExpr(c *OrExprContext)

	// EnterTrivialLogicalAndExpr is called when entering the TrivialLogicalAndExpr production.
	EnterTrivialLogicalAndExpr(c *TrivialLogicalAndExprContext)

	// EnterAndExpr is called when entering the AndExpr production.
	EnterAndExpr(c *AndExprContext)

	// EnterTrivialEqualityExpr is called when entering the TrivialEqualityExpr production.
	EnterTrivialEqualityExpr(c *TrivialEqualityExprContext)

	// EnterEqualExpr is called when entering the EqualExpr production.
	EnterEqualExpr(c *EqualExprContext)

	// EnterNotEqualExpr is called when entering the NotEqualExpr production.
	EnterNotEqualExpr(c *NotEqualExprContext)

	// EnterTrivialRelationalExpr is called when entering the TrivialRelationalExpr production.
	EnterTrivialRelationalExpr(c *TrivialRelationalExprContext)

	// EnterGreaterThanExpr is called when entering the GreaterThanExpr production.
	EnterGreaterThanExpr(c *GreaterThanExprContext)

	// EnterGreaterThanEqualExpr is called when entering the GreaterThanEqualExpr production.
	EnterGreaterThanEqualExpr(c *GreaterThanEqualExprContext)

	// EnterLessThanExpr is called when entering the LessThanExpr production.
	EnterLessThanExpr(c *LessThanExprContext)

	// EnterLessThanEqualExpr is called when entering the LessThanEqualExpr production.
	EnterLessThanEqualExpr(c *LessThanEqualExprContext)

	// EnterTrivialAdditiveExpr is called when entering the TrivialAdditiveExpr production.
	EnterTrivialAdditiveExpr(c *TrivialAdditiveExprContext)

	// EnterAddExpr is called when entering the AddExpr production.
	EnterAddExpr(c *AddExprContext)

	// EnterSubExpr is called when entering the SubExpr production.
	EnterSubExpr(c *SubExprContext)

	// EnterTrivialMultiplicativeExpr is called when entering the TrivialMultiplicativeExpr production.
	EnterTrivialMultiplicativeExpr(c *TrivialMultiplicativeExprContext)

	// EnterMulExpr is called when entering the MulExpr production.
	EnterMulExpr(c *MulExprContext)

	// EnterDivExpr is called when entering the DivExpr production.
	EnterDivExpr(c *DivExprContext)

	// EnterModExpr is called when entering the ModExpr production.
	EnterModExpr(c *ModExprContext)

	// EnterTrivialUnaryExpr is called when entering the TrivialUnaryExpr production.
	EnterTrivialUnaryExpr(c *TrivialUnaryExprContext)

	// EnterNotExpr is called when entering the NotExpr production.
	EnterNotExpr(c *NotExprContext)

	// EnterNegExpr is called when entering the NegExpr production.
	EnterNegExpr(c *NegExprContext)

	// EnterTrivialPrimaryExpr is called when entering the TrivialPrimaryExpr production.
	EnterTrivialPrimaryExpr(c *TrivialPrimaryExprContext)

	// EnterParenthesizedExpr is called when entering the ParenthesizedExpr production.
	EnterParenthesizedExpr(c *ParenthesizedExprContext)

	// EnterFunctionCallExpr is called when entering the FunctionCallExpr production.
	EnterFunctionCallExpr(c *FunctionCallExprContext)

	// EnterColumnExpr is called when entering the ColumnExpr production.
	EnterColumnExpr(c *ColumnExprContext)

	// EnterLiteralExpr is called when entering the LiteralExpr production.
	EnterLiteralExpr(c *LiteralExprContext)

	// EnterListExpr is called when entering the ListExpr production.
	EnterListExpr(c *ListExprContext)

	// EnterTrueExpr is called when entering the TrueExpr production.
	EnterTrueExpr(c *TrueExprContext)

	// EnterFalseExpr is called when entering the FalseExpr production.
	EnterFalseExpr(c *FalseExprContext)

	// EnterStringListExpr is called when entering the StringListExpr production.
	EnterStringListExpr(c *StringListExprContext)

	// EnterIntegerListExpr is called when entering the IntegerListExpr production.
	EnterIntegerListExpr(c *IntegerListExprContext)

	// EnterDecimalListExpr is called when entering the DecimalListExpr production.
	EnterDecimalListExpr(c *DecimalListExprContext)

	// EnterFuncCall is called when entering the funcCall production.
	EnterFuncCall(c *FuncCallContext)

	// EnterExprList is called when entering the exprList production.
	EnterExprList(c *ExprListContext)

	// EnterStringExpr is called when entering the StringExpr production.
	EnterStringExpr(c *StringExprContext)

	// EnterIntegerExpr is called when entering the IntegerExpr production.
	EnterIntegerExpr(c *IntegerExprContext)

	// EnterDecimalExpr is called when entering the DecimalExpr production.
	EnterDecimalExpr(c *DecimalExprContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitOrExpr is called when exiting the OrExpr production.
	ExitOrExpr(c *OrExprContext)

	// ExitTrivialLogicalAndExpr is called when exiting the TrivialLogicalAndExpr production.
	ExitTrivialLogicalAndExpr(c *TrivialLogicalAndExprContext)

	// ExitAndExpr is called when exiting the AndExpr production.
	ExitAndExpr(c *AndExprContext)

	// ExitTrivialEqualityExpr is called when exiting the TrivialEqualityExpr production.
	ExitTrivialEqualityExpr(c *TrivialEqualityExprContext)

	// ExitEqualExpr is called when exiting the EqualExpr production.
	ExitEqualExpr(c *EqualExprContext)

	// ExitNotEqualExpr is called when exiting the NotEqualExpr production.
	ExitNotEqualExpr(c *NotEqualExprContext)

	// ExitTrivialRelationalExpr is called when exiting the TrivialRelationalExpr production.
	ExitTrivialRelationalExpr(c *TrivialRelationalExprContext)

	// ExitGreaterThanExpr is called when exiting the GreaterThanExpr production.
	ExitGreaterThanExpr(c *GreaterThanExprContext)

	// ExitGreaterThanEqualExpr is called when exiting the GreaterThanEqualExpr production.
	ExitGreaterThanEqualExpr(c *GreaterThanEqualExprContext)

	// ExitLessThanExpr is called when exiting the LessThanExpr production.
	ExitLessThanExpr(c *LessThanExprContext)

	// ExitLessThanEqualExpr is called when exiting the LessThanEqualExpr production.
	ExitLessThanEqualExpr(c *LessThanEqualExprContext)

	// ExitTrivialAdditiveExpr is called when exiting the TrivialAdditiveExpr production.
	ExitTrivialAdditiveExpr(c *TrivialAdditiveExprContext)

	// ExitAddExpr is called when exiting the AddExpr production.
	ExitAddExpr(c *AddExprContext)

	// ExitSubExpr is called when exiting the SubExpr production.
	ExitSubExpr(c *SubExprContext)

	// ExitTrivialMultiplicativeExpr is called when exiting the TrivialMultiplicativeExpr production.
	ExitTrivialMultiplicativeExpr(c *TrivialMultiplicativeExprContext)

	// ExitMulExpr is called when exiting the MulExpr production.
	ExitMulExpr(c *MulExprContext)

	// ExitDivExpr is called when exiting the DivExpr production.
	ExitDivExpr(c *DivExprContext)

	// ExitModExpr is called when exiting the ModExpr production.
	ExitModExpr(c *ModExprContext)

	// ExitTrivialUnaryExpr is called when exiting the TrivialUnaryExpr production.
	ExitTrivialUnaryExpr(c *TrivialUnaryExprContext)

	// ExitNotExpr is called when exiting the NotExpr production.
	ExitNotExpr(c *NotExprContext)

	// ExitNegExpr is called when exiting the NegExpr production.
	ExitNegExpr(c *NegExprContext)

	// ExitTrivialPrimaryExpr is called when exiting the TrivialPrimaryExpr production.
	ExitTrivialPrimaryExpr(c *TrivialPrimaryExprContext)

	// ExitParenthesizedExpr is called when exiting the ParenthesizedExpr production.
	ExitParenthesizedExpr(c *ParenthesizedExprContext)

	// ExitFunctionCallExpr is called when exiting the FunctionCallExpr production.
	ExitFunctionCallExpr(c *FunctionCallExprContext)

	// ExitColumnExpr is called when exiting the ColumnExpr production.
	ExitColumnExpr(c *ColumnExprContext)

	// ExitLiteralExpr is called when exiting the LiteralExpr production.
	ExitLiteralExpr(c *LiteralExprContext)

	// ExitListExpr is called when exiting the ListExpr production.
	ExitListExpr(c *ListExprContext)

	// ExitTrueExpr is called when exiting the TrueExpr production.
	ExitTrueExpr(c *TrueExprContext)

	// ExitFalseExpr is called when exiting the FalseExpr production.
	ExitFalseExpr(c *FalseExprContext)

	// ExitStringListExpr is called when exiting the StringListExpr production.
	ExitStringListExpr(c *StringListExprContext)

	// ExitIntegerListExpr is called when exiting the IntegerListExpr production.
	ExitIntegerListExpr(c *IntegerListExprContext)

	// ExitDecimalListExpr is called when exiting the DecimalListExpr production.
	ExitDecimalListExpr(c *DecimalListExprContext)

	// ExitFuncCall is called when exiting the funcCall production.
	ExitFuncCall(c *FuncCallContext)

	// ExitExprList is called when exiting the exprList production.
	ExitExprList(c *ExprListContext)

	// ExitStringExpr is called when exiting the StringExpr production.
	ExitStringExpr(c *StringExprContext)

	// ExitIntegerExpr is called when exiting the IntegerExpr production.
	ExitIntegerExpr(c *IntegerExprContext)

	// ExitDecimalExpr is called when exiting the DecimalExpr production.
	ExitDecimalExpr(c *DecimalExprContext)
}
