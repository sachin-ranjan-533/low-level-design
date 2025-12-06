package main

import (
	"mediator-pattern/colleague"
	"mediator-pattern/mediator"
)

func main() {
	auction := mediator.NewAuction()

	bidder1 := colleague.NewBidder("Alice", auction)
	bidder2 := colleague.NewBidder("Bob", auction)
	bidder3 := colleague.NewBidder("Charlie", auction)
	bidder1.PlaceBid(100.0)
	bidder2.PlaceBid(150.0)
	bidder3.PlaceBid(200.0)
}
