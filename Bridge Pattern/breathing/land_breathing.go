package breathing

import "fmt"

type LandBreathing struct{}

func (lb *LandBreathing) Breathe() {
	fmt.Println("Breathing on land")
}
