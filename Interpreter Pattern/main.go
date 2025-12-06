package main

import (
	"fmt"
	"interpreter-pattern/context"
	"interpreter-pattern/expression"
)

func main() {
	var m map[string]float64 = map[string]float64{
		"a": 2.0,
		"b": 4.0,
		"c": 6.0,
		"d": 8.0,
	}
	context := context.NewContext(m)
	subExpression1 := expression.NewBinaryExpression(expression.NewNumberExpression("a"), expression.NewNumberExpression("b"), "*")
	subExpression2 := expression.NewBinaryExpression(expression.NewNumberExpression("c"), expression.NewNumberExpression("d"), "*")
	expression := expression.NewBinaryExpression(subExpression1, subExpression2, "+")
	result := expression.Interpret(context)
	fmt.Println("Result is", result)
}
