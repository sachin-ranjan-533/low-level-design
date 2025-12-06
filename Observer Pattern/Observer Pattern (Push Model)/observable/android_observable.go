package observable

import (
	observerable_data "observable-pattern/observable_data"
	"observable-pattern/observer"
)

type AndroidObservable struct {
	count          int
	observers      []observer.Observer
	observableData *observerable_data.ObservableData
}

func NewAndroidObservable() *AndroidObservable {
	return &AndroidObservable{
		observableData: observerable_data.NewObservableData(0),
	}
}

func (ao *AndroidObservable) Add(observer observer.Observer) {
	ao.observers = append(ao.observers, observer)
}

func (ao *AndroidObservable) Remove(observer observer.Observer) {
	for i, obs := range ao.observers {
		if obs == observer {
			ao.observers = append(ao.observers[:i], ao.observers[i+1:]...)
		}
	}
}

func (ao *AndroidObservable) Notify() {
	for _, observer := range ao.observers {
		observer.Update(ao.observableData)
	}
}

func (ao *AndroidObservable) SetData(count int) {
	currentCount := ao.observableData.Count
	ao.observableData.SetData(count)
	if currentCount == 0 {
		ao.Notify()
	}
}
