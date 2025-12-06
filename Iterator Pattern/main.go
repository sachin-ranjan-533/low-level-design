package main

import "iterator-pattern/collection"

func main() {
	array := collection.NewArray([]int{1, 2, 3, 4, 5})
	iterator := array.Iterator()

	for iterator.HasNext() {
		value := iterator.Next()
		println(value)
	}
}
