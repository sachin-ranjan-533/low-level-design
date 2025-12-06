package memento

type ConfigMemento struct {
	X int
	Y int
}

func NewConfigMemento(x int, y int) *ConfigMemento {
	return &ConfigMemento{X: x, Y: y}
}
