package bill

import (
	"car-rental/reservation"
	"car-rental/strategy"
)

type Bill struct {
	reservation *reservation.Reservation
	amount      float64
	strategy    strategy.BillStrategy
}

func NewBill(reservation *reservation.Reservation, strategy strategy.BillStrategy) *Bill {
	return &Bill{
		reservation: reservation,
		strategy:    strategy,
	}
}

func (b *Bill) CalculateTotal(duration int) float64 {
	b.amount = b.strategy.CalculateBill(duration)
	return b.amount
}
