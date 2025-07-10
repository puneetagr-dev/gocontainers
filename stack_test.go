package gocontainers

import (
	"testing"
)

func TestStack_Push(t *testing.T) {
	// Test pushing elements
	s := NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	// Verify elements are in correct order (last pushed is at the end)
	if len(s.elements) != 3 {
		t.Errorf("Stack should have 3 elements, got %d", len(s.elements))
	}
	if s.elements[0] != 1 {
		t.Error("First element should be 1")
	}
	if s.elements[1] != 2 {
		t.Error("Second element should be 2")
	}
	if s.elements[2] != 3 {
		t.Error("Third element should be 3")
	}
}

func TestStack_Pop(t *testing.T) {
	// Test popping elements
	s := NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	// Pop should return elements in reverse order
	if got := s.Pop(); got != 3 {
		t.Errorf("First pop should return 3, got %v", got)
	}
	if got := s.Pop(); got != 2 {
		t.Errorf("Second pop should return 2, got %v", got)
	}
	if got := s.Pop(); got != 1 {
		t.Errorf("Third pop should return 1, got %v", got)
	}

	// Stack should be empty after all pops
	if !s.IsEmpty() {
		t.Error("Stack should be empty after all pops")
	}

	// Test pop on empty stack (should panic)
	defer func() {
		if r := recover(); r == nil {
			t.Error("Pop on empty stack should panic")
		}
	}()
	s.Pop()
}

func TestStack_Size(t *testing.T) {
	// Test size method
	s := NewStack[int]()
	if s.Size() != 0 {
		t.Errorf("NewHeap stack should have size 0, got %d", s.Size())
	}

	s.Push(1)
	if s.Size() != 1 {
		t.Errorf("Stack with one element should have size 1, got %d", s.Size())
	}

	s.Push(2)
	if s.Size() != 2 {
		t.Errorf("Stack with two elements should have size 2, got %d", s.Size())
	}

	s.Pop()
	if s.Size() != 1 {
		t.Errorf("Stack after pop should have size 1, got %d", s.Size())
	}
}

func TestStack_IsEmpty(t *testing.T) {
	// Test IsEmpty method
	s := NewStack[int]()
	if !s.IsEmpty() {
		t.Error("NewHeap stack should be empty")
	}

	s.Push(1)
	if s.IsEmpty() {
		t.Error("Stack with one element should not be empty")
	}

	s.Pop()
	if !s.IsEmpty() {
		t.Error("Stack after pop should be empty")
	}
}

func TestStack_Peek(t *testing.T) {
	// Test Peek method
	s := NewStack[int]()

	s.Push(1)
	if s.Peek() != 1 {
		t.Errorf("Peek should return top element 1, got %v", s.Peek())
	}

	s.Push(2)
	if s.Peek() != 2 {
		t.Errorf("Peek should return top element 2, got %v", s.Peek())
	}

	s.Pop()
	if s.Peek() != 1 {
		t.Errorf("Peek after pop should return 1, got %v", s.Peek())
	}
}
