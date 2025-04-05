package btreezigzaglevelordertraversal

import "testing"

func TestZigzagLevelOrder(t *testing.T) {
	tests := []struct {
		name     string
		input    *TreeNode
		expected [][]int
	}{
		{
			name: "normal tree",
			input: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			expected: [][]int{{3}, {20, 9}, {15, 7}},
		},
		{
			name:     "nil tree",
			input:    nil,
			expected: [][]int{},
		},
		{
			name: "single node",
			input: &TreeNode{
				Val: 1,
			},
			expected: [][]int{{1}},
		},
		{
			name: "complete tree with null nodes",
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 5,
					},
				},
			},
			expected: [][]int{{1}, {3, 2}, {4, 5}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := zigzagLevelOrder(tt.input)
			if len(result) != len(tt.expected) {
				t.Errorf("got length %v, want length %v", len(result), len(tt.expected))
				return
			}
			for i := range result {
				if len(result[i]) != len(tt.expected[i]) {
					t.Errorf("at level %v: got length %v, want length %v", i, len(result[i]), len(tt.expected[i]))
					continue
				}
				for j := range result[i] {
					if result[i][j] != tt.expected[i][j] {
						t.Errorf("at level %v index %v: got %v, want %v", i, j, result[i][j], tt.expected[i][j])
					}
				}
			}
		})
	}
}
