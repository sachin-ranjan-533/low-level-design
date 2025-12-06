package parking_spot

import "parking-lot/vehicle"

// TwoWheelerParkingSpot represents a parking spot for two-wheeler vehicles
type TwoWheelerParkingSpot struct {
	Price   float64          // Price for parking at this spot
	IsFree  bool             // Indicates whether the spot is currently free
	Vehicle *vehicle.Vehicle // Vehicle currently parked in this spot, nil if empty
}

// Constructor for creating a new TwoWheelerParkingSpot with a given price
func NewTwoWheelerParkingSpot(price float64) *TwoWheelerParkingSpot {
	return &TwoWheelerParkingSpot{
		Price:  price,
		IsFree: true, // By default, the spot is free when created
	}
}

// GetPrice returns the parking price for this spot as an integer
func (twps *TwoWheelerParkingSpot) GetPrice() int {
	return int(twps.Price)
}

// IsAvailable checks if the parking spot is free
func (twps *TwoWheelerParkingSpot) IsAvailable() bool {
	return twps.IsFree
}

// GetVehicleType returns the type of vehicle currently parked in the spot
func (twps *TwoWheelerParkingSpot) GetVehicleType() vehicle.EnumVehicleType {
	if twps.Vehicle == nil {
		// If no vehicle is parked, return a default type (motorcycle for two-wheeler)
		return vehicle.EnumVehicleTypeMotorcycle
	}
	return twps.Vehicle.VehicleType
}
