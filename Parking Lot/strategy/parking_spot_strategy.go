package strategy

import "parking-lot/parking_spot"

// ParkingSpotStrategy defines a strategy to select a parking spot from a list
type ParkingSpotStrategy interface {
	// SelectSpot selects one parking spot from the given list based on the strategy
	SelectSpot(spots []parking_spot.ParkingSpot) parking_spot.ParkingSpot
}
