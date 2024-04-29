package list

import (
	"errors"
	"fmt"
	"strings"
)

// Array is a struct that holds a list of T
type Array[T any] struct {
	List[T]
	elements []T
	equals   func(T, T) bool // Function to compare elements
}

// NewArray creates a new array with a default capacity of 0 using the provided equals function for further operations
//
// This implementation is not thread-safe, see list/ConcArray for a thread-safe implementation
func NewArray[T any](equals func(T, T) bool) *Array[T] {
	return &Array[T]{elements: make([]T, 0), equals: equals}
}

// NewArrayWithCapacity creates a new array with a specified capacity.
// If the capacity is negative, the result of NewArray is returned instead (defaulting to a capacity of 0)
func NewArrayWithCapacity[T any](capacity int, equals func(T, T) bool) (array *Array[T], err error) {
	if capacity < 0 {
		return NewArray(equals), errors.New("capacity must be non-negative")
	}
	return &Array[T]{elements: make([]T, 0, capacity), equals: equals}, nil
}

func (l *Array[T]) Add(element T) {
	l.elements = append(l.elements, element)
}

func (l *Array[T]) Remove(element T) bool {
	index := l.IndexOf(element)
	if index == -1 {
		return false
	}
	l.elements = append(l.elements[:index], l.elements[index+1:]...)
	return true
}

func (l *Array[T]) Contains(element T) bool {
	for _, e := range l.elements {
		if l.equals(e, element) {
			return true
		}
	}
	return false
}

func (l *Array[T]) ToString() string {
	var elements []string
	for _, e := range l.elements {
		elements = append(elements, fmt.Sprintf("%v", e))
	}
	return fmt.Sprintf("[%s]", strings.Join(elements, ", "))
}

func (l *Array[T]) Size() int {
	return len(l.elements)
}

func (l *Array[T]) Get(index int) T {
	return l.elements[index]
}

func (l *Array[T]) Set(index int, element T) {
	l.elements[index] = element
}

func (l *Array[T]) Clear() {
	l.elements = make([]T, 0, cap(l.elements))
}

func (l *Array[T]) IndexOf(element T) int {
	for i, e := range l.elements {
		if l.equals(e, element) {
			return i
		}
	}
	return -1
}

func (l *Array[T]) LastIndexOf(element T) int {
	for i := len(l.elements) - 1; i >= 0; i-- {
		if l.equals(l.elements[i], element) {
			return i
		}
	}
	return -1
}

func (l *Array[T]) First() T {
	return l.elements[0]
}

func (l *Array[T]) Last() T {
	return l.elements[len(l.elements)-1]
}

func (l *Array[T]) SubList(from, to int) *Array[T] {

	return &Array[T]{elements: l.SubSlice(from, to), equals: l.equals}
}

func (l *Array[T]) SubSlice(from, to int) []T {
	return l.elements[from:to]
}

// AddAll adds all elements from another list
func (l *Array[T]) AddAll(other *Array[T]) {
	l.elements = append(l.elements, other.elements...)
}

// RemoveAll removes all elements from another list
func (l *Array[T]) RemoveAll(other *Array[T]) {
	//TODO: Optimize
	for _, e := range other.elements {
		l.Remove(e)
	}
}

// RetainAll retains all elements from another list
func (l *Array[T]) RetainAll(other *Array[T]) {
	//TODO: Optimize
	for _, e := range l.elements {
		if !other.Contains(e) {
			l.Remove(e)
		}
	}
}

// IsEmpty checks if the list is empty
func (l *Array[T]) IsEmpty() bool {
	return len(l.elements) == 0
}

// Equals checks if two lists are equal
func (this *Array[T]) Equals(other *Array[T]) bool {
	if other == nil || this.Size() != other.Size() {
		return false
	}

	if !other.equals(this.elements[0], this.elements[0]) {
		return false
	}
	for i, e := range this.elements {
		if !this.equals(e, other.Get(i)) {
			return false
		}
	}
	return true
}

// Clone returns a shallow copy of the list
func (l *Array[T]) Clone() *Array[T] {
	return &Array[T]{elements: l.elements, equals: l.equals}
}

// ToSlice returns a slice of the list
func (l *Array[T]) ToSlice() []T {
	return l.elements
}

// ForEach applies a function to each element in the list
func (l *Array[T]) ForEach(f func(T)) {
	for _, e := range l.elements {
		f(e)
	}
}

// Filter returns a new list with elements that satisfy a predicate
func (l *Array[T]) Filter(f func(T) bool) *Array[T] {
	filtered, _ := NewArrayWithCapacity[T](cap(l.elements), l.equals)
	for _, e := range l.elements {
		if f(e) {
			filtered.Add(e)
		}
	}
	return filtered
}

// Map applies a function f to each element of list L and returns a new list of a different type K.
func Map[T any, K any](l *Array[T], f func(T) K, newEquals func(K, K) bool) *Array[K] {
	mapped, _ := NewArrayWithCapacity[K](cap(l.elements), newEquals)
	for _, e := range l.elements {
		mapped.Add(f(e))
	}
	return mapped
}

func (this *Array[T]) Iterator() *ArrayIterator[T] {
	return NewArrayIterator(this.elements)
}
