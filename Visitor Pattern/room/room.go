package room

type Operation interface {
	Visit(room Room)
}

type Room interface {
	Accept(operation Operation)
}
