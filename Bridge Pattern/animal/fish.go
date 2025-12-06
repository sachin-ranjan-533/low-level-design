package animal

import "bridge-pattern/breathing"

type Fish struct {
	Breathing breathing.Breathing
}

func NewFish(b breathing.Breathing) *Fish {
	return &Fish{Breathing: b}
}

func (f *Fish) Breathe() {
	f.Breathing.Breathe()
}
