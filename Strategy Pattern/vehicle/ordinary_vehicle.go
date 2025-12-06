package vehicle

import "strategy-pattern/strategy"

// OrdinaryVehicle is another concrete type that uses a normal drive strategy.

type OrdinaryVehicle struct {
	DriveStrategy strategy.DriveStrategy
}

// Drive executes the assigned driving strategy.

func (ov *OrdinaryVehicle) Drive() {
	ov.DriveStrategy.Drive()
}
