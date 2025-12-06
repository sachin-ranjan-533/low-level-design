package main

import "null-pattern/vehicle"

type Factory struct{}

func (f *Factory) CreateVehicle(vehicleType string) vehicle.Vehicle {
	switch vehicleType {
	case "Bike":
		return &vehicle.Bike{}
	case "Car":
		return &vehicle.Car{}
	default:
		return &vehicle.DefaultVehicle{}
	}
}
