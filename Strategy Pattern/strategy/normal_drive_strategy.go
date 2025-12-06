package strategy

// NormalDrive implements DriveStrategy for normal vehicles.

type NormalDriveStrategy struct{}

func (nds *NormalDriveStrategy) Drive() {
	println("Driving in normal mode.")
}
