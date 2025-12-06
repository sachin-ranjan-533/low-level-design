package main

import (
	"bridge-pattern/animal"
	"bridge-pattern/breathing"
)

func main() {
	landBreathing := &breathing.LandBreathing{}
	fish := animal.NewFish(landBreathing)
	fish.Breathe()

	waterBreathing := &breathing.WaterBreathing{}
	fishWater := animal.NewFish(waterBreathing)
	fishWater.Breathe()

	dog := animal.NewDog(landBreathing)
	dog.Breathe()
}
