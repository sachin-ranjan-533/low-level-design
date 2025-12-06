package room

type SingleRoom struct {
	RoomNumber int
}

func NewRoom(roomnumber int) *SingleRoom {
	return &SingleRoom{RoomNumber: roomnumber}
}

func (sr *SingleRoom) Accept(operation Operation) {
	operation.Visit(sr)
}
