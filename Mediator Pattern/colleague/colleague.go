package colleague

// Colleague interface is implemented by Bidder or any participant in the auction.
// It mirrors the mediator.Colleague interface but is defined here for clarity.
// (You don’t strictly need this file, but it’s useful for structure.)
type Colleague interface {
	GetName() string
	PlaceBid(amount float64)
}
