package main

import car "abstract-factory/create_car"

func main() {
	car := car.NewCar()
	normalCarType := car.GetCarType("normal")
	luxuryCarType := car.GetCarType("luxury")

	normalCarExterior := normalCarType.GetExterior()
	normalCarInterior := normalCarType.GetInterior()

	luxuryCarExterior := luxuryCarType.GetExterior()
	luxuryCarInterior := luxuryCarType.GetInterior()

	normalCarExterior.GetExteriorDetails()
	normalCarInterior.GetInteriorDetails()

	luxuryCarExterior.GetExteriorDetails()
	luxuryCarInterior.GetInteriorDetails()
}
