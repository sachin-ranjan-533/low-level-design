package mediator

type Colleague interface {
	GetName() string
}

type AuctionMediator interface {
	AddBidder(bidder Colleague)
	ReceiveBid(bidder Colleague, amount float64)
}
