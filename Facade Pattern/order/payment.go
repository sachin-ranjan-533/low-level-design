package order

import "fmt"

type Payment struct {
	Method string
	Amount float64
}

func NewPayment(method string, amount float64) *Payment {
	return &Payment{
		Method: method,
		Amount: amount,
	}
}

func (p *Payment) GetPaymentInfo() string {
	return "Payment Method: " + p.Method + ", Amount: $" + fmt.Sprintf("%.2f", p.Amount)
}
