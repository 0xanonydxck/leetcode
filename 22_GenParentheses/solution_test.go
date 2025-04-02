package genparentheses

import "testing"

func TestGenerateParenthesis(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected []string
	}{
		{
			name:     "n=1",
			n:        1,
			expected: []string{"()"},
		},
		{
			name:     "n=2",
			n:        2,
			expected: []string{"(())", "()()"},
		},
		{
			name:     "n=3",
			n:        3,
			expected: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			name:     "n=0",
			n:        0,
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateParenthesis(tt.n)
			if len(result) != len(tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
				return
			}

			// Convert slices to maps for comparison since order doesn't matter
			resultMap := make(map[string]bool)
			expectedMap := make(map[string]bool)
			for _, v := range result {
				resultMap[v] = true
			}
			for _, v := range tt.expected {
				expectedMap[v] = true
			}

			for k := range resultMap {
				if !expectedMap[k] {
					t.Errorf("got unexpected value %v", k)
				}
			}
			for k := range expectedMap {
				if !resultMap[k] {
					t.Errorf("missing expected value %v", k)
				}
			}
		})
	}
}
