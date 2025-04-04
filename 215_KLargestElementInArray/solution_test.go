package klargestelementinarray

import (
	"testing"
)

func TestFindKthLargest(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{
			name:     "single element",
			nums:     []int{1},
			k:        1,
			expected: 1,
		},
		{
			name:     "sorted array",
			nums:     []int{1, 2, 3, 4, 5},
			k:        2,
			expected: 4,
		},
		{
			name:     "unsorted array",
			nums:     []int{3, 2, 1, 5, 6, 4},
			k:        2,
			expected: 5,
		},
		{
			name:     "array with duplicates",
			nums:     []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:        4,
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findKthLargest(tt.nums, tt.k)
			if result != tt.expected {
				t.Errorf("findKthLargest() = %v, want %v", result, tt.expected)
			}
		})
	}
}
