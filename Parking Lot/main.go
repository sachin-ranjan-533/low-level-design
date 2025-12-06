package main

import (
	"fmt"
	"parking-lot/factory"
	"parking-lot/gates"
	"parking-lot/parking_spot"
	"parking-lot/parking_spot_manager"
	"parking-lot/vehicle"
)

func main() {
	// Create parking spot manager factory
	managerFactory := factory.NewParkingSpotManagerFactory()

	// Create managers for each vehicle type (as interface type ParkingSpotManager)
	twoWManager := managerFactory.CreateParkingSpotManager(vehicle.EnumVehicleTypeMotorcycle)
	fourWManager := managerFactory.CreateParkingSpotManager(vehicle.EnumVehicleTypeCar)

	// Map of vehicle type -> manager
	managers := map[vehicle.EnumVehicleType]parking_spot_manager.ParkingSpotManager{
		vehicle.EnumVehicleTypeMotorcycle: twoWManager,
		vehicle.EnumVehicleTypeCar:        fourWManager,
	}

	// Create entrance and exit gates using the managers map
	entranceGate := gates.NewEntranceGate(managers)
	exitGate := gates.NewExitGate(managers)

	// Add parking spots
	twoWManager.AddParkingSpot(parking_spot.NewTwoWheelerParkingSpot(10))
	twoWManager.AddParkingSpot(parking_spot.NewTwoWheelerParkingSpot(15))
	fourWManager.AddParkingSpot(parking_spot.NewFourWheelerParkingSpot(20))
	fourWManager.AddParkingSpot(parking_spot.NewFourWheelerParkingSpot(25))

	// Create vehicles
	motorcycle := vehicle.NewVehicle(1, vehicle.EnumVehicleTypeMotorcycle)
	car := vehicle.NewVehicle(2, vehicle.EnumVehicleTypeCar)

	// Allow entry
	ticket1 := entranceGate.AllowEntry(motorcycle)
	ticket2 := entranceGate.AllowEntry(car)

	if ticket1 != nil {
		fmt.Println("Motorcycle parked. Ticket ID:", ticket1.Id)
	} else {
		fmt.Println("No spot available for motorcycle")
	}

	if ticket2 != nil {
		fmt.Println("Car parked. Ticket ID:", ticket2.Id)
	} else {
		fmt.Println("No spot available for car")
	}

	// Remove vehicles on exit
	if ticket1 != nil {
		exitGate.RemoveVehicle(ticket1)
	}
	if ticket2 != nil {
		exitGate.RemoveVehicle(ticket2)
	}
}
