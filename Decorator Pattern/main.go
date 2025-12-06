package main

import (
	"decorator-pattern/decorator"
	"decorator-pattern/pizza"
)

func main() {
	basePizza := &pizza.VegDelight{}
	cheesePizza := &decorator.Cheese{BasePizza: basePizza}
	mushroomCheesePizza := &decorator.Mushroom{BasePizza: cheesePizza}
	paneerMushroomCheesePizza := &decorator.Paneer{BasePizza: mushroomCheesePizza}

	finalPrice := paneerMushroomCheesePizza.GetPrice()
	println("Final Pizza Price:", finalPrice)
}
