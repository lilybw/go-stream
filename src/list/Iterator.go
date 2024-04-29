package list

import (
	"go-stream/src/supplier"
)

type Iterator[T any] interface {
	supplier.Supplier[T]
}

type ArrayIterator[T any] struct {
	Iterator[T]
	slice []T
	index int
}

func NewArrayIterator[T any](slice []T) *ArrayIterator[T] {
	return &ArrayIterator[T]{slice: slice, index: -1}
}
func (iter *ArrayIterator[T]) HasNext() bool {
	return iter.index+1 < len(iter.slice)
}

func (iter *ArrayIterator[T]) Next() T {
	iter.index++
	return iter.slice[iter.index]
}
