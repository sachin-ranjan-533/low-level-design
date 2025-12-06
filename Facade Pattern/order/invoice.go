package order

import "fmt"

type Invoice struct {
	Number string
	Amount float64
}

func NewInvoice(number string, amount float64) *Invoice {
	return &Invoice{
		Number: number,
		Amount: amount,
	}
}

func (i *Invoice) GetInfo() string {
	return "Invoice Number: " + i.Number + ", Amount: $" + fmt.Sprintf("%.2f", i.Amount)
}
