package booking_user

import "car-rental/location"

type BookingUser struct {
	FirstName string
	LastName  string
	location  *location.Location
}

func NewBookingUser(FirstName string, LastName string) *BookingUser {
	return &BookingUser{
		FirstName: FirstName,
		LastName:  LastName,
	}
}

func (bu *BookingUser) SetLocation(location *location.Location) {
	bu.location = location
}
