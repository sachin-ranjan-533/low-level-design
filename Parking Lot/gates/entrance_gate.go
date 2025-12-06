package gates

import (
	"parking-lot/parking_spot_manager"
	"parking-lot/ticket"
	"parking-lot/vehicle"
)

// EntranceGate now stores managers in a map keyed by vehicle type
type EntranceGate struct {
	managers map[vehicle.EnumVehicleType]parking_spot_manager.ParkingSpotManager
}

// Constructor accepts a map of vehicleType -> manager
func NewEntranceGate(managers map[vehicle.EnumVehicleType]parking_spot_manager.ParkingSpotManager) *EntranceGate {
	return &EntranceGate{
		managers: managers,
	}
}

// AllowEntry finds the correct manager and parks the vehicle
func (eg *EntranceGate) AllowEntry(v *vehicle.Vehicle) *ticket.Ticket {
	manager, ok := eg.managers[v.VehicleType]
	if !ok || manager == nil {
		return nil // no manager for this vehicle type
	}

	spot := manager.FindParkingSpot()
	if spot == nil {
		return nil // no free spot
	}

	return manager.ParkVehicle(spot, v)
}
