package subset

import "testing"

func Test_subsets(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{
			name: "empty array",
			nums: []int{},
			expected: [][]int{
				{},
			},
		},
		{
			name: "single element",
			nums: []int{1},
			expected: [][]int{
				{},
				{1},
			},
		},
		{
			name: "three elements",
			nums: []int{1, 2, 3},
			expected: [][]int{
				{},
				{1},
				{2},
				{3},
				{1, 2},
				{1, 3},
				{2, 3},
				{1, 2, 3},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := subsets(tt.nums)
			if !equalIntSlices(got, tt.expected) {
				t.Errorf("subsets() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// Helper function to compare two int slice slices regardless of order
func equalIntSlices(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	used := make([]bool, len(b))
	for _, sliceA := range a {
		found := false
		for j, sliceB := range b {
			if !used[j] && equalSlice(sliceA, sliceB) {
				used[j] = true
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// Helper function to compare two int slices
func equalSlice(a, b []int) bool {
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
