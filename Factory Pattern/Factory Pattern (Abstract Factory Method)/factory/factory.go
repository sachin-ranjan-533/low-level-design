package factory

import "abstract-factory/vehicle"

type Factory interface {
	CreateVehicle(model string) vehicle.Vehicle
}
