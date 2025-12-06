package payment

import (
	"car-rental/strategy"
)

type Payment struct {
	id       int
	amount   float64
	strategy strategy.PaymentStrategy
}

func NewPayment(id int, amount float64, strategy strategy.PaymentStrategy) *Payment {
	return &Payment{
		id:       id,
		amount:   amount,
		strategy: strategy,
	}
}

func (p *Payment) ProcessPayment() {
	p.strategy.ProcessPayment(p.id, p.amount)
}
