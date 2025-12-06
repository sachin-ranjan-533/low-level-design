package observable

import "observer-patterns/interfaces"

type IphoneObservable struct {
	observers []interfaces.Observer
}

func NewIphoneObservable() *IphoneObservable {
	return &IphoneObservable{
		observers: make([]interfaces.Observer, 0),
	}
}

func (io *IphoneObservable) Subscribe(observer interfaces.Observer) {
	io.observers = append(io.observers, observer)
}

func (io *IphoneObservable) Unsubscribe(observer interfaces.Observer) {
	for i, obs := range io.observers {
		if obs == observer {
			io.observers = append(io.observers[:i], io.observers[i+1:]...)
			break
		}
	}
}

func (io *IphoneObservable) Notify(data interface{}) {
	for _, observer := range io.observers {
		observer.Update(data.(string))
	}
}
