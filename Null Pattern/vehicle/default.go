package vehicle

import "fmt"

type DefaultVehicle struct{}

func (d *DefaultVehicle) Drive() {
	fmt.Println("No vehicle selected")
}
