package templates

import "fmt"

type PayToMerchant struct {
	senderAccountNo   int
	receiverAccountNo int
	amount            float64
}

func NewPayToMerchant(senderAccountNo int, receiverAccountNo int, amount float64) *PayToMerchant {
	return &PayToMerchant{
		senderAccountNo:   senderAccountNo,
		receiverAccountNo: receiverAccountNo,
		amount:            amount,
	}
}

func (ptm *PayToMerchant) ReceiveRequest(senderAccountNo int, receiverAccountNo int) {
	fmt.Println("Received Request to transfer from user", senderAccountNo, "to merchant", receiverAccountNo)
}

func (ptm *PayToMerchant) DebitMoney(amount float64, senderAccountNo int) {
	fmt.Println("Rupees", amount, "debited from user account no", senderAccountNo)
}

func (ptm *PayToMerchant) CreditMoney(amount float64, receiverAccountNo int) {
	fmt.Println("Rupees", amount, "credtied to merchant account no", receiverAccountNo)
}

func (ptm *PayToMerchant) SendNotification(senderAccountNo int, receiverAccountNo int) {
	fmt.Println("Money transfer from user ", senderAccountNo, "to merchant", receiverAccountNo)
}
