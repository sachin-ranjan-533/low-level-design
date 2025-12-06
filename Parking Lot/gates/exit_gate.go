package gates

import (
	"fmt"
	"parking-lot/parking_spot_manager"
	"parking-lot/ticket"
	"parking-lot/vehicle"
)

// ExitGate handles vehicle exit and keeps track of ParkingSpotManagers for each vehicle type
type ExitGate struct {
	// Map of vehicle type -> ParkingSpotManager
	managers map[vehicle.EnumVehicleType]parking_spot_manager.ParkingSpotManager
}

// Constructor for ExitGate
// Accepts a map of vehicle type to corresponding ParkingSpotManager
func NewExitGate(managers map[vehicle.EnumVehicleType]parking_spot_manager.ParkingSpotManager) *ExitGate {
	return &ExitGate{
		managers: managers,
	}
}

// RemoveVehicle un-parks a vehicle and prints its parking price
func (eg *ExitGate) RemoveVehicle(t *ticket.Ticket) {
	// Get the correct ParkingSpotManager based on the vehicle type of the parking spot
	manager, ok := eg.managers[t.ParkingSpot.GetVehicleType()]
	if !ok || manager == nil {
		fmt.Println("No manager found for vehicle type") // Safety check
		return
	}

	// Unpark the vehicle
	manager.UnparkVehicle(t.ParkingSpot)

	// Print the parking fee
	fmt.Println("Parking price is:", t.ParkingSpot.GetPrice())

	// Confirm successful un-parking
	fmt.Println("Vehicle unparked successfully. Ticket ID:", t.Id)
}
