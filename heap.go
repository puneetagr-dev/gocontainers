package gocontainers

import (
	"container/heap"
	"fmt"
)

type Item[T any] struct {
	val   T
	index int // index in the heap slice
}

func NewItem[T any](val T) *Item[T] {
	return &Item[T]{val: val, index: -1}
}

func (i *Item[T]) Update(val T) {
	i.val = val
}

func (i *Item[T]) Get() T {
	return i.val
}

// Heap is a generic heap with elements of type T.
// The comparator defines the ordering: comparator(a, b) == true means element a has higher priority than element b.
type Heap[T any] struct {
	data       []*Item[T]
	comparator func(a, b T) bool
}

// NewHeap creates a new Heap with the given comparator function.
// The comparator should return true if element a has higher priority than element b.
func NewHeap[T any](comparator func(a, b T) bool) *Heap[T] {
	return &Heap[T]{
		comparator: comparator,
		data:       []*Item[T]{},
	}
}

// Len returns the number of items in the heap.
func (h *Heap[T]) Len() int {
	return len(h.data)
}

// Less compares two elements by their priority using the comparator.
func (h *Heap[T]) Less(i, j int) bool {
	return h.comparator(h.data[i].val, h.data[j].val)
}

// Swap swaps two elements and updates their indices.
func (h *Heap[T]) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
	h.data[i].index = i
	h.data[j].index = j
}

// PushItem adds x as the last element in the heap.
func (h *Heap[T]) PushItem(item *Item[T]) {
	heap.Push(h, item)
}

func (h *Heap[T]) Push(x any) {
	n := len(h.data)
	item := x.(*Item[T])
	item.index = n
	h.data = append(h.data, item)
}

// PopItem removes and returns the last element of the heap.
func (h *Heap[T]) PopItem() *Item[T] {
	x := heap.Pop(h)
	return x.(*Item[T])
}

func (h *Heap[T]) Pop() any {
	old := h.data
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	h.data = old[0 : n-1]
	return item
}

// Update updates the value of an existing item and fixes the heap order.
// Caller must ensure the updated value respects the heap ordering rules.
// the item passed should be with the updated value
func (h *Heap[T]) Update(item *Item[T]) {
	heap.Fix(h, item.index)
}

// Remove removes an item from the heap.
func (h *Heap[T]) Remove(item *Item[T]) {
	heap.Remove(h, item.index)
}

// Init initializes the heap (useful if you have pre-populated data).
func (h *Heap[T]) Init() {
	heap.Init(h)
}

// Peek returns the highest priority value without removing it.
// Returns false if the heap is empty.
func (h *Heap[T]) Peek() (*Item[T], bool) {
	if h.Len() == 0 {
		return nil, false
	}
	return h.data[0], true
}

// String returns a string representation of the heap (for debugging).
func (h *Heap[T]) String() string {
	s := "Heap: ["
	for i, item := range h.data {
		s += fmt.Sprintf("%v", item.val)
		if i < len(h.data)-1 {
			s += ", "
		}
	}
	s += "]"
	return s
}
