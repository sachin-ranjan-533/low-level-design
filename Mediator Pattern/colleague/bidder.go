package colleague

import (
	"fmt"
	"mediator-pattern/mediator"
)

type Bidder struct {
	Name            string
	AuctionMediator mediator.AuctionMediator
}

func NewBidder(name string, auctionMediator mediator.AuctionMediator) *Bidder {
	b := &Bidder{
		Name:            name,
		AuctionMediator: auctionMediator,
	}
	auctionMediator.AddBidder(b)
	return b
}

func (b *Bidder) GetName() string {
	return b.Name
}

func (b *Bidder) PlaceBid(amount float64) {
	fmt.Printf("%s is placing a bid of %.2f\n", b.Name, amount)
	b.AuctionMediator.ReceiveBid(b, amount)
}
