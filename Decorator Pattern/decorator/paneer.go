package decorator

import "decorator-pattern/pizza"

type Paneer struct {
	BasePizza pizza.Base
}

func (p *Paneer) GetPrice() float64 {
	return p.BasePizza.GetPrice() + 4.0
}
