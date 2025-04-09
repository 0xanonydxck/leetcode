package triangle

import "testing"

func Test_minimumTotal(t *testing.T) {
	tests := []struct {
		name     string
		triangle [][]int
		want     int
	}{
		{
			name: "example 1",
			triangle: [][]int{
				{2},
				{3, 4},
				{6, 5, 7},
				{4, 1, 8, 3},
			},
			want: 11,
		},
		{
			name: "example 2",
			triangle: [][]int{
				{-1},
				{2, 3},
				{1, -1, 3},
				{4, 1, 8, 3},
				{1, 2, 2, 4, 5},
			},
			want: -1,
		},
		{
			name: "single element",
			triangle: [][]int{
				{-10},
			},
			want: -10,
		},
		{
			name: "two rows",
			triangle: [][]int{
				{1},
				{2, 3},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumTotal(tt.triangle); got != tt.want {
				t.Errorf("minimumTotal() = %v, want %v", got, tt.want)
			}
		})
	}
}
