package caretaker

import (
	"memento-pattern/memento"
	"memento-pattern/originator"
)

type Caretaker interface {
	AddMemento(memento.Memento)
	Undo(originator.Originator)
}
