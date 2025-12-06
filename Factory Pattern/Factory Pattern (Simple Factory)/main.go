package main

func main() {
	factory := &Factory{}

	car := factory.CreateVehicle("Car")
	car.Drive()

	bike := factory.CreateVehicle("Bike")
	bike.Drive()
}
