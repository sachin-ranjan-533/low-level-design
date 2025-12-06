package arthimetic_expression

type Expression struct {
	leftExpression  ArthimeticExpression
	rightExpression ArthimeticExpression
	operator        string
}

func NewExpression(leftExpression ArthimeticExpression, rightExpression ArthimeticExpression, operator string) *Expression {
	return &Expression{
		leftExpression:  leftExpression,
		rightExpression: rightExpression,
		operator:        operator,
	}
}

func (e *Expression) Evaluate() float64 {
	switch e.operator {
	case "+":
		return e.leftExpression.Evaluate() + e.rightExpression.Evaluate()
	case "-":
		return e.leftExpression.Evaluate() - e.rightExpression.Evaluate()
	case "*":
		return e.leftExpression.Evaluate() * e.rightExpression.Evaluate()
	case "/":
		return e.leftExpression.Evaluate() / e.rightExpression.Evaluate()
	default:
		return 0
	}
}
