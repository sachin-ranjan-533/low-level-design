package strategy

type BillStrategy interface {
	CalculateBill(duration int) float64
}
