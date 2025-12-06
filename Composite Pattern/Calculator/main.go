package main

import (
	"calculator/arthimetic_expression"
	"fmt"
)

func main() {
	expression := arthimetic_expression.NewExpression(
		arthimetic_expression.NewExpression(
			arthimetic_expression.NewNumber(5),
			arthimetic_expression.NewNumber(3),
			"+",
		),
		arthimetic_expression.NewExpression(
			arthimetic_expression.NewNumber(10),
			arthimetic_expression.NewNumber(2),
			"*",
		),
		"-",
	)

	result := expression.Evaluate()
	fmt.Println("Result:", result)
}
