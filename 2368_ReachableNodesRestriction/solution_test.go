package reachablenodesrestriction

import "testing"

func TestReachableNodes(t *testing.T) {
	testCases := []struct {
		name       string
		n          int
		edges      [][]int
		restricted []int
		want       int
	}{
		{
			name:       "example 1",
			n:          7,
			edges:      [][]int{{0, 1}, {1, 2}, {3, 1}, {4, 0}, {0, 5}, {5, 6}},
			restricted: []int{4, 5},
			want:       4,
		},
		{
			name:       "single node",
			n:          1,
			edges:      [][]int{},
			restricted: []int{},
			want:       1,
		},
		{
			name:       "all nodes restricted except root",
			n:          3,
			edges:      [][]int{{0, 1}, {0, 2}},
			restricted: []int{1, 2},
			want:       1,
		},
		{
			name:       "linear path",
			n:          4,
			edges:      [][]int{{0, 1}, {1, 2}, {2, 3}},
			restricted: []int{2},
			want:       2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := reachableNodes(tc.n, tc.edges, tc.restricted)
			if got != tc.want {
				t.Errorf("reachableNodes(%v, %v, %v) = %v, want %v",
					tc.n, tc.edges, tc.restricted, got, tc.want)
			}
		})
	}
}
