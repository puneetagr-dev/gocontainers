package gocontainers

import (
	"testing"
)

func TestSetEmpty(t *testing.T) {
	set := NewSet[int]()
	if !set.IsEmpty() {
		t.Error("New set should be empty")
	}
	if set.Size() != 0 {
		t.Errorf("New set size should be 0, got %d", set.Size())
	}
}

func TestSetAdd(t *testing.T) {
	set := NewSet[int]()
	set.Add(1)
	set.Add(2)
	set.Add(3)

	if set.Size() != 3 {
		t.Errorf("Set size should be 3 after adding 3 elements, got %d", set.Size())
	}
	for _, v := range []int{1, 2, 3} {
		if !set.Contains(v) {
			t.Errorf("Set should contain %d", v)
		}
	}
}

func TestSetRemove(t *testing.T) {
	set := NewSet[int]()
	set.Add(1)
	set.Add(2)
	set.Add(3)

	// Test removing existing element
	set.Remove(2)
	if set.Contains(2) {
		t.Error("Set should not contain removed element")
	}
	if set.Size() != 2 {
		t.Errorf("Set size should be 2 after removing one element, got %d", set.Size())
	}

	// Test removing non-existent element
	set.Remove(99)
	if set.Size() != 2 {
		t.Errorf("Removing non-existent element should not change size")
	}
}

func TestSetClear(t *testing.T) {
	set := NewSet[int]()
	set.Add(1)
	set.Add(2)
	set.Add(3)

	set.Clear()
	if !set.IsEmpty() {
		t.Error("Set should be empty after clear")
	}
	if set.Size() != 0 {
		t.Errorf("Set size should be 0 after clear, got %d", set.Size())
	}
}

func TestSetContains(t *testing.T) {
	set := NewSet[int]()
	set.Add(1)
	set.Add(2)
	set.Add(3)

	// Test existing elements
	for _, v := range []int{1, 2, 3} {
		if !set.Contains(v) {
			t.Errorf("Set should contain %d", v)
		}
	}

	// Test non-existent element
	if set.Contains(4) {
		t.Error("Set should not contain non-existent element")
	}
}

func TestSetUnion(t *testing.T) {
	tests := []struct {
		name         string
		set1Elements []int
		set2Elements []int
		expected     *Set[int]
	}{
		{
			name:         "no overlap",
			set1Elements: []int{1, 2, 3},
			set2Elements: []int{4, 5, 6},
			expected: func() *Set[int] {
				set := NewSet[int]()
				set.Add(1)
				set.Add(2)
				set.Add(3)
				set.Add(4)
				set.Add(5)
				set.Add(6)
				return set
			}(),
		},
		{
			name:         "with overlap",
			set1Elements: []int{1, 2, 3},
			set2Elements: []int{3, 4, 5},
			expected: func() *Set[int] {
				set := NewSet[int]()
				set.Add(1)
				set.Add(2)
				set.Add(3)
				set.Add(4)
				set.Add(5)
				return set
			}(),
		},
		{
			name:         "first set empty",
			set1Elements: []int{},
			set2Elements: []int{1, 2, 3},
			expected: func() *Set[int] {
				set := NewSet[int]()
				set.Add(1)
				set.Add(2)
				set.Add(3)
				return set
			}(),
		},
		{
			name:         "second set empty",
			set1Elements: []int{},
			set2Elements: []int{4, 5, 6},
			expected: func() *Set[int] {
				set := NewSet[int]()
				set.Add(4)
				set.Add(5)
				set.Add(6)
				return set
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set1 := NewSet[int]()
			for _, v := range tt.set1Elements {
				set1.Add(v)
			}
			set2 := NewSet[int]()
			for _, v := range tt.set2Elements {
				set2.Add(v)
			}

			result := set1.Union(set2)
			if !result.Equal(tt.expected) {
				t.Errorf("%s: Union = %v, want %v", tt.name, result, tt.expected)
			}
		})
	}
}

func TestSetEquality(t *testing.T) {
	tests := []struct {
		name         string
		set1Elements []int
		set2Elements []int
		expected     bool
	}{
		{
			name:         "empty sets",
			set1Elements: []int{},
			set2Elements: []int{},
			expected:     true,
		},
		{
			name:         "equal sets",
			set1Elements: []int{1, 2, 3},
			set2Elements: []int{1, 2, 3},
			expected:     true,
		},
		{
			name:         "different sizes",
			set1Elements: []int{1, 2, 3},
			set2Elements: []int{1, 2, 3, 4},
			expected:     false,
		},
		{
			name:         "same elements different order",
			set1Elements: []int{1, 2, 3},
			set2Elements: []int{3, 1, 2},
			expected:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set1 := NewSet[int]()
			for _, v := range tt.set1Elements {
				set1.Add(v)
			}
			set2 := NewSet[int]()
			for _, v := range tt.set2Elements {
				set2.Add(v)
			}

			if set1.Equal(set2) != tt.expected {
				t.Errorf("%s: Equal = %v, want %v", tt.name, set1.Equal(set2), tt.expected)
			}
		})
	}
}

func TestSetToSlice(t *testing.T) {
	tests := []struct {
		name     string
		set      *Set[int]
		add      []int
		expected []int
	}{
		{
			name:     "empty set",
			set:      NewSet[int](),
			expected: []int{},
		},
		{
			name:     "with elements",
			set:      NewSet[int](),
			add:      []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, v := range tt.add {
				tt.set.Add(v)
			}

			slice := tt.set.ToSlice()
			if len(slice) != len(tt.expected) {
				t.Errorf("%s: ToSlice length = %d, want %d", tt.name, len(slice), len(tt.expected))
			}

			elements := make(map[int]struct{})
			for _, v := range slice {
				elements[v] = struct{}{}
			}
			if len(elements) != len(tt.expected) {
				t.Errorf("%s: ToSlice contains duplicates", tt.name)
			}
			for _, v := range tt.expected {
				if _, ok := elements[v]; !ok {
					t.Errorf("%s: ToSlice missing element %d", tt.name, v)
				}
			}
		})
	}
}
