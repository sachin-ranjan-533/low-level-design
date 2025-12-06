package cartype

import (
	carexterior "abstract-factory/car_exterior"
	carinterior "abstract-factory/car_interior"
)

type NormalCarType struct{}

func NewNormalCarType() *NormalCarType {
	return &NormalCarType{}
}

func (nct *NormalCarType) GetExterior() carexterior.CarExterior {
	return carexterior.NewNormalCarExterior()
}

func (nct *NormalCarType) GetInterior() carinterior.CarInterior {
	return carinterior.NewNormalCarInterior()
}
