package stream

import (
	"go-stream/src/list"
)

type Stream[T any] interface {
	ForEach(f func(T))
	Filter(f func(T) bool) *Stream[T]
	//Come on go... Map[K any](f func(T) K, newEquals func(K, K) bool) *Stream[K]
	Reduce(f func(T, T) T) *T
}

type IteratorStream[T any] struct {
	Stream[T]
	iterator list.Iterator[T]
}

func Of[E any](elements ...E) *IteratorStream[E] {
	return &IteratorStream[E]{iterator: list.NewArrayIterator(elements)}
}
