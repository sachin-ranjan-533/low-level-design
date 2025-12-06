package factory

import robot "flyweight-pattern/Robot"

type Factory struct{}

var robots = map[string]robot.Robot{}

func NewFactory() *Factory {
	return &Factory{}
}

func (f *Factory) GetRobot(robotType string) robot.Robot {
	if r, exists := robots[robotType]; exists {
		return r
	}

	var r robot.Robot
	switch robotType {
	case "Car":
		r = robot.NewCarRobot()
	case "Human":
		r = robot.NewHumanRobot()
	}

	robots[robotType] = r
	return r
}
