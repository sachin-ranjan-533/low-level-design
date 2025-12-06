package expression

import "interpreter-pattern/context"

type BinaryExpression struct {
	leftExpression  Expression
	rightExpression Expression
	operator        string
}

func NewBinaryExpression(left Expression, right Expression, operator string) *BinaryExpression {
	return &BinaryExpression{
		leftExpression:  left,
		rightExpression: right,
		operator:        operator,
	}
}

func (be *BinaryExpression) Interpret(ctx *context.Context) float64 {
	switch be.operator {
	case "+":
		return be.leftExpression.Interpret(ctx) + be.rightExpression.Interpret(ctx)
	case "-":
		return be.leftExpression.Interpret(ctx) - be.rightExpression.Interpret(ctx)
	case "*":
		return be.leftExpression.Interpret(ctx) * be.rightExpression.Interpret(ctx)
	case "/":
		return be.leftExpression.Interpret(ctx) / be.rightExpression.Interpret(ctx)
	}
	return 0
}
