package carexterior

import "fmt"

type NormalCarExterior struct{}

func NewNormalCarExterior() *NormalCarExterior {
	return &NormalCarExterior{}
}

func (lce *NormalCarExterior) GetExteriorDetails() {
	fmt.Println("Normal Car Exterior: Standard design with basic features.")
}
