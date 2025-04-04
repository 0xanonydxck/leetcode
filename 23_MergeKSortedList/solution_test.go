package mergeksortedlist

import (
	"testing"
)

func TestMergeKLists(t *testing.T) {
	tests := []struct {
		name     string
		lists    []*ListNode
		expected *ListNode
	}{
		{
			name:     "empty lists",
			lists:    []*ListNode{},
			expected: nil,
		},
		{
			name: "single list",
			lists: []*ListNode{
				{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}},
			},
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}},
		},
		{
			name: "multiple sorted lists",
			lists: []*ListNode{
				{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}},
				{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
				{Val: 2, Next: &ListNode{Val: 6}},
			},
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}}}}}}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := mergeKLists(tt.lists)

			// Compare result with expected
			current := result
			expected := tt.expected
			for current != nil && expected != nil {
				if current.Val != expected.Val {
					t.Errorf("got %v, want %v", current.Val, expected.Val)
				}
				current = current.Next
				expected = expected.Next
			}

			if current != nil || expected != nil {
				t.Errorf("lists have different lengths")
			}
		})
	}
}
