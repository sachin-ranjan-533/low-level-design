package strategy

import "fmt"

type UpiPayment struct{}

func NewUPIPaymentStrategy() *UpiPayment {
	return &UpiPayment{}
}

func (up *UpiPayment) ProcessPayment(id int, amount float64) {
	fmt.Println("Processing UPI payment of", amount, "for payment ID:", id)
}
