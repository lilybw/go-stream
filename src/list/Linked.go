package list

import (
	"fmt"
	"strings"
)

type Node[E any] struct {
	value E
	next  *Node[E]
	prev  *Node[E]
}

type Linked[E any] struct {
	List[E]
	head   *Node[E]
	tail   *Node[E]
	length int
	equals func(E, E) bool
}

func NewLinked[E any](equals func(E, E) bool) *Linked[E] {
	return &Linked[E]{head: nil, tail: nil, length: 0, equals: equals}
}

func (l *Linked[E]) Push(element E) {
	node := &Node[E]{value: element}
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		node.next = l.head
		l.head.prev = node
		l.head = node
	}
	l.length++
}

func (l *Linked[E]) Add(element E) {
	node := &Node[E]{value: element}
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		l.tail.next = node
		node.prev = l.tail
		l.tail = node
	}
	l.length++
}

func (l *Linked[E]) Remove(element E) bool {
	if l.head == nil {
		return false
	}
	if l.equals(l.head.value, element) {
		l.head = l.head.next
		l.length--
		return true
	}
	current := l.head
	for current.next != nil {
		if l.equals(current.next.value, element) {
			current.next = current.next.next
			l.length--
			return true
		}
		current = current.next
	}
	return false
}

func (l *Linked[E]) Contains(element E) bool {
	current := l.head
	for current != nil {
		if l.equals(current.value, element) {
			return true
		}
		current = current.next
	}
	return false
}

func (l *Linked[E]) ToString() string {
	var elements []string
	current := l.head
	for current != nil {
		elements = append(elements, fmt.Sprintf("%v", current))
		current = current.next
	}
	return fmt.Sprintf("[%s]", strings.Join(elements, ", "))
}

func (l *Linked[E]) Size() int {
	return l.length
}

func (l *Linked[E]) Get(index int) E {
	current := l.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	return current.value
}

func (l *Linked[E]) Set(index int, element E) {
	current := l.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	current.value = element
}

func (l *Linked[E]) Clear() {
	l.head = nil
	l.tail = nil
	l.length = 0
}

func (l *Linked[E]) IndexOf(element E) int {
	current := l.head
	for i := 0; current != nil; i++ {
		if l.equals(current.value, element) {
			return i
		}
		current = current.next
	}
	return -1
}

func (l *Linked[E]) LastIndexOf(element E) int {
	lastIndex := -1
	current := l.head
	for i := 0; current != nil; i++ {
		if l.equals(current.value, element) {
			lastIndex = i
		}
		current = current.next
	}
	return lastIndex
}

func (l *Linked[E]) First() E {
	return l.head.value
}

func (l *Linked[E]) Last() E {
	return l.tail.value
}

func (l *Linked[E]) SubList(from, to int) List[E] {
	sub := NewLinked[E](l.equals)
	current := l.head
	for i := 0; i < from; i++ {
		current = current.next
	}
	for i := from; i < to; i++ {
		sub.Push(current.value)
		current = current.next
	}
	return sub
}
