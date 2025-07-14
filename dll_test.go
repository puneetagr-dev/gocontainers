package gocontainers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestDLLInt tests DLL operations with integer values
type intTest struct {
	name    string
	setup   func(dll *DLL[int])
	check   func(t *testing.T, dll *DLL[int])
	cleanup func(dll *DLL[int])
}

func TestDLLInt(t *testing.T) {
	tests := []intTest{
		{
			name: "AddFront",
			setup: func(dll *DLL[int]) {
				dll.AddFront(NewNode(1))
				dll.AddFront(NewNode(2))
				dll.AddFront(NewNode(3))
			},
			check: func(t *testing.T, dll *DLL[int]) {
				if dll.GetFront().Get() != 3 {
					t.Error("GetFront should return 3")
				}
				if dll.GetBack().Get() != 1 {
					t.Error("GetBack should return 1")
				}
				if dll.Size() != 3 {
					t.Errorf("Size should be 3, got %d", dll.Size())
				}
			},
			cleanup: func(dll *DLL[int]) {
				dll.Clear()
			},
		},
		{
			name: "RemoveFront",
			setup: func(dll *DLL[int]) {
				dll.AddFront(NewNode(1))
				dll.AddFront(NewNode(2))
				dll.AddFront(NewNode(3))
				dll.RemoveFront()
			},
			check: func(t *testing.T, dll *DLL[int]) {
				if dll.GetFront().Get() != 2 {
					t.Error("After RemoveFront, front should be 2")
				}
				if dll.Size() != 2 {
					t.Errorf("Size should be 2, got %d", dll.Size())
				}
			},
			cleanup: func(dll *DLL[int]) {
				dll.Clear()
			},
		},
		{
			name: "RemoveBack",
			setup: func(dll *DLL[int]) {
				dll.AddFront(NewNode(1))
				dll.AddFront(NewNode(2))
				dll.AddFront(NewNode(3))
				dll.RemoveBack()
			},
			check: func(t *testing.T, dll *DLL[int]) {
				if dll.GetBack().Get() != 2 {
					t.Error("After RemoveBack, back should be 2")
				}
				if dll.Size() != 2 {
					t.Errorf("Size should be 2, got %d", dll.Size())
				}
			},
			cleanup: func(dll *DLL[int]) {
				dll.Clear()
			},
		},
		{
			name: "DeleteMatch",
			setup: func(dll *DLL[int]) {
				dll.AddFront(NewNode(1))
				dll.AddFront(NewNode(2))
				dll.AddFront(NewNode(3))
				dll.DeleteMatch(2)
			},
			check: func(t *testing.T, dll *DLL[int]) {
				if dll.Size() != 2 {
					t.Errorf("Size should be 2 after deletion, got %d", dll.Size())
				}
				if dll.GetFront().Get() != 3 {
					t.Error("Front should be 3 after deleting 2")
				}
				if dll.GetBack().Get() != 1 {
					t.Error("Back should be 1 after deleting 2")
				}
			},
			cleanup: func(dll *DLL[int]) {
				dll.Clear()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dll := NewDLL[int]()
			defer tt.cleanup(dll)
			tt.setup(dll)
			tt.check(t, dll)
		})
	}
}

// TestDLLString tests DLL operations with string values
type stringTest struct {
	name    string
	setup   func(dll *DLL[string])
	check   func(t *testing.T, dll *DLL[string])
	cleanup func(dll *DLL[string])
}

func TestDLLString(t *testing.T) {
	tests := []stringTest{
		{
			name: "AddFront",
			setup: func(dll *DLL[string]) {
				dll.AddFront(NewNode("hello"))
				dll.AddFront(NewNode("world"))
			},
			check: func(t *testing.T, dll *DLL[string]) {
				if dll.GetFront().Get() != "world" {
					t.Error("GetFront should return world")
				}
				if dll.GetBack().Get() != "hello" {
					t.Error("GetBack should return hello")
				}
				if dll.Size() != 2 {
					t.Errorf("Size should be 2, got %d", dll.Size())
				}
			},
			cleanup: func(dll *DLL[string]) {
				dll.Clear()
			},
		},
		{
			name: "DeleteMatch",
			setup: func(dll *DLL[string]) {
				dll.AddFront(NewNode("hello"))
				dll.AddFront(NewNode("world"))
				dll.DeleteMatch("hello")
			},
			check: func(t *testing.T, dll *DLL[string]) {
				if dll.Size() != 1 {
					t.Errorf("Size should be 1 after deletion, got %d", dll.Size())
				}
				if dll.GetFront().Get() != "world" {
					t.Error("Front should be world after deleting hello")
				}
			},
			cleanup: func(dll *DLL[string]) {
				dll.Clear()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dll := NewDLL[string]()
			defer tt.cleanup(dll)
			tt.setup(dll)
			tt.check(t, dll)
		})
	}
}

func TestDeleteByNode(t *testing.T) {
	// single node list
	dll := NewDLL[string]()
	dll.AddBack(NewNode("hello"))

	assert.Equal(t, 1, dll.Size())
	node := dll.GetBack()
	dll.DeleteNode(node)
	assert.Equal(t, 0, dll.Size())
	assert.Nil(t, dll.head)
	assert.Nil(t, dll.tail)

	// two nodes, delete head
	dll = NewDLL[string]()
	dll.AddFront(NewNode("hello"))
	dll.AddFront(NewNode("how"))
	assert.Equal(t, 2, dll.Size())
	node = dll.GetFront()
	dll.DeleteNode(node)
	assert.Equal(t, 1, dll.Size())
	assert.Equal(t, dll.tail, dll.head)
	assert.Equal(t, "hello", dll.head.Get())
	assert.Equal(t, "hello", dll.tail.Get())

	// two nodes, delete tail
	dll = NewDLL[string]()
	dll.AddFront(NewNode("hello"))
	dll.AddFront(NewNode("how"))
	assert.Equal(t, 2, dll.Size())
	node = dll.GetBack()
	dll.DeleteNode(node)
	assert.Equal(t, 1, dll.Size())
	assert.Equal(t, dll.tail, dll.head)
	assert.Equal(t, "how", dll.head.Get())
	assert.Equal(t, "how", dll.tail.Get())

	// more than two nodes
	dll = NewDLL[string]()
	n1 := NewNode("hello")
	n2 := NewNode("how")
	n3 := NewNode("are")
	dll.AddFront(n1)
	dll.AddFront(n2)
	dll.AddFront(n3)
	assert.Equal(t, 3, dll.Size())
	dll.DeleteNode(n2)
	assert.Equal(t, 2, dll.Size())
	assert.Equal(t, dll.tail.prev, dll.head)
	assert.Equal(t, dll.head.next, dll.tail)
	assert.Nil(t, dll.head.prev)
	assert.Nil(t, dll.tail.next)
	assert.Equal(t, "are", dll.head.Get())
	assert.Equal(t, "hello", dll.tail.Get())
}
