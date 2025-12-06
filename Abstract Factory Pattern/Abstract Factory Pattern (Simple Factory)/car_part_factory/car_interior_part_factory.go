package carpartfactory

import carinterior "abstract-factory/car_interior"

type CarInteriorPartFactory struct{}

func NewCarInteriorPartFactory() *CarInteriorPartFactory {
	return &CarInteriorPartFactory{}
}

func (cepf *CarInteriorPartFactory) CreatePart(classType string) interface{} {
	if classType == "Normal" {
		return carinterior.NewNormalCarInterior()
	} else if classType == "Luxury" {
		return carinterior.NewLuxuryCarInterior()
	}
	return nil
}
