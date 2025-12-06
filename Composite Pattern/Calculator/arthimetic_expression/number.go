package arthimetic_expression

type Number struct {
	value float64
}

func NewNumber(value float64) *Number {
	return &Number{value: value}
}

func (n *Number) Evaluate() float64 {
	return n.value
}
