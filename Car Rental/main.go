package main

import (
	"car-rental/bill"
	"car-rental/booking_user"
	"car-rental/factory"
	"car-rental/location"
	"car-rental/payment"
	"car-rental/reservation"
	"car-rental/store"
	"car-rental/store_manager"
	"car-rental/strategy"
	"car-rental/vehicle"
	"fmt"

	"google.golang.org/genproto/googleapis/type/datetime"
)

func main() {
	// creating location
	location := location.NewLocation("123 Main St", "Apt 4B", "Metropolis", 12345)

	// creating vehicles and inventories
	bike1 := vehicle.NewBike(1, "Yamaha FZ")
	bike2 := vehicle.NewBike(2, "Honda CBR")

	vehicleInventoryFactory := factory.NewVehicleInventoryFactory()

	twoWheelerInventory := vehicleInventoryFactory.CreateVehicleInventory("two_wheeler")
	twoWheelerInventory.AddVehicle(bike1)
	twoWheelerInventory.AddVehicle(bike2)

	car1 := vehicle.NewCar(1, "Toyota Camry")
	car2 := vehicle.NewCar(2, "Honda Accord")

	fourWheelerInventory := vehicleInventoryFactory.CreateVehicleInventory("four_wheeler")
	fourWheelerInventory.AddVehicle(car1)
	fourWheelerInventory.AddVehicle(car2)

	// creating store and initializing inventories
	store := store.NewStore(1, "Test Store", location)
	store.InitializeInventory(twoWheelerInventory)
	store.InitializeInventory(fourWheelerInventory)

	storeManager := store_manager.NewStoreManager(location)
	storeManager.AddStore(*store)

	// Below are the flow for booking a vehicle. Workflow: Get User, Retrieve store Manager for user location, retrieve stores, pick a vehicle, create a reservation, generate the bill using an billing strategy, and process the payment via payment strategy.

	// creating booking user and setting location
	bookingUser := booking_user.NewBookingUser("John", "Doe")
	bookingUser.SetLocation(location)

	// add logic to get the store manager based on location (iska function kaha pe define karnege?)

	// Getting the stores from store manager
	stores := storeManager.GetStores()

	// Get the vehicle factory inventory for the type of vehicle they want to rent
	vehicleInventory := stores[0].GetVehicleInventory("four_wheeler")

	vehicles := vehicleInventory.GetVehicles()

	// Creating a reservation for the first vehicle in the inventory
	reservation := reservation.NewReservation(1, bookingUser, &vehicles[0], datetime.DateTime{})

	store.AddReservation(reservation)

	// Applying an hourly billing strategy to calculate the total bill
	hourBillStrategy := strategy.NewHourBillStrategy()

	bill := bill.NewBill(reservation, hourBillStrategy)
	total := bill.CalculateTotal(10)

	fmt.Println("Total Bill Amount:", total)

	// Processing the payment via UPI strategy
	upiPaymentStrategy := strategy.NewUPIPaymentStrategy()
	payment := payment.NewPayment(1, total, upiPaymentStrategy)

	payment.ProcessPayment()
	fmt.Println("Payment processed successfully")

}
