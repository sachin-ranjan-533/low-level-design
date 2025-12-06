package main

import (
	carexterior "abstract-factory/car_exterior"
	carfactory "abstract-factory/car_factory"
	carinterior "abstract-factory/car_interior"
)

func main() {
	factory := carfactory.NewCarFactory()

	interiorFactory := factory.CreatePart("Interior")
	interiorPart := interiorFactory.CreatePart("Luxury")
	interiorPart.(carinterior.CarInterior).GetInteriorDetails()

	exteriorFactory := factory.CreatePart("Exterior")
	exteriorPart := exteriorFactory.CreatePart("Normal")
	exteriorPart.(carexterior.CarExterior).GetExteriorDetails()
}
