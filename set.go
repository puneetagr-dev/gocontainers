package gocontainers

type Set[T comparable] struct {
	elements map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{elements: make(map[T]struct{})}
}

func (s *Set[T]) Add(element T) {
	s.elements[element] = struct{}{}
}

func (s *Set[T]) Remove(element T) {
	delete(s.elements, element)
}

func (s *Set[T]) Contains(element T) bool {
	_, exists := s.elements[element]
	return exists
}

func (s *Set[T]) Size() int {
	return len(s.elements)
}

func (s *Set[T]) Clear() {
	s.elements = make(map[T]struct{})
}

func (s *Set[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Set[T]) ToSlice() []T {
	result := make([]T, 0, len(s.elements))
	for element := range s.elements {
		result = append(result, element)
	}
	return result
}

// Union returns a new set containing all elements from both sets.
// If an element is present in both sets, it will only appear once in the result.
func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for element := range s.elements {
		result.Add(element)
	}
	for element := range other.elements {
		result.Add(element)
	}
	return result
}

func (s *Set[T]) Equal(other *Set[T]) bool {
	if s.Size() != other.Size() {
		return false
	}

	for element := range s.elements {
		if !other.Contains(element) {
			return false
		}
	}
	return true
}
