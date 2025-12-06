package operation

import (
	"fmt"
	"visitor-pattern/room"
)

type RoomFacilityOperation struct{}

func NewRoomFacilityOperation() *RoomFacilityOperation {
	return &RoomFacilityOperation{}
}

func (rfo *RoomFacilityOperation) Visit(r room.Room) {
	switch roomType := r.(type) {
	case *room.SingleRoom:
		rfo.VisitSingleRoom(roomType)
	case *room.DoubleRoom:
		rfo.VisitDoubleRoom(roomType)
	case *room.DeluxeRoom:
		rfo.VisitDeluxeRoom(roomType)
	}
}

func (rfo *RoomFacilityOperation) VisitSingleRoom(sr *room.SingleRoom) {
	fmt.Println("Single Room Facility Operation for Room Number:", sr.RoomNumber)
}

func (rfo *RoomFacilityOperation) VisitDoubleRoom(dr *room.DoubleRoom) {
	fmt.Println("Double Room Facility Operation for Room Number:", dr.RoomNumber)
}

func (rfo *RoomFacilityOperation) VisitDeluxeRoom(dr *room.DeluxeRoom) {
	fmt.Println("Deluxe Room Facility Operation for Room Number:", dr.RoomNumber)
}
