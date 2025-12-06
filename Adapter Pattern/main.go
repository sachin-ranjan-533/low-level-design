package main

import (
	"adapter-pattern/adaptee"
	"adapter-pattern/adapter"
	"fmt"
)

func main() {
	babyMachine := adaptee.NewWeightMachineBaby(50) // 50 pounds
	adapterObj := adapter.NewAdapter(babyMachine)

	fmt.Printf("Weight in KG: %.2f\n", adapterObj.GetWeightInKg())
}
