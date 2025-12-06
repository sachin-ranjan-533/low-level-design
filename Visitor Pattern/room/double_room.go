package room

type DoubleRoom struct {
	RoomNumber int
}

func NewDOubleRoom(roomnumber int) *DoubleRoom {
	return &DoubleRoom{RoomNumber: roomnumber}
}

func (dr *DoubleRoom) Accept(operation Operation) {
	operation.Visit(dr)
}
