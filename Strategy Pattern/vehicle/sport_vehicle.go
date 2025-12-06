package vehicle

import "strategy-pattern/strategy"

// SportsVehicle uses a sports drive strategy for fast performance.

type SportsVehicle struct {
	DriveStrategy strategy.DriveStrategy
}

// Drive executes the assigned sports driving strategy.

func (sv *SportsVehicle) Drive() {
	sv.DriveStrategy.Drive()
}
