package observer

import "observer-patterns/observable"

type PhoneObserver struct {
	Email      string
	observable observable.Observable
}

func NewPhoneObserver(email string) *PhoneObserver {
	return &PhoneObserver{
		Email: email,
	}
}

func (po *PhoneObserver) Update(data string) {
	println("Sending sms to", po.Email, "with message:", data)
}
