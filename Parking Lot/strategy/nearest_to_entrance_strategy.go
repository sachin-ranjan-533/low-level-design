package strategy

import "parking-lot/parking_spot"

// NearestToEntranceStrategy picks the first available spot (closest to entrance)
type NearestToEntranceStrategy struct{}

// NewNearestToEntranceStrategy creates a new NearestToEntranceStrategy instance
func NewNearestToEntranceStrategy() *NearestToEntranceStrategy {
	return &NearestToEntranceStrategy{}
}

// SelectSpot selects the first available spot from the list
func (s *NearestToEntranceStrategy) SelectSpot(spots []parking_spot.ParkingSpot) parking_spot.ParkingSpot {
	for _, spot := range spots {
		if spot.IsAvailable() {
			return spot
		}
	}
	return nil
}
