package findpathexistingraph

import "testing"

func TestValidPath(t *testing.T) {
	testCases := []struct {
		name        string
		n           int
		edges       [][]int
		source      int
		destination int
		want        bool
	}{
		{
			name:        "example 1",
			n:           3,
			edges:       [][]int{{0, 1}, {1, 2}, {2, 0}},
			source:      0,
			destination: 2,
			want:        true,
		},
		{
			name:        "example 2",
			n:           6,
			edges:       [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}},
			source:      0,
			destination: 5,
			want:        false,
		},
		{
			name:        "single node",
			n:           1,
			edges:       [][]int{},
			source:      0,
			destination: 0,
			want:        true,
		},
		{
			name:        "linear path",
			n:           4,
			edges:       [][]int{{0, 1}, {1, 2}, {2, 3}},
			source:      0,
			destination: 3,
			want:        true,
		},
		{
			name:        "disconnected components",
			n:           5,
			edges:       [][]int{{0, 1}, {2, 3}, {3, 4}},
			source:      0,
			destination: 4,
			want:        false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := validPath(tc.n, tc.edges, tc.source, tc.destination)
			if got != tc.want {
				t.Errorf("validPath(%v, %v, %v, %v) = %v, want %v",
					tc.n, tc.edges, tc.source, tc.destination, got, tc.want)
			}
		})
	}

}
