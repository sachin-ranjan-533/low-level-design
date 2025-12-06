package main

import (
	"observer-patterns/observable"
	"observer-patterns/observer"
)

func main() {
	iphoneObservable := observable.NewIphoneObservable()
	emailObserver := observer.NewEmailObserver("sachinranjan533@gmail.com")
	phoneObserver := observer.NewPhoneObserver("+916205757958")

	iphoneObservable.Subscribe(emailObserver)
	iphoneObservable.Subscribe(phoneObserver)

	iphoneObservable.Notify("New iPhone 15 Released!")

	iphoneObservable.Unsubscribe(phoneObserver)

	iphoneObservable.Notify("iPhone 15 Pro Launched!")
}
