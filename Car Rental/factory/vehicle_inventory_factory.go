package factory

import (
	"car-rental/vehicle_inventory"
)

type VehicleInventoryFactory struct{}

func NewVehicleInventoryFactory() *VehicleInventoryFactory {
	return &VehicleInventoryFactory{}
}

func (vif *VehicleInventoryFactory) CreateVehicleInventory(vehicleType string) vehicle_inventory.VehicleInventory {
	if vehicleType == "four_wheeler" {
		return vehicle_inventory.NewFourWheelerVehicleInventory()
	} else if vehicleType == "two_wheeler" {
		return vehicle_inventory.NewTwoWheelerVehicleInventory()
	}
	return nil
}
