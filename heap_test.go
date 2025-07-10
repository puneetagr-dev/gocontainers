package gocontainers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHeapBasicOperations(t *testing.T) {
	cmp := func(a, b int) bool { return a > b } // max-heap comparator
	h := NewHeap(cmp)

	// Create items
	item1 := NewItem[int](10)
	item2 := NewItem[int](5)
	item3 := NewItem[int](20)

	// Push items
	h.PushItem(item1)
	h.PushItem(item2)
	h.PushItem(item3)

	assert.Equal(t, 3, h.Len())

	// Peek the highest priority
	val, ok := h.Peek()
	assert.True(t, ok)
	assert.Equal(t, 20, val.Get())

	// Pop the highest priority
	popped := h.PopItem()
	assert.Equal(t, 20, popped.Get())
	assert.Equal(t, 2, h.Len())

	// Pop next
	popped = h.PopItem()
	assert.Equal(t, 10, popped.Get())
	assert.Equal(t, 1, h.Len())

	// Pop last
	popped = h.PopItem()
	assert.Equal(t, 5, popped.Get())
	assert.Equal(t, 0, h.Len())

	// Pop empty
	assert.Equal(t, 0, h.Len())
}

func TestHeapUpdate(t *testing.T) {
	cmp := func(a, b int) bool { return a > b } // max-heap comparator
	h := NewHeap(cmp)

	item1 := NewItem[int](10)
	item2 := NewItem[int](5)
	item3 := NewItem[int](20)

	h.PushItem(item1)
	h.PushItem(item2)
	h.PushItem(item3)

	// Update item2 from 5 to 25 (should become the new max)
	item2.Update(25)
	h.Update(item2)

	val, ok := h.Peek()
	assert.True(t, ok)
	assert.Equal(t, 25, val.Get())

	// Pop and check order
	popped := h.PopItem()
	assert.Equal(t, 25, popped.Get())

	popped = h.PopItem()
	assert.Equal(t, 20, popped.Get())

	popped = h.PopItem()
	assert.Equal(t, 10, popped.Get())
}

func TestHeapRemove(t *testing.T) {
	cmp := func(a, b int) bool { return a > b } // max-heap comparator
	h := NewHeap(cmp)

	item1 := NewItem[int](10)
	item2 := NewItem[int](5)
	item3 := NewItem[int](20)

	h.PushItem(item1)
	h.PushItem(item2)
	h.PushItem(item3)

	// Remove item3 (20)
	h.Remove(item3)
	assert.Equal(t, 2, h.Len())

	val, ok := h.Peek()
	assert.True(t, ok)
	assert.Equal(t, 10, val.Get())

	// Remove item1 (10)
	h.Remove(item1)
	assert.Equal(t, 1, h.Len())

	val, ok = h.Peek()
	assert.True(t, ok)
	assert.Equal(t, 5, val.Get())

	// Remove item2 (5)
	h.Remove(item2)
	assert.Equal(t, 0, h.Len())

	// Peek empty
	_, ok = h.Peek()
	assert.False(t, ok)
}

func TestHeapEmptyPeekPop(t *testing.T) {
	cmp := func(a, b int) bool { return a > b }
	h := NewHeap(cmp)

	val, ok := h.Peek()
	assert.False(t, ok)
	assert.Nil(t, val)
	assert.Equal(t, 0, h.Len())
}
