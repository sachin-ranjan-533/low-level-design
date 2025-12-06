package originator

import (
	"memento-pattern/memento"
)

type ConfigOriginator struct {
	X int
	Y int
}

func NewConfigOriginator() *ConfigOriginator {
	return &ConfigOriginator{X: 0, Y: 0}
}

func (co *ConfigOriginator) SetData(x int, y int) {
	co.X = x
	co.Y = y
}

func (co *ConfigOriginator) CreateMemento() *memento.ConfigMemento {
	return memento.NewConfigMemento(co.X, co.Y)
}

func (co *ConfigOriginator) RestoreMemento(memento *memento.ConfigMemento) {
	co.X = memento.X
	co.Y = memento.Y
}
