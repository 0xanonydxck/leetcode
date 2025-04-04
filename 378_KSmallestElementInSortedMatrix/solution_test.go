package ksmallestelementinsortedmatrix

import (
	"testing"
)

func TestKthSmallest(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]int
		k        int
		expected int
	}{
		{
			name:     "1x1 matrix",
			matrix:   [][]int{{1}},
			k:        1,
			expected: 1,
		},
		{
			name: "3x3 matrix",
			matrix: [][]int{
				{1, 5, 9},
				{10, 11, 13},
				{12, 13, 15},
			},
			k:        8,
			expected: 13,
		},
		{
			name: "2x2 matrix",
			matrix: [][]int{
				{1, 2},
				{3, 4},
			},
			k:        2,
			expected: 2,
		},
		{
			name: "matrix with duplicates",
			matrix: [][]int{
				{1, 2, 2},
				{2, 3, 3},
				{3, 3, 4},
			},
			k:        4,
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := kthSmallest(tt.matrix, tt.k)
			if result != tt.expected {
				t.Errorf("kthSmallest() = %v, want %v", result, tt.expected)
			}
		})
	}
}
