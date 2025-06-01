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

func NewDLL[T comparable]() *DLL[T] {
	return &DLL[T]{head: nil, tail: nil, size: 0}
}

func (dll *DLL[T]) AddFront(element T) {
	node := &Node[T]{element: element}
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

func (dll *DLL[T]) AddBack(element T) {
	node := &Node[T]{element: element}
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

func (dll *DLL[T]) GetFront() T {
	return dll.head.element
}

func (dll *DLL[T]) GetBack() T {
	return dll.tail.element
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

// Clear removes all elements from the DLL
func (dll *DLL[T]) Clear() {
	dll.head = nil
	dll.tail = nil
	dll.size = 0
}
