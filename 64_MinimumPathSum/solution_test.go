package minimumpathsum

import "testing"

func Test_minPathSum(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "3x3 grid",
			grid: [][]int{
				{1, 3, 1},
				{1, 5, 1},
				{4, 2, 1},
			},
			want: 7,
		},
		{
			name: "2x3 grid",
			grid: [][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
			want: 12,
		},
		{
			name: "1x1 grid",
			grid: [][]int{
				{1},
			},
			want: 1,
		},
		{
			name: "2x2 grid",
			grid: [][]int{
				{1, 2},
				{1, 1},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minPathSum(tt.grid); got != tt.want {
				t.Errorf("minPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
