package strategy

// SportsDrive implements DriveStrategy for sports vehicles.

type SportDriveStrategy struct{}

func (sds *SportDriveStrategy) Drive() {
	println("Driving in sporty mode.")
}
