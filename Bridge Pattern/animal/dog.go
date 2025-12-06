package animal

import "bridge-pattern/breathing"

type Dog struct {
	Breathing breathing.Breathing
}

func NewDog(b breathing.Breathing) *Dog {
	return &Dog{Breathing: b}
}

func (d *Dog) Breathe() {
	d.Breathing.Breathe()
}
