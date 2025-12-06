package collection

import "iterator-pattern/iterator"

type Array struct {
	items    []int
	iterator iterator.Iterator
}

func NewArray(items []int) *Array {
	return &Array{items: items}
}

func (a *Array) Iterator() iterator.Iterator {
	a.iterator = iterator.NewArrayIterator(a.items)
	return a.iterator
}
