package main

import "abstract-factory/abstract_factory_method"

func main() {
	car := abstract_factory_method.AbstractFactoryMethod("car", "Sedan")
	car.Drive()

	bike := abstract_factory_method.AbstractFactoryMethod("bike", "Mountain Bike")
	bike.Drive()
}
