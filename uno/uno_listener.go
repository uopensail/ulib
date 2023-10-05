// Code generated from uno.g4 by ANTLR 4.13.1. DO NOT EDIT.

package uno // uno
import "github.com/antlr4-go/antlr/v4"

// unoListener is a complete listener for a parse tree produced by unoParser.
type unoListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterCmpBooleanExpression is called when entering the CmpBooleanExpression production.
	EnterCmpBooleanExpression(c *CmpBooleanExpressionContext)

	// EnterNotBooleanExpression is called when entering the NotBooleanExpression production.
	EnterNotBooleanExpression(c *NotBooleanExpressionContext)

	// EnterPlainBooleanExpression is called when entering the PlainBooleanExpression production.
	EnterPlainBooleanExpression(c *PlainBooleanExpressionContext)

	// EnterOrBooleanExpression is called when entering the OrBooleanExpression production.
	EnterOrBooleanExpression(c *OrBooleanExpressionContext)

	// EnterTrueBooleanExpression is called when entering the TrueBooleanExpression production.
	EnterTrueBooleanExpression(c *TrueBooleanExpressionContext)

	// EnterAndBooleanExpression is called when entering the AndBooleanExpression production.
	EnterAndBooleanExpression(c *AndBooleanExpressionContext)

	// EnterNotInBooleanExpression is called when entering the NotInBooleanExpression production.
	EnterNotInBooleanExpression(c *NotInBooleanExpressionContext)

	// EnterFalseBooleanExpression is called when entering the FalseBooleanExpression production.
	EnterFalseBooleanExpression(c *FalseBooleanExpressionContext)

	// EnterInBooleanExpression is called when entering the InBooleanExpression production.
	EnterInBooleanExpression(c *InBooleanExpressionContext)

	// EnterPlainArithmeticExpression is called when entering the PlainArithmeticExpression production.
	EnterPlainArithmeticExpression(c *PlainArithmeticExpressionContext)

	// EnterAddArithmeticExpression is called when entering the AddArithmeticExpression production.
	EnterAddArithmeticExpression(c *AddArithmeticExpressionContext)

	// EnterStringArithmeticExpression is called when entering the StringArithmeticExpression production.
	EnterStringArithmeticExpression(c *StringArithmeticExpressionContext)

	// EnterIntegerArithmeticExpression is called when entering the IntegerArithmeticExpression production.
	EnterIntegerArithmeticExpression(c *IntegerArithmeticExpressionContext)

	// EnterDecimalArithmeticExpression is called when entering the DecimalArithmeticExpression production.
	EnterDecimalArithmeticExpression(c *DecimalArithmeticExpressionContext)

	// EnterFuncArithmeticExpression is called when entering the FuncArithmeticExpression production.
	EnterFuncArithmeticExpression(c *FuncArithmeticExpressionContext)

	// EnterColumnArithmeticExpression is called when entering the ColumnArithmeticExpression production.
	EnterColumnArithmeticExpression(c *ColumnArithmeticExpressionContext)

	// EnterDivArithmeticExpression is called when entering the DivArithmeticExpression production.
	EnterDivArithmeticExpression(c *DivArithmeticExpressionContext)

	// EnterFieldColumnArithmeticExpression is called when entering the FieldColumnArithmeticExpression production.
	EnterFieldColumnArithmeticExpression(c *FieldColumnArithmeticExpressionContext)

	// EnterSubArithmeticExpression is called when entering the SubArithmeticExpression production.
	EnterSubArithmeticExpression(c *SubArithmeticExpressionContext)

	// EnterModArithmeticExpression is called when entering the ModArithmeticExpression production.
	EnterModArithmeticExpression(c *ModArithmeticExpressionContext)

	// EnterRuntTimeFuncArithmeticExpression is called when entering the RuntTimeFuncArithmeticExpression production.
	EnterRuntTimeFuncArithmeticExpression(c *RuntTimeFuncArithmeticExpressionContext)

	// EnterMulArithmeticExpression is called when entering the MulArithmeticExpression production.
	EnterMulArithmeticExpression(c *MulArithmeticExpressionContext)

	// EnterType_marker is called when entering the type_marker production.
	EnterType_marker(c *Type_markerContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitCmpBooleanExpression is called when exiting the CmpBooleanExpression production.
	ExitCmpBooleanExpression(c *CmpBooleanExpressionContext)

	// ExitNotBooleanExpression is called when exiting the NotBooleanExpression production.
	ExitNotBooleanExpression(c *NotBooleanExpressionContext)

	// ExitPlainBooleanExpression is called when exiting the PlainBooleanExpression production.
	ExitPlainBooleanExpression(c *PlainBooleanExpressionContext)

	// ExitOrBooleanExpression is called when exiting the OrBooleanExpression production.
	ExitOrBooleanExpression(c *OrBooleanExpressionContext)

	// ExitTrueBooleanExpression is called when exiting the TrueBooleanExpression production.
	ExitTrueBooleanExpression(c *TrueBooleanExpressionContext)

	// ExitAndBooleanExpression is called when exiting the AndBooleanExpression production.
	ExitAndBooleanExpression(c *AndBooleanExpressionContext)

	// ExitNotInBooleanExpression is called when exiting the NotInBooleanExpression production.
	ExitNotInBooleanExpression(c *NotInBooleanExpressionContext)

	// ExitFalseBooleanExpression is called when exiting the FalseBooleanExpression production.
	ExitFalseBooleanExpression(c *FalseBooleanExpressionContext)

	// ExitInBooleanExpression is called when exiting the InBooleanExpression production.
	ExitInBooleanExpression(c *InBooleanExpressionContext)

	// ExitPlainArithmeticExpression is called when exiting the PlainArithmeticExpression production.
	ExitPlainArithmeticExpression(c *PlainArithmeticExpressionContext)

	// ExitAddArithmeticExpression is called when exiting the AddArithmeticExpression production.
	ExitAddArithmeticExpression(c *AddArithmeticExpressionContext)

	// ExitStringArithmeticExpression is called when exiting the StringArithmeticExpression production.
	ExitStringArithmeticExpression(c *StringArithmeticExpressionContext)

	// ExitIntegerArithmeticExpression is called when exiting the IntegerArithmeticExpression production.
	ExitIntegerArithmeticExpression(c *IntegerArithmeticExpressionContext)

	// ExitDecimalArithmeticExpression is called when exiting the DecimalArithmeticExpression production.
	ExitDecimalArithmeticExpression(c *DecimalArithmeticExpressionContext)

	// ExitFuncArithmeticExpression is called when exiting the FuncArithmeticExpression production.
	ExitFuncArithmeticExpression(c *FuncArithmeticExpressionContext)

	// ExitColumnArithmeticExpression is called when exiting the ColumnArithmeticExpression production.
	ExitColumnArithmeticExpression(c *ColumnArithmeticExpressionContext)

	// ExitDivArithmeticExpression is called when exiting the DivArithmeticExpression production.
	ExitDivArithmeticExpression(c *DivArithmeticExpressionContext)

	// ExitFieldColumnArithmeticExpression is called when exiting the FieldColumnArithmeticExpression production.
	ExitFieldColumnArithmeticExpression(c *FieldColumnArithmeticExpressionContext)

	// ExitSubArithmeticExpression is called when exiting the SubArithmeticExpression production.
	ExitSubArithmeticExpression(c *SubArithmeticExpressionContext)

	// ExitModArithmeticExpression is called when exiting the ModArithmeticExpression production.
	ExitModArithmeticExpression(c *ModArithmeticExpressionContext)

	// ExitRuntTimeFuncArithmeticExpression is called when exiting the RuntTimeFuncArithmeticExpression production.
	ExitRuntTimeFuncArithmeticExpression(c *RuntTimeFuncArithmeticExpressionContext)

	// ExitMulArithmeticExpression is called when exiting the MulArithmeticExpression production.
	ExitMulArithmeticExpression(c *MulArithmeticExpressionContext)

	// ExitType_marker is called when exiting the type_marker production.
	ExitType_marker(c *Type_markerContext)
}
