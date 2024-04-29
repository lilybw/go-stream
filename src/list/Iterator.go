package list

type Iterator[T any] interface {
	HasNext() bool
	Next() T
}

type ArrayIterator[T any] struct {
	Iterator[T]
	slice []T
	index int
}

func NewArrayIterator[T any](slice []T) *ArrayIterator[T] {
	return &ArrayIterator[T]{slice: slice, index: -1}
}
func (i *ArrayIterator[T]) HasNext() bool {
	return i.index+1 < len(i.slice)
}

func (i *ArrayIterator[T]) Next() T {
	i.index++
	return i.slice[i.index]
}
