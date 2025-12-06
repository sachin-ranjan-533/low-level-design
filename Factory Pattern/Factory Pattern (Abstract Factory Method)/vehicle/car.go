package vehicle

import "fmt"

type Car struct {
	model string
}

func NewCar(model string) *Car {
	return &Car{model: model}
}

func (b *Car) Drive() {
	fmt.Println("Riding Car:", b.model)
}
