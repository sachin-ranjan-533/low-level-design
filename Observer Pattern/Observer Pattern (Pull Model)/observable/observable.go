package observable

import "observer-patterns/interfaces"

type Observable interface {
	Subscribe(observer interfaces.Observer)
	Unsubscribe(observer interfaces.Observer)
	Notify(data interface{})
}
