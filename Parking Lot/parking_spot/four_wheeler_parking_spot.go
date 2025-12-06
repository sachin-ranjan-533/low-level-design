package parking_spot

import "parking-lot/vehicle"

// FourWheelerParkingSpot represents a parking spot for four-wheeler vehicles
type FourWheelerParkingSpot struct {
	Price   float64          // Price for parking at this spot
	IsFree  bool             // Indicates whether the spot is currently free
	Vehicle *vehicle.Vehicle // Vehicle currently parked in this spot, nil if empty
}

// Constructor for creating a new FourWheelerParkingSpot with a given price
func NewFourWheelerParkingSpot(price float64) *FourWheelerParkingSpot {
	return &FourWheelerParkingSpot{
		Price:  price,
		IsFree: true, // By default, the spot is free when created
	}
}

// GetPrice returns the parking price for this spot as an integer
func (fwps *FourWheelerParkingSpot) GetPrice() int {
	return int(fwps.Price)
}

// IsAvailable checks if the parking spot is free
func (fwps *FourWheelerParkingSpot) IsAvailable() bool {
	return fwps.IsFree
}

// GetVehicleType returns the type of vehicle currently parked in the spot
func (fwps *FourWheelerParkingSpot) GetVehicleType() vehicle.EnumVehicleType {
	if fwps.Vehicle == nil {
		// If no vehicle is parked, return a default value or handle appropriately
		return vehicle.EnumVehicleTypeCar // Defaulting to car for four-wheeler spot
	}
	return fwps.Vehicle.VehicleType
}
