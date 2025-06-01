package gocontainers

import (
	"testing"
)

func TestQueue_Enqueue(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	if q.elements[0] != 1 {
		t.Error("Front element should be 1")
	}
	if q.elements[1] != 2 {
		t.Error("Middle element should be 2")
	}
	if q.elements[2] != 3 {
		t.Error("Back element should be 3")
	}
	if q.Size() != 3 {
		t.Errorf("Size should be 3, got %d", q.Size())
	}
}

func TestQueue_Dequeue(t *testing.T) {
	q := NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	if q.Dequeue() != 1 {
		t.Error("Front element should be 1")
	}
	if q.Dequeue() != 2 {
		t.Error("Middle element should be 2")
	}
	if q.Dequeue() != 3 {
		t.Error("Back element should be 3")
	}
}

func TestQueue_Size(t *testing.T) {
	tests := []struct {
		name  string
		setup func(q *Queue[int])
		want  int
	}{
		{
			name:  "empty queue",
			setup: func(q *Queue[int]) {},
			want:  0,
		},
		{
			name: "one element",
			setup: func(q *Queue[int]) {
				q.Enqueue(42)
			},
			want: 1,
		},
		{
			name: "multiple elements",
			setup: func(q *Queue[int]) {
				q.Enqueue(1)
				q.Enqueue(2)
				q.Enqueue(3)
			},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := NewQueue[int]()
			tt.setup(q)
			if got := q.Size(); got != tt.want {
				t.Errorf("%s: got %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestQueue_IsEmpty(t *testing.T) {
	tests := []struct {
		name  string
		setup func(q *Queue[int])
		want  bool
	}{
		{
			name:  "empty queue",
			setup: func(q *Queue[int]) {},
			want:  true,
		},
		{
			name: "non-empty queue",
			setup: func(q *Queue[int]) {
				q.Enqueue(42)
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := NewQueue[int]()
			tt.setup(q)
			if got := q.IsEmpty(); got != tt.want {
				t.Errorf("%s: got %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestQueue_Clear(t *testing.T) {
	tests := []struct {
		name  string
		setup func(q *Queue[int])
	}{
		{
			name:  "clear empty queue",
			setup: func(q *Queue[int]) {},
		},
		{
			name: "clear non-empty queue",
			setup: func(q *Queue[int]) {
				q.Enqueue(42)
				q.Enqueue(1)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := NewQueue[int]()
			tt.setup(q)
			q.Clear()
			if !q.IsEmpty() {
				t.Errorf("%s: queue should be empty after Clear()", tt.name)
			}
			if len(q.elements) != 0 {
				t.Errorf("%s: queue should be empty after Clear()", tt.name)
			}
		})
	}
}
