package main

import (
	"visitor-pattern/operation"
	"visitor-pattern/room"
)

func main() {
	facilityOperation := operation.NewRoomFacilityOperation()
	priceOperation := operation.NewRoomPriceOperation()

	singleRoom := room.NewRoom(101)
	doubleRoom := room.NewDOubleRoom(202)
	deluxeRoom := room.NewDeluxeRoom(303)

	singleRoom.Accept(priceOperation)
	doubleRoom.Accept(priceOperation)
	deluxeRoom.Accept(priceOperation)

	singleRoom.Accept(facilityOperation)
	doubleRoom.Accept(facilityOperation)
	deluxeRoom.Accept(facilityOperation)
}
