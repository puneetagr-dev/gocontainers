package gocontainers

import (
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
				dll.AddFront(1)
				dll.AddFront(2)
				dll.AddFront(3)
			},
			check: func(t *testing.T, dll *DLL[int]) {
				if dll.GetFront() != 3 {
					t.Error("GetFront should return 3")
				}
				if dll.GetBack() != 1 {
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
				dll.AddFront(1)
				dll.AddFront(2)
				dll.AddFront(3)
				dll.RemoveFront()
			},
			check: func(t *testing.T, dll *DLL[int]) {
				if dll.GetFront() != 2 {
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
				dll.AddFront(1)
				dll.AddFront(2)
				dll.AddFront(3)
				dll.RemoveBack()
			},
			check: func(t *testing.T, dll *DLL[int]) {
				if dll.GetBack() != 2 {
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
				dll.AddFront(1)
				dll.AddFront(2)
				dll.AddFront(3)
				dll.DeleteMatch(2)
			},
			check: func(t *testing.T, dll *DLL[int]) {
				if dll.Size() != 2 {
					t.Errorf("Size should be 2 after deletion, got %d", dll.Size())
				}
				if dll.GetFront() != 3 {
					t.Error("Front should be 3 after deleting 2")
				}
				if dll.GetBack() != 1 {
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
				dll.AddFront("hello")
				dll.AddFront("world")
			},
			check: func(t *testing.T, dll *DLL[string]) {
				if dll.GetFront() != "world" {
					t.Error("GetFront should return world")
				}
				if dll.GetBack() != "hello" {
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
				dll.AddFront("hello")
				dll.AddFront("world")
				dll.DeleteMatch("hello")
			},
			check: func(t *testing.T, dll *DLL[string]) {
				if dll.Size() != 1 {
					t.Errorf("Size should be 1 after deletion, got %d", dll.Size())
				}
				if dll.GetFront() != "world" {
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
