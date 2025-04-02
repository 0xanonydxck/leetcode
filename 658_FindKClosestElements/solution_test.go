package findkclosestelements

import "testing"

func TestFindClosestElements(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		k        int
		x        int
		expected []int
	}{
		{
			name:     "basic case",
			arr:      []int{1, 2, 3, 4, 5},
			k:        4,
			x:        3,
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "x outside array on left",
			arr:      []int{1, 2, 3, 4, 5},
			k:        3,
			x:        -1,
			expected: []int{1, 2, 3},
		},
		{
			name:     "x outside array on right",
			arr:      []int{1, 2, 3, 4, 5},
			k:        3,
			x:        10,
			expected: []int{3, 4, 5},
		},
		{
			name:     "empty array",
			arr:      []int{},
			k:        3,
			x:        5,
			expected: []int{},
		},
		{
			name:     "k equals array length",
			arr:      []int{1, 2, 3},
			k:        3,
			x:        2,
			expected: []int{1, 2, 3},
		},
		{
			name:     "x equals element in array",
			arr:      []int{1, 2, 3, 4, 5},
			k:        1,
			x:        3,
			expected: []int{3},
		},
		{
			name:     "x between two distinct values",
			arr:      []int{1, 1, 1, 10, 10, 10},
			k:        1,
			x:        9,
			expected: []int{10},
		},
		{
			name:     "x between two distinct values",
			arr:      []int{1, 2, 2, 3, 5, 5, 6, 8},
			k:        4,
			x:        4,
			expected: []int{2, 3, 5, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findClosestElements(tt.arr, tt.k, tt.x)
			if len(got) != len(tt.expected) {
				t.Errorf("findClosestElements() got length = %v, want length %v", len(got), len(tt.expected))
				return
			}
			for i := range got {
				if got[i] != tt.expected[i] {
					t.Errorf("findClosestElements() = %v, want %v", got, tt.expected)
					return
				}
			}
		})
	}
}
