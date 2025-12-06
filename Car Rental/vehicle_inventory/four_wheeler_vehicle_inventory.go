package vehicle_inventory

import (
	"car-rental/vehicle"
)

type FourWheelerVehicleInventory struct {
	vehicles []vehicle.Vehicle
}

func NewFourWheelerVehicleInventory() *FourWheelerVehicleInventory {
	return &FourWheelerVehicleInventory{}
}

func (fwvi *FourWheelerVehicleInventory) AddVehicle(vehicle vehicle.Vehicle) {
	fwvi.vehicles = append(fwvi.vehicles, vehicle)
}

func (fwvi *FourWheelerVehicleInventory) GetVehicles() []vehicle.Vehicle {
	return fwvi.vehicles
}

func (fwvi *FourWheelerVehicleInventory) GetVehicleType() string {
	return "four_wheeler"
}
