package factory

import "abstract-factory/vehicle"

type CarFactory struct{}

func NewCarFactory() *CarFactory {
	return &CarFactory{}
}

func (bf *CarFactory) CreateVehicle(model string) vehicle.Vehicle {
	return vehicle.NewCar(model)
}
