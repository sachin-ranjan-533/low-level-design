package parking_spot

import "parking-lot/vehicle"

// ParkingSpot defines the common behavior for all types of parking spots
type ParkingSpot interface {
	// GetPrice returns the price for parking at this spot
	GetPrice() int

	// IsAvailable checks if the parking spot is currently free
	IsAvailable() bool

	// GetVehicleType returns the type of vehicle that is (or can be) parked in this spot
	GetVehicleType() vehicle.EnumVehicleType
}
