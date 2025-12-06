package vehicle

// EnumVehicleType defines the type of vehicle as a string
type EnumVehicleType string

// Supported vehicle types
const (
	EnumVehicleTypeCar        EnumVehicleType = "car"
	EnumVehicleTypeTruck      EnumVehicleType = "truck"
	EnumVehicleTypeMotorcycle EnumVehicleType = "motorcycle"
)

// Vehicle represents a vehicle entering the parking lot
type Vehicle struct {
	Id          int             // Unique vehicle ID
	VehicleType EnumVehicleType // Type of the vehicle (Car, Truck, Motorcycle)
}

// NewVehicle creates a new Vehicle instance with the given ID and type
func NewVehicle(id int, Type EnumVehicleType) *Vehicle {
	return &Vehicle{
		Id:          id,
		VehicleType: Type,
	}
}
