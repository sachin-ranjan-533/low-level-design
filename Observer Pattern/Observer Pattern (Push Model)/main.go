package main

import (
	"observable-pattern/observable"
	"observable-pattern/observer"
)

func main() {
	io := observable.NewIphoneObservable()
	io.Add(observer.NewEmailObserver("sachin.ranjan@unthinkable.co"))
	io.Add(observer.NewEmailObserver("prince.breja@unthinkable.co"))
	io.Add(observer.NewSmsObserver("1234567890"))
	io.Add(observer.NewSmsObserver("0987654321"))
	io.SetData(10)
}
