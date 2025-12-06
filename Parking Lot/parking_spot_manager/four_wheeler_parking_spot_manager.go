package parking_spot_manager

import (
	"parking-lot/parking_spot"
	"parking-lot/strategy"
	"parking-lot/ticket"
	"parking-lot/vehicle"
)

// FourWheelerParkingSpotManager manages parking spots for four-wheeler vehicles
type FourWheelerParkingSpotManager struct {
	parkingSpots        []parking_spot.FourWheelerParkingSpot // List of concrete four-wheeler parking spots
	parkingSpotStrategy strategy.ParkingSpotStrategy          // Strategy to choose parking spots
}

// Constructor to create a new FourWheelerParkingSpotManager with a parking strategy
func NewFourWheelerParkingSpotManager(parkingSpotStrategy strategy.ParkingSpotStrategy) *FourWheelerParkingSpotManager {
	return &FourWheelerParkingSpotManager{
		parkingSpotStrategy: parkingSpotStrategy,
		parkingSpots:        make([]parking_spot.FourWheelerParkingSpot, 0),
	}
}

// AddParkingSpot adds a new parking spot to the manager
func (fwpsm *FourWheelerParkingSpotManager) AddParkingSpot(spot parking_spot.ParkingSpot) {
	// Type assert the generic ParkingSpot to FourWheelerParkingSpot
	fwpsm.parkingSpots = append(fwpsm.parkingSpots, *spot.(*parking_spot.FourWheelerParkingSpot))
}

// FindParkingSpot selects a parking spot using the strategy
func (fwpsm *FourWheelerParkingSpotManager) FindParkingSpot() parking_spot.ParkingSpot {
	if fwpsm.parkingSpotStrategy != nil {
		// Convert concrete slice to interface slice
		spots := make([]parking_spot.ParkingSpot, len(fwpsm.parkingSpots))
		for i := range fwpsm.parkingSpots {
			spots[i] = &fwpsm.parkingSpots[i] // pointer to concrete spot
		}
		return fwpsm.parkingSpotStrategy.SelectSpot(spots)
	}
	return nil
}

// ParkVehicle assigns a vehicle to the given parking spot and returns a new Ticket
func (fwpsm *FourWheelerParkingSpotManager) ParkVehicle(parkingSpot parking_spot.ParkingSpot, vehicle *vehicle.Vehicle) *ticket.Ticket {
	spot := parkingSpot.(*parking_spot.FourWheelerParkingSpot)
	spot.Vehicle = vehicle
	spot.IsFree = false
	return ticket.NewTicket(1, parkingSpot)
}

// UnparkVehicle frees up the parking spot and removes the vehicle
func (fwpsm *FourWheelerParkingSpotManager) UnparkVehicle(parkingSpot parking_spot.ParkingSpot) {
	spot := parkingSpot.(*parking_spot.FourWheelerParkingSpot)
	spot.Vehicle = nil
	spot.IsFree = true
}
