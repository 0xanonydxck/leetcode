package stepbystepdirectionbtree

import "testing"

func TestGetDirections(t *testing.T) {
	tests := []struct {
		name       string
		root       *TreeNode
		startValue int
		destValue  int
		want       string
	}{
		{
			name: "Example 1",
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
			},
			startValue: 3,
			destValue:  6,
			want:       "UURL",
		},
		{
			name: "Example 2",
			root: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
			},
			startValue: 2,
			destValue:  1,
			want:       "L",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDirections(tt.root, tt.startValue, tt.destValue); got != tt.want {
				t.Errorf("getDirections() = %v, want %v", got, tt.want)
			}
		})
	}
}
