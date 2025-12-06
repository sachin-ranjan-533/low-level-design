package room

type DeluxeRoom struct {
	RoomNumber int
}

func NewDeluxeRoom(roomnumber int) *DeluxeRoom {
	return &DeluxeRoom{RoomNumber: roomnumber}
}

func (dr *DeluxeRoom) Accept(operation Operation) {
	operation.Visit(dr)
}
