package expression

import "interpreter-pattern/context"

type NumberExpression struct {
	key string
}

func NewNumberExpression(key string) *NumberExpression {
	return &NumberExpression{key: key}
}

func (n *NumberExpression) Interpret(ctx *context.Context) float64 {
	v, _ := ctx.GetNumber(n.key)
	return v
}
