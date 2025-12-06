package robot

import "fmt"

type HumanRobot struct {
	Type string
}

func NewHumanRobot() *HumanRobot {
	return &HumanRobot{Type: "Human"}
}

func (hr *HumanRobot) Display(x int, y int) {
	fmt.Println("Human Robot displayed at position:", x, y)
}
