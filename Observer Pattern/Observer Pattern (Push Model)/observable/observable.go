package observable

import "observable-pattern/observer"

type Observable interface {
	Add(observer.Observer)
	Remove(observer.Observer)
	Notify()
	SetData()
}
