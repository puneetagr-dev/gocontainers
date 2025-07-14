package gocontainers

type DLL[T comparable] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

type Node[T comparable] struct {
	element T
	prev    *Node[T]
	next    *Node[T]
}

func NewNode[T comparable](element T) *Node[T] {
	return &Node[T]{element: element}
}

func (n *Node[T]) Get() T {
	return n.element
}

func (n *Node[T]) Update(val T) {
	n.element = val
}

func NewDLL[T comparable]() *DLL[T] {
	return &DLL[T]{head: nil, tail: nil, size: 0}
}

func (dll *DLL[T]) AddFront(node *Node[T]) {
	if dll.head == nil {
		dll.head = node
		dll.tail = node
	} else {
		node.next = dll.head
		dll.head.prev = node
		dll.head = node
	}
	dll.size++
}

func (dll *DLL[T]) AddBack(node *Node[T]) {
	if dll.tail == nil {
		dll.head = node
		dll.tail = node
	} else {
		node.prev = dll.tail
		dll.tail.next = node
		dll.tail = node
	}
	dll.size++
}

func (dll *DLL[T]) RemoveFront() {
	if dll.head == nil {
		return
	}
	if dll.head == dll.tail {
		dll.head = nil
		dll.tail = nil
		dll.size = 0
		return
	}
	dll.head = dll.head.next
	dll.head.prev = nil
	dll.size--
}

func (dll *DLL[T]) RemoveBack() {
	if dll.tail == nil {
		return
	}
	if dll.head == dll.tail {
		dll.head = nil
		dll.tail = nil
		dll.size = 0
		return
	}
	dll.tail = dll.tail.prev
	dll.tail.next = nil
	dll.size--
}

func (dll *DLL[T]) Size() int {
	return dll.size
}

func (dll *DLL[T]) IsEmpty() bool {
	return dll.head == nil
}

func (dll *DLL[T]) GetFront() *Node[T] {
	return dll.head
}

func (dll *DLL[T]) GetBack() *Node[T] {
	return dll.tail
}

func (dll *DLL[T]) DeleteMatch(element T) {
	if dll.head == nil {
		return
	}

	current := dll.head
	for current != nil {
		if current.element == element {
			next := current.next
			if current.prev != nil {
				current.prev.next = current.next
			} else {
				dll.head = current.next
			}
			if current.next != nil {
				current.next.prev = current.prev
			} else {
				dll.tail = current.prev
			}
			dll.size--
			current = next
		} else {
			current = current.next
		}
	}
}

func (dll *DLL[T]) DeleteNode(node *Node[T]) {
	if dll.Size() == 0 {
		return
	}

	if dll.size == 1 && dll.head == node {
		dll.head = nil
		dll.tail = nil
		dll.size = 0
		return
	}

	if dll.head == node {
		dll.head = dll.head.next
		dll.tail.prev = nil
		dll.size--
		return
	}

	if dll.tail == node {
		dll.tail = dll.tail.prev
		dll.head.next = nil
		dll.size--
		return
	}

	node.prev.next, node.next.prev = node.next, node.prev
	dll.size--
}

// Clear removes all elements from the DLL
func (dll *DLL[T]) Clear() {
	dll.head = nil
	dll.tail = nil
	dll.size = 0
}

type Iterator[T comparable] struct {
	dll     *DLL[T]
	current *Node[T]
}

func (dll *DLL[T]) Iterator() *Iterator[T] {
	return &Iterator[T]{dll: dll, current: dll.head}
}

func (it *Iterator[T]) HasNext() bool {
	return it.current != nil
}

func (it *Iterator[T]) Next() T {
	if it.current == nil {
		panic("Iterator is at the end")
	}

	elem := it.current.element
	it.current = it.current.next
	return elem
}
