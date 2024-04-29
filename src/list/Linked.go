package list

type Node[E any] struct {
	value E
	next  *Node[E]
}

type Linked[E any] struct {
	head   *Node[E]
	tail   *Node[E]
	length int
}
