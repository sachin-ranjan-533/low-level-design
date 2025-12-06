package cartype

import (
	carexterior "abstract-factory/car_exterior"
	carinterior "abstract-factory/car_interior"
)

type LuxuryCarType struct{}

func NewLuxuryCarType() *LuxuryCarType {
	return &LuxuryCarType{}
}

func (lct *LuxuryCarType) GetExterior() carexterior.CarExterior {
	return carexterior.NewLuxuryCarExterior()
}

func (lct *LuxuryCarType) GetInterior() carinterior.CarInterior {
	return carinterior.NewLuxuryCarInterior()
}
