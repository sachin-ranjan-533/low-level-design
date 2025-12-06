package decorator

import "decorator-pattern/pizza"

type Mushroom struct {
	BasePizza pizza.Base
}

func (m *Mushroom) GetPrice() float64 {
	return m.BasePizza.GetPrice() + 3.0
}
