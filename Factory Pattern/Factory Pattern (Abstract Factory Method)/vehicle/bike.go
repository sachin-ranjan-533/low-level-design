package vehicle

import "fmt"

type Bike struct {
	model string
}

func NewBike(model string) *Bike {
	return &Bike{model: model}
}

func (b *Bike) Drive() {
	fmt.Println("Riding Bike:", b.model)
}
