package gocontainers

type Queue[T comparable] struct {
	elements []T
}

func NewQueue[T comparable]() *Queue[T] {
	return &Queue[T]{elements: make([]T, 0)}
}

func (q *Queue[T]) Enqueue(element T) {
	q.elements = append(q.elements, element)
}

func (q *Queue[T]) Dequeue() T {
	if len(q.elements) == 0 {
		panic("Dequeue from empty queue")
	}

	front := q.elements[0]
	q.elements = q.elements[1:]
	return front
}

func (q *Queue[T]) Size() int {
	return len(q.elements)
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.elements) == 0
}

func (q *Queue[T]) Clear() {
	q.elements = nil
}
