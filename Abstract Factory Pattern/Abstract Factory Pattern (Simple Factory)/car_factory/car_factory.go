package carfactory

import carpartfactory "abstract-factory/car_part_factory"

type CarFactory struct{}

func NewCarFactory() *CarFactory {
	return &CarFactory{}
}

func (cf *CarFactory) CreatePart(partType string) carpartfactory.CarPartFactory {
	if partType == "Interior" {
		return carpartfactory.NewCarInteriorPartFactory()
	} else if partType == "Exterior" {
		return carpartfactory.NewCarExteriorPartFactory()
	}
	return nil
}
