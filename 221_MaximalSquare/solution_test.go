package maximalsquare

import "testing"

func Test_maximalSquare(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]byte
		want   int
	}{
		{
			name:   "empty matrix",
			matrix: [][]byte{},
			want:   0,
		},
		{
			name: "single row matrix",
			matrix: [][]byte{
				{'1', '0', '1', '0', '0'},
			},
			want: 1,
		},
		{
			name: "single column matrix",
			matrix: [][]byte{
				{'1'},
				{'0'},
				{'1'},
			},
			want: 1,
		},
		{
			name: "2x2 square",
			matrix: [][]byte{
				{'1', '1'},
				{'1', '1'},
			},
			want: 4,
		},
		{
			name: "3x3 matrix with 2x2 square",
			matrix: [][]byte{
				{'1', '1', '0'},
				{'1', '1', '1'},
				{'1', '1', '1'},
			},
			want: 4,
		},
		{
			name: "larger matrix",
			matrix: [][]byte{
				{'1', '0', '1', '0', '0'},
				{'1', '0', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '0', '0', '1', '0'},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximalSquare(tt.matrix); got != tt.want {
				t.Errorf("maximalSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
