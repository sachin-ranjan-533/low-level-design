package main

import (
	"fmt"
	"memento-pattern/caretaker"
	"memento-pattern/originator"
)

func main() {
	originator := originator.NewConfigOriginator()
	caretaker := caretaker.NewConfigCaretaker()
	originator.SetData(10, 20)
	memento1 := originator.CreateMemento()
	caretaker.AddMemento(*memento1)

	originator.SetData(20, 30)
	memento2 := originator.CreateMemento()
	caretaker.AddMemento(*memento2)

	originator.SetData(30, 40)
	memento3 := originator.CreateMemento()
	caretaker.AddMemento(*memento3)

	caretaker.Undo(originator)
	caretaker.Undo(originator)

	fmt.Printf("X value is %d and Y value is %d\n", originator.X, originator.Y)
}
