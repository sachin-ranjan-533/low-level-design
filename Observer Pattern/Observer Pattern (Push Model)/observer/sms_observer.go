package observer

import (
	"fmt"
	observerable_data "observable-pattern/observable_data"
)

type SmsObserver struct {
	number string
}

func NewSmsObserver(number string) *SmsObserver {
	return &SmsObserver{
		number: number,
	}
}

func (s *SmsObserver) Update(observableData *observerable_data.ObservableData) {
	fmt.Println("Sms:", s.number, "Data Updated:", observableData.Count)
}
