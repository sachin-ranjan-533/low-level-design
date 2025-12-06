package strategy

type HourBillStrategy struct{}

func NewHourBillStrategy() *HourBillStrategy {
	return &HourBillStrategy{}
}

func (hbs *HourBillStrategy) CalculateBill(duration int) float64 {
	return float64(duration) * 10.0
}
