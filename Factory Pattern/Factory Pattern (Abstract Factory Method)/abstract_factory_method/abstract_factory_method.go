package abstract_factory_method

import (
	"abstract-factory/factory"
	"abstract-factory/vehicle"
)

func AbstractFactoryMethod(vehicleType string, model string) vehicle.Vehicle {
	if vehicleType == "car" {
		carFactory := factory.NewCarFactory()
		return carFactory.CreateVehicle(model)
	} else if vehicleType == "bike" {
		bikeFactory := factory.NewBikeFactory()
		return bikeFactory.CreateVehicle(model)
	}
	return nil
}
