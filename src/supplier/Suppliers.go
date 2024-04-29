package supplier

type Supplier[T any] interface {
	HasNext() bool
	Next() T
}
