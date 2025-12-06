package carinterior

import "fmt"

type NormalCarInterior struct{}

func NewNormalCarInterior() *NormalCarInterior {
	return &NormalCarInterior{}
}

func (lce *NormalCarInterior) GetInteriorDetails() {
	fmt.Println("Normal Car Interior: Standard design with basic features.")
}
