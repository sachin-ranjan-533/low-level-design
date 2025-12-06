package ticket

import (
	parkingspot "parking-lot/parking_spot"

	"google.golang.org/genproto/googleapis/type/datetime"
)

// Ticket represents a parking ticket issued when a vehicle enters the parking lot
type Ticket struct {
	Id          int                     // Unique ticket ID
	ParkingSpot parkingspot.ParkingSpot // The parking spot assigned to the vehicle
	EntryTime   datetime.DateTime       // Timestamp of vehicle entry
}

// NewTicket creates a new Ticket with the given ID and ParkingSpot
// The EntryTime is initialized to the zero value of datetime.DateTime
func NewTicket(id int, parkingSpot parkingspot.ParkingSpot) *Ticket {
	return &Ticket{
		Id:          id,
		ParkingSpot: parkingSpot,
		EntryTime:   datetime.DateTime{}, // Can be updated when vehicle enters
	}
}
