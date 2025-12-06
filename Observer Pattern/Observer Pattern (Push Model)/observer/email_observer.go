package observer

import (
	"fmt"
	observerable_data "observable-pattern/observable_data"
)

type EmailObserver struct {
	email string
}

func NewEmailObserver(email string) *EmailObserver {
	return &EmailObserver{
		email: email,
	}
}

func (eo *EmailObserver) Update(observableData *observerable_data.ObservableData) {
	fmt.Println("Email:", eo.email, "Data Updated:", observableData.Count)
}
