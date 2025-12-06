package vehicle

import "strategy-pattern/strategy"

// NormalVehicle is a concrete type that uses a normal drive strategy.

type NormalVehicle struct {
	DriveStrategy strategy.DriveStrategy
}

// Drive executes the assigned driving strategy.

func (nv *NormalVehicle) Drive() {
	nv.DriveStrategy.Drive()
}
