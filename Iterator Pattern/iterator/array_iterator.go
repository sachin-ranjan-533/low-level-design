package iterator

type ArrayIterator struct {
	items []int
	index int
}

func NewArrayIterator(items []int) *ArrayIterator {
	return &ArrayIterator{items: items, index: 0}
}

func (ai *ArrayIterator) HasNext() bool {
	return ai.index < len(ai.items)
}

func (ai *ArrayIterator) Next() int {
	value := ai.items[ai.index]
	ai.index++
	return value
}
