package strategy

type MinuteBillStrategy struct{}

func NewMinuteBillStrategy() *MinuteBillStrategy {
	return &MinuteBillStrategy{}
}

func (mbs *MinuteBillStrategy) CalculateBill(duration int) float64 {
	return float64(duration) * 1.0
}
