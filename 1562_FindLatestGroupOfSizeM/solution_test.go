package findlatestgroupofsizem

import "testing"

func Test_findLatestStep(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		m    int
		want int
	}{
		{
			name: "example 1",
			arr:  []int{3, 5, 1, 2, 4},
			m:    1,
			want: 4,
		},
		{
			name: "example 2",
			arr:  []int{3, 1, 5, 4, 2},
			m:    2,
			want: -1,
		},
		{
			name: "example 3",
			arr:  []int{1},
			m:    1,
			want: 1,
		},
		{
			name: "example 4",
			arr:  []int{2, 1},
			m:    2,
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLatestStep(tt.arr, tt.m); got != tt.want {
				t.Errorf("findLatestStep() = %v, want %v", got, tt.want)
			}
		})
	}
}
