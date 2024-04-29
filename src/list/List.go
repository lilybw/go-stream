package list

type List[T any] interface {
	// Add adds an element to the list
	Add(element T)
	// Remove removes an element from the list using IndexOf
	Remove(element T) bool
	// Contains checks if the list contains an element
	Contains(element T) bool
	// String returns a string representation of the list
	ToString() string
	// Size returns the size of the list
	Size() int
	// Get returns the element at the specified index
	Get(index int) T
	// Set sets the element at the specified index
	Set(index int, element T)
	// Clear removes all elements from the list
	Clear()
	// Get the index of an element
	IndexOf(element T) int
	// Get the last index of an element
	LastIndexOf(element T) int
	// Get the first element in the list
	First() T
	// Get the last element in the list
	Last() T
	// Get a sublist of the list
	SubList(from, to int) List[T]
	// Get a subslice of the list
	SubSlice(from, to int) []T
	// AddAll adds all elements from another list
	AddAll(other List[T])
	// RemoveAll removes all elements from another list
	RemoveAll(other List[T])
	// RetainAll retains all elements from another list
	RetainAll(other List[T])
	// IsEmpty checks if the list is empty
	IsEmpty() bool
	// Equals checks if two lists are equal
	Equals(other List[T]) bool
	Iterator() Iterator[T]
}
