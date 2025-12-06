package decorator

import "decorator-pattern/pizza"

type Cheese struct {
	BasePizza pizza.Base
}

func (c *Cheese) GetPrice() float64 {
	return c.BasePizza.GetPrice() + 2.0
}
