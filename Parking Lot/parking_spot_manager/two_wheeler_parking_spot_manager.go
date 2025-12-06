package parking_spot_manager

import (
	parkingspot "parking-lot/parking_spot"
	"parking-lot/strategy"
	"parking-lot/ticket"
	"parking-lot/vehicle"
)

// TwoWheelerParkingSpotManager manages parking spots for two-wheeler vehicles
type TwoWheelerParkingSpotManager struct {
	parkingSpots        []parkingspot.TwoWheelerParkingSpot // List of managed two-wheeler parking spots
	parkingSpotStrategy strategy.ParkingSpotStrategy        // Strategy to choose parking spots
}

// Constructor to create a new TwoWheelerParkingSpotManager with a parking strategy
func NewTwoWheelerParkingSpotManager(parkingSpotStrategy strategy.ParkingSpotStrategy) *TwoWheelerParkingSpotManager {
	return &TwoWheelerParkingSpotManager{
		parkingSpotStrategy: parkingSpotStrategy,
	}
}

// AddParkingSpot adds a new parking spot to the manager
func (twpsm *TwoWheelerParkingSpotManager) AddParkingSpot(spot parkingspot.ParkingSpot) {
	// Type assert the generic ParkingSpot interface to TwoWheelerParkingSpot
	twpsm.parkingSpots = append(twpsm.parkingSpots, *spot.(*parkingspot.TwoWheelerParkingSpot))
}

// FindParkingSpot returns the first available parking spot
func (twpsm *TwoWheelerParkingSpotManager) FindParkingSpot() parkingspot.ParkingSpot {
	for i := range twpsm.parkingSpots {
		if twpsm.parkingSpots[i].IsAvailable() {
			return &twpsm.parkingSpots[i]
		}
	}
	return nil // No available spot found
}

// ParkVehicle assigns a vehicle to the given parking spot and returns a new Ticket
func (twpsm *TwoWheelerParkingSpotManager) ParkVehicle(parkingSpot parkingspot.ParkingSpot, vehicle *vehicle.Vehicle) *ticket.Ticket {
	spot := parkingSpot.(*parkingspot.TwoWheelerParkingSpot) // Type assertion
	spot.Vehicle = vehicle
	spot.IsFree = false
	return ticket.NewTicket(1, parkingSpot) // Create and return a ticket
}

// UnparkVehicle frees up the parking spot and removes the vehicle
func (twpsm *TwoWheelerParkingSpotManager) UnparkVehicle(parkingSpot parkingspot.ParkingSpot) {
	spot := parkingSpot.(*parkingspot.TwoWheelerParkingSpot) // Type assertion
	spot.Vehicle = nil
	spot.IsFree = true
}
