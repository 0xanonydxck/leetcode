package uniquepath

import "testing"

func Test_uniquePaths(t *testing.T) {
	tests := []struct {
		name string
		m    int
		n    int
		want int
	}{
		{
			name: "3x2 grid",
			m:    3,
			n:    2,
			want: 3,
		},
		{
			name: "3x7 grid",
			m:    3,
			n:    7,
			want: 28,
		},
		{
			name: "7x3 grid",
			m:    7,
			n:    3,
			want: 28,
		},
		{
			name: "3x3 grid",
			m:    3,
			n:    3,
			want: 6,
		},
		{
			name: "1x1 grid",
			m:    1,
			n:    1,
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePaths(tt.m, tt.n); got != tt.want {
				t.Errorf("uniquePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
