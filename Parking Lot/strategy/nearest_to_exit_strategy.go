package strategy

import "parking-lot/parking_spot"

// NearestToExitStrategy picks the last available spot (closest to exit)
type NearestToExitStrategy struct{}

// NewNearestToExitStrategy creates a new NearestToExitStrategy instance
func NewNearestToExitStrategy() *NearestToExitStrategy {
	return &NearestToExitStrategy{}
}

// SelectSpot selects the last available spot from the list
func (s *NearestToExitStrategy) SelectSpot(spots []parking_spot.ParkingSpot) parking_spot.ParkingSpot {
	for i := len(spots) - 1; i >= 0; i-- {
		if spots[i].IsAvailable() {
			return spots[i]
		}
	}
	return nil
}
