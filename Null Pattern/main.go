package main

func main() {
	factory := &Factory{}

	vehicle1 := factory.CreateVehicle("Bike")
	vehicle1.Drive()

	vehicle2 := factory.CreateVehicle("Car")
	vehicle2.Drive()

	vehicle3 := factory.CreateVehicle("Truck")
	vehicle3.Drive()
}
