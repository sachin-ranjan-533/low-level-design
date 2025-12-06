package observable

import (
	observerable_data "observable-pattern/observable_data"
	"observable-pattern/observer"
)

type IphoneObservable struct {
	count          int
	observers      []observer.Observer
	observableData *observerable_data.ObservableData
}

func NewIphoneObservable() *IphoneObservable {
	return &IphoneObservable{
		observableData: observerable_data.NewObservableData(0),
	}
}

func (io *IphoneObservable) Add(observer observer.Observer) {
	io.observers = append(io.observers, observer)
}

func (io *IphoneObservable) Remove(observer observer.Observer) {
	for i, obs := range io.observers {
		if obs == observer {
			io.observers = append(io.observers[:i], io.observers[i+1:]...)
		}
	}
}

func (io *IphoneObservable) Notify() {
	for _, observer := range io.observers {
		observer.Update(io.observableData)
	}
}

func (io *IphoneObservable) SetData(count int) {
	currentCount := io.observableData.Count
	io.observableData.SetData(count)
	if currentCount == 0 {
		io.Notify()
	}
}
