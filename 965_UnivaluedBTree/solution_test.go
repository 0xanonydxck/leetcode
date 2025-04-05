package univaluedbtree

import "testing"

func TestIsUnivalTree(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected bool
	}{
		{
			name:     "Single node tree",
			root:     &TreeNode{Val: 1},
			expected: true,
		},
		{
			name: "All nodes same value",
			root: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 1},
			},
			expected: true,
		},
		{
			name: "Different values",
			root: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 2},
			},
			expected: false,
		},
		{
			name:     "Nil root",
			root:     nil,
			expected: true,
		},
		{
			name: "Larger unival tree",
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:   5,
					Left:  &TreeNode{Val: 5},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val:  5,
					Left: &TreeNode{Val: 5},
				},
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isUnivalTree(tt.root); got != tt.expected {
				t.Errorf("isUnivalTree() = %v, want %v", got, tt.expected)
			}
		})
	}
}
