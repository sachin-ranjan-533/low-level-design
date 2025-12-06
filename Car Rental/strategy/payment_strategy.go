package strategy

type PaymentStrategy interface {
	ProcessPayment(id int, amount float64)
}
