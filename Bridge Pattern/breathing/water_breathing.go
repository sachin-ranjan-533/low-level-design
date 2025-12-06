package breathing

import "fmt"

type WaterBreathing struct{}

func (wb *WaterBreathing) Breathe() {
	fmt.Println("Breathing on water")
}
