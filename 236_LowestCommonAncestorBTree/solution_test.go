package lowestcommonancestorbtree

import "testing"

func TestLowestCommonAncestor(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		p        *TreeNode
		q        *TreeNode
		expected *TreeNode
	}{
		{
			name: "Example 1",
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 7,
						},
						Right: &TreeNode{
							Val: 4,
						},
					},
				},
				Right: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 0,
					},
					Right: &TreeNode{
						Val: 8,
					},
				},
			},
			p:        &TreeNode{Val: 5},
			q:        &TreeNode{Val: 1},
			expected: &TreeNode{Val: 3},
		},
		{
			name: "Example 2",
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 7,
						},
						Right: &TreeNode{
							Val: 4,
						},
					},
				},
				Right: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 0,
					},
					Right: &TreeNode{
						Val: 8,
					},
				},
			},
			p:        &TreeNode{Val: 5},
			q:        &TreeNode{Val: 4},
			expected: &TreeNode{Val: 5},
		},
		{
			name: "Single node tree",
			root: &TreeNode{
				Val: 1,
			},
			p:        &TreeNode{Val: 1},
			q:        &TreeNode{Val: 1},
			expected: &TreeNode{Val: 1},
		},
		{
			name: "Example 3",
			root: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 0,
					},
					Right: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 3,
						},
						Right: &TreeNode{
							Val: 5,
						},
					},
				},
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: 7,
					},
					Right: &TreeNode{
						Val: 10,
						Left: &TreeNode{
							Val: 9,
						},
						Right: &TreeNode{
							Val: 11,
						},
					},
				},
			},
			p:        &TreeNode{Val: 9},
			q:        &TreeNode{Val: 11},
			expected: &TreeNode{Val: 10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lowestCommonAncestor(tt.root, tt.p, tt.q)
			if got == nil || got.Val != tt.expected.Val {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.expected)
			}
		})
	}
}
