package observer

import "observer-patterns/observable"

type EmailObserver struct {
	Email      string
	observable observable.Observable
}

func NewEmailObserver(email string) *EmailObserver {
	return &EmailObserver{
		Email: email,
	}
}

func (po *EmailObserver) Update(data string) {
	println("Sending email to", po.Email, "with message:", data)
}
