package expression

import "interpreter-pattern/context"

type Expression interface {
	Interpret(ctx *context.Context) float64
}
