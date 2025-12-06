package strategy

import "fmt"

type CardPayment struct{}

func NewCardPayment() *CardPayment {
	return &CardPayment{}
}

func (cp *CardPayment) ProcessPayment(id int, amount float64) {
	fmt.Println("Processing Card payment of", amount, "for payment ID:", id)
}
