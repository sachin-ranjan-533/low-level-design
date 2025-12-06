package templates

import "fmt"

type PayToFriend struct {
	senderAccountNo   int
	receiverAccountNo int
	amount            float64
}

func NewPayToFriend(senderAccountNo int, receiverAccountNo int, amount float64) *PayToFriend {
	return &PayToFriend{
		senderAccountNo:   senderAccountNo,
		receiverAccountNo: receiverAccountNo,
		amount:            amount,
	}
}

func (ptf *PayToFriend) ReceiveRequest(senderAccountNo int, receiverAccountNo int) {
	fmt.Println("Received Request to transfer from user", senderAccountNo, "to user", receiverAccountNo)
}

func (ptf *PayToFriend) DebitMoney(amount float64, senderAccountNo int) {
	fmt.Println("Rupees", amount, "debited from user account no", senderAccountNo)
}

func (ptf *PayToFriend) CreditMoney(amount float64, receiverAccountNo int) {
	fmt.Println("Rupees", amount, "credtied to user account no", receiverAccountNo)

}

func (ptf *PayToFriend) SendNotification(senderAccountNo int, receiverAccountNo int) {
	fmt.Println("Money transfer from user", senderAccountNo, "to user", receiverAccountNo)
}
