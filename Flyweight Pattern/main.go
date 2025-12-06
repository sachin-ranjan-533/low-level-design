package main

import "flyweight-pattern/factory"

func main() {
	factory := factory.NewFactory()
	humanRobot := factory.GetRobot("Human")
	humanRobot.Display(10, 20)
	carRobot := factory.GetRobot("Car")
	carRobot.Display(30, 40)
	humanRobot = factory.GetRobot("Human")
	humanRobot.Display(50, 60)
}
