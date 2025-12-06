package main

func main() {
	order := OrderCreation{
		ProductID:     "P123",
		ProductName:   "Laptop",
		ProductPrice:  999.99,
		PaymentMethod: "Credit Card",
		InvoiceNumber: "INV456",
	}

	productDetails, paymentInfo, invoiceInfo := CreateOrder(order)

	println(productDetails)
	println(paymentInfo)
	println(invoiceInfo)
}
