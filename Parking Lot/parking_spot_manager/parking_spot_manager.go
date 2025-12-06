package parking_spot_manager

import (
	parkingspot "parking-lot/parking_spot"
	"parking-lot/ticket"
	"parking-lot/vehicle"
)

// ParkingSpotManager defines the common behavior for managing parking spots
// It acts as an interface for different types of parking spot managers
// (e.g., TwoWheelerParkingSpotManager, FourWheelerParkingSpotManager)
type ParkingSpotManager interface {
	// AddParkingSpot adds a new parking spot to the manager
	AddParkingSpot(spot parkingspot.ParkingSpot)

	// FindParkingSpot returns the first available parking spot
	FindParkingSpot() parkingspot.ParkingSpot

	// ParkVehicle assigns a vehicle to the given parking spot and returns a Ticket
	ParkVehicle(parkingSpot parkingspot.ParkingSpot, vehicle *vehicle.Vehicle) *ticket.Ticket

	// UnparkVehicle frees the parking spot and removes the vehicle
	UnparkVehicle(parkingSpot parkingspot.ParkingSpot)
}
