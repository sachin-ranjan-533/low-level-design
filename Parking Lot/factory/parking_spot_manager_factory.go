package factory

import (
	parking_spot_manager "parking-lot/parking_spot_manager"
	"parking-lot/strategy"
	"parking-lot/vehicle"
)

// ParkingSpotManagerFactory is a factory to create parking spot managers
type ParkingSpotManagerFactory struct{}

// NewParkingSpotManagerFactory creates a new instance of ParkingSpotManagerFactory
func NewParkingSpotManagerFactory() *ParkingSpotManagerFactory {
	return &ParkingSpotManagerFactory{}
}

// CreateParkingSpotManager returns the appropriate parking spot manager
// based on the type of vehicle. It applies the strategy pattern to select
// the parking spot selection strategy for each manager.
func (psmf *ParkingSpotManagerFactory) CreateParkingSpotManager(vehicleType vehicle.EnumVehicleType) parking_spot_manager.ParkingSpotManager {
	switch vehicleType {
	case vehicle.EnumVehicleTypeCar:
		// For cars, use NearestToEntranceStrategy
		strategy := strategy.NewNearestToEntranceStrategy()
		return parking_spot_manager.NewFourWheelerParkingSpotManager(strategy)

	case vehicle.EnumVehicleTypeMotorcycle:
		// For motorcycles, use NearestToExitStrategy
		strategy := strategy.NewNearestToExitStrategy()
		return parking_spot_manager.NewTwoWheelerParkingSpotManager(strategy)

	default:
		// Return nil if vehicle type is not supported
		return nil
	}
}
