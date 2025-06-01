package gocontainers

type Stack[T comparable] struct {
	elements []T
}

func NewStack[T comparable]() *Stack[T] {
	return &Stack[T]{elements: make([]T, 0)}
}

func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) Pop() T {
	if len(s.elements) == 0 {
		panic("Pop from empty stack")
	}
	lastIndex := len(s.elements) - 1
	lastElement := s.elements[lastIndex]
	s.elements = s.elements[:lastIndex]
	return lastElement
}

func (s *Stack[T]) Size() int {
	return len(s.elements)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack[T]) Clear() {
	s.elements = nil
}

func (s *Stack[T]) Peek() T {
	if len(s.elements) == 0 {
		panic("Peek from empty stack")
	}
	return s.elements[len(s.elements)-1]
}
