package mediator

import "fmt"

type Auction struct {
	Bidders []Colleague
}

func NewAuction() *Auction {
	return &Auction{Bidders: []Colleague{}}
}

func (a *Auction) AddBidder(bidder Colleague) {
	a.Bidders = append(a.Bidders, bidder)
}

func (a *Auction) ReceiveBid(bidder Colleague, amount float64) {
	fmt.Printf("%s placed a bid of %.2f\n", bidder.GetName(), amount)

	for _, b := range a.Bidders {
		if b.GetName() != bidder.GetName() {
			fmt.Printf("Notifying %s of new bid %.2f\n", b.GetName(), amount)
		}
	}
}
