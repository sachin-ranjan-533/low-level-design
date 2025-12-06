package carinterior

import "fmt"

type LuxuryCarInterior struct{}

func NewLuxuryCarInterior() *LuxuryCarInterior {
	return &LuxuryCarInterior{}
}

func (lce *LuxuryCarInterior) GetInteriorDetails() {
	fmt.Println("Luxury Car Interior: Sleek design with premium materials and advanced lighting.")
}
