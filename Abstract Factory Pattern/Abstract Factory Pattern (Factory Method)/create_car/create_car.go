package car

import cartype "abstract-factory/car_type"

type Car struct{}

func NewCar() *Car {
	return &Car{}
}

func (c *Car) GetCarType(carType string) cartype.CarType {
	if carType == "normal" {
		return cartype.NewNormalCarType()
	} else if carType == "luxury" {
		return cartype.NewLuxuryCarType()
	}
	return nil
}
