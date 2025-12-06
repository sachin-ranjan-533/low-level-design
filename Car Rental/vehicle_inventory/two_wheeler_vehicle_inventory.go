package vehicle_inventory

import (
	"car-rental/vehicle"
)

type TwoWheelerVehicleInventory struct {
	vehicles []vehicle.Vehicle
}

func NewTwoWheelerVehicleInventory() *TwoWheelerVehicleInventory {
	return &TwoWheelerVehicleInventory{}
}

func (twvi *TwoWheelerVehicleInventory) AddVehicle(vehicle vehicle.Vehicle) {
	twvi.vehicles = append(twvi.vehicles, vehicle)
}

func (twvi *TwoWheelerVehicleInventory) GetVehicles() []vehicle.Vehicle {
	return twvi.vehicles
}

func (twvi *TwoWheelerVehicleInventory) GetVehicleType() string {
	return "two_wheeler"
}
