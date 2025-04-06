package circularpermutationbinaryrepresentation

import "testing"

func TestCircularPermutation(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		start    int
		expected []int
	}{
		{
			name:     "n=2 start=3",
			n:        2,
			start:    3,
			expected: []int{3, 2, 0, 1},
		},
		{
			name:     "n=3 start=2",
			n:        3,
			start:    2,
			expected: []int{2, 3, 1, 0, 4, 5, 7, 6},
		},
		{
			name:     "n=1 start=0",
			n:        1,
			start:    0,
			expected: []int{0, 1},
		},
		{
			name:     "n=2 start=0",
			n:        2,
			start:    0,
			expected: []int{0, 1, 3, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := circularPermutation(tt.n, tt.start)
			if !equalIntSlice(got, tt.expected) {
				t.Errorf("circularPermutation() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// Helper function to compare two int slices
func equalIntSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
