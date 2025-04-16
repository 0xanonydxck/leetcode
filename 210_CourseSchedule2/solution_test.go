package courseschedule2

import "testing"

func TestFindOrder(t *testing.T) {
	tests := []struct {
		name          string
		numCourses    int
		prerequisites [][]int
		want          []int
	}{
		{
			name:          "Example 1",
			numCourses:    2,
			prerequisites: [][]int{{1, 0}},
			want:          []int{0, 1},
		},
		{
			name:          "Example 2",
			numCourses:    4,
			prerequisites: [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
			want:          []int{0, 1, 2, 3},
		},
		{
			name:          "No prerequisites",
			numCourses:    3,
			prerequisites: [][]int{},
			want:          []int{0, 1, 2},
		},
		{
			name:          "Cycle case",
			numCourses:    3,
			prerequisites: [][]int{{1, 0}, {1, 2}, {0, 1}},
			want:          []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findOrder(tt.numCourses, tt.prerequisites)
			if !equalSlices(got, tt.want) {
				t.Errorf("findOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func equalSlices(a, b []int) bool {
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
