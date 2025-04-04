package kclosestpointstoorigin

import "testing"

func TestKClosest(t *testing.T) {
	tests := []struct {
		name   string
		points [][]int
		k      int
		want   [][]int
	}{
		{
			name:   "example 1",
			points: [][]int{{1, 3}, {-2, 2}},
			k:      1,
			want:   [][]int{{-2, 2}},
		},
		{
			name:   "example 2",
			points: [][]int{{3, 3}, {5, -1}, {-2, 4}},
			k:      2,
			want:   [][]int{{3, 3}, {-2, 4}},
		},
		{
			name:   "single point",
			points: [][]int{{1, 1}},
			k:      1,
			want:   [][]int{{1, 1}},
		},
		{
			name:   "all points",
			points: [][]int{{1, 1}, {2, 2}, {3, 3}},
			k:      3,
			want:   [][]int{{1, 1}, {2, 2}, {3, 3}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := kClosest(tt.points, tt.k)
			if len(got) != len(tt.want) {
				t.Errorf("kClosest() returned %v points, want %v points", len(got), len(tt.want))
			}
			// Check each point is in the result
			for _, wantPoint := range tt.want {
				found := false
				for _, gotPoint := range got {
					if gotPoint[0] == wantPoint[0] && gotPoint[1] == wantPoint[1] {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("kClosest() = %v, want %v", got, tt.want)
					break
				}
			}
		})
	}
}
