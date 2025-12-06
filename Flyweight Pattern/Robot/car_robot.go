package robot

import "fmt"

type CarRobot struct {
	Type string
}

func NewCarRobot() *CarRobot {
	return &CarRobot{Type: "Car"}
}

func (cr *CarRobot) Display(x int, y int) {
	fmt.Println("Car Robot displayed at position:", x, y)
}
