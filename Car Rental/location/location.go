package location

type Location struct {
	address1 string
	address2 string
	city     string
	pincode  int
}

func NewLocation(address1 string, address2 string, city string, pincode int) *Location {
	return &Location{
		address1: address1,
		address2: address2,
		city:     city,
		pincode:  pincode,
	}
}
