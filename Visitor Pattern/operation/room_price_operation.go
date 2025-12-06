package operation

import (
	"fmt"
	"visitor-pattern/room"
)

type RoomPriceOperation struct{}

func NewRoomPriceOperation() *RoomPriceOperation {
	return &RoomPriceOperation{}
}

func (rop *RoomPriceOperation) Visit(r room.Room) {
	switch roomType := r.(type) {
	case *room.SingleRoom:
		rop.VisitSingleRoom(roomType)
	case *room.DoubleRoom:
		rop.VisitDoubleRoom(roomType)
	case *room.DeluxeRoom:
		rop.VisitDeluxeRoom(roomType)
	}
}

func (rop *RoomPriceOperation) VisitSingleRoom(sr *room.SingleRoom) {
	fmt.Printf("Single Room Price Operation for Room Number: %d\n", sr.RoomNumber)
}

func (rop *RoomPriceOperation) VisitDoubleRoom(dr *room.DoubleRoom) {
	fmt.Printf("Double Room Price Operation for Room Number: %d\n", dr.RoomNumber)
}

func (rop *RoomPriceOperation) VisitDeluxeRoom(dr *room.DeluxeRoom) {
	fmt.Printf("Deluxe Room Price Operation for Room Number: %d\n", dr.RoomNumber)
}
