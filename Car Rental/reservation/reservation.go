package reservation

import (
	"car-rental/booking_user"
	"car-rental/vehicle"

	"google.golang.org/genproto/googleapis/type/datetime"
)

type Reservation struct {
	id              int
	bookingUser     *booking_user.BookingUser
	vehicle         *vehicle.Vehicle
	bookingDateTime datetime.DateTime
}

func NewReservation(id int, bookingUser *booking_user.BookingUser, vehicle *vehicle.Vehicle, bookingDateTime datetime.DateTime) *Reservation {
	return &Reservation{
		id:              id,
		bookingUser:     bookingUser,
		vehicle:         vehicle,
		bookingDateTime: bookingDateTime,
	}
}
