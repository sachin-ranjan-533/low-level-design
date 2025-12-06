package main

import (
	"strategy-pattern/strategy"
	"strategy-pattern/vehicle"
)

func main() {
	// creating normal drive strategy
	normalStrategy := &strategy.NormalDriveStrategy{}

	// creating sport drive strategy
	sportStrategy := &strategy.SportDriveStrategy{}

	//  creating normal vehicle with normal drive strategy
	normalVehicle := &vehicle.NormalVehicle{DriveStrategy: normalStrategy}
	normalVehicle.Drive()

	// creating sports vehicle with sport drive strategy
	sportsVehicle := &vehicle.SportsVehicle{DriveStrategy: sportStrategy}
	sportsVehicle.Drive()

	// creating ordinary vehicle with normal drive strategy
	ordinaryVehicle := &vehicle.OrdinaryVehicle{DriveStrategy: normalStrategy}
	ordinaryVehicle.Drive()
}
