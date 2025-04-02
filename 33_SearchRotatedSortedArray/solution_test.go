package searchrotatedsortedarray

import "testing"

func TestSearch(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "target exists in rotated array",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   0,
			expected: 4,
		},
		{
			name:     "target does not exist",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   3,
			expected: -1,
		},
		{
			name:     "empty array",
			nums:     []int{},
			target:   5,
			expected: -1,
		},
		{
			name:     "single element array - found",
			nums:     []int{1},
			target:   1,
			expected: 0,
		},
		{
			name:     "single element array - not found",
			nums:     []int{1},
			target:   0,
			expected: -1,
		},
		{
			name:     "rotated array - target at start",
			nums:     []int{5, 1, 3},
			target:   5,
			expected: 0,
		},
		{
			name:     "rotated array - target in middle",
			nums:     []int{4, 5, 6, 7, 8, 1, 2, 3},
			target:   8,
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.nums, tt.target); got != tt.expected {
				t.Errorf("search() = %v, want %v", got, tt.expected)
			}
		})
	}
}
