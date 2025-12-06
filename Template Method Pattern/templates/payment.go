package templates

type Payment interface {
	ReceiveRequest(senderAccountNo int, receiverAccountNo int)
	DebitMoney(amount float64, senderAccountNo int)
	CreditMoney(amount float64, receiverAccountNo int)
	SendNotification(senderAccountNo int, receiverAccountNo int)
}

func SendMoney(p Payment, senderAccountNo int, receiverAccountNo int, amount float64) {
	p.ReceiveRequest(senderAccountNo, receiverAccountNo)
	p.DebitMoney(amount, senderAccountNo)
	p.CreditMoney(amount, receiverAccountNo)
	p.SendNotification(senderAccountNo, receiverAccountNo)
}
