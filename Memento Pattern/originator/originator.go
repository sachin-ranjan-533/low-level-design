package originator

import "memento-pattern/memento"

type Originator interface {
	SetData(x int, y int)
	AddMemento() *memento.Memento
	RestoreMemento()
}
