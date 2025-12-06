package factory

import "abstract-factory/vehicle"

type BikeFactory struct{}

func NewBikeFactory() *BikeFactory {
	return &BikeFactory{}
}

func (bf *BikeFactory) CreateVehicle(model string) vehicle.Vehicle {
	return vehicle.NewBike(model)
}
