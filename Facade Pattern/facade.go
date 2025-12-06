package main

import orderpkg "facade-pattern/order"

type OrderCreation struct {
	ProductID     string
	ProductName   string
	ProductPrice  float64
	PaymentMethod string
	InvoiceNumber string
}

func CreateOrder(order OrderCreation) (string, string, string) {
	product := orderpkg.NewProduct(order.ProductID, order.ProductName, order.ProductPrice)
	payment := orderpkg.NewPayment(order.PaymentMethod, order.ProductPrice)
	invoice := orderpkg.NewInvoice(order.InvoiceNumber, order.ProductPrice)

	return product.GetDetails(), payment.GetPaymentInfo(), invoice.GetInfo()
}
