package animal

import "bridge-pattern/breathing"

type Cat struct {
	Breathing breathing.Breathing
}

func NewCat(b breathing.Breathing) *Cat {
	return &Cat{Breathing: b}
}

func (c *Cat) Breathe() {
	c.Breathing.Breathe()
}
