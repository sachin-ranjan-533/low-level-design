package vehicle_inventory

import "car-rental/vehicle"

type VehicleInventory interface {
	AddVehicle(vehicle vehicle.Vehicle)
	GetVehicles() []vehicle.Vehicle
	GetVehicleType() string
}
