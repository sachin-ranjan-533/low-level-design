package carpartfactory

import carexterior "abstract-factory/car_exterior"

type CarExteriorPartFactory struct{}

func NewCarExteriorPartFactory() *CarExteriorPartFactory {
	return &CarExteriorPartFactory{}
}

func (cepf *CarExteriorPartFactory) CreatePart(classType string) interface{} {
	if classType == "Normal" {
		return carexterior.NewNormalCarExterior()
	} else if classType == "Luxury" {
		return carexterior.NewLuxuryCarExterior()
	}
	return nil
}
