package carexterior

import "fmt"

type LuxuryCarExterior struct{}

func NewLuxuryCarExterior() *LuxuryCarExterior {
	return &LuxuryCarExterior{}
}

func (lce *LuxuryCarExterior) GetExteriorDetails() {
	fmt.Println("Luxury Car Exterior: Sleek design with premium materials and advanced lighting.")
}
