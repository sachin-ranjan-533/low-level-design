package operation

import "visitor-pattern/room"

type Operation interface {
	Visit(r room.Room)
}
