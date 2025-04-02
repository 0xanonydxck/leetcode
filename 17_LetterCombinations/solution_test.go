package lettercombinations

import "testing"

func TestLetterCombinations(t *testing.T) {
	tests := []struct {
		name     string
		digits   string
		expected []string
	}{
		{
			name:     "empty string",
			digits:   "",
			expected: []string{},
		},
		{
			name:     "single digit",
			digits:   "2",
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "two digits",
			digits:   "23",
			expected: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			name:     "three digits",
			digits:   "234",
			expected: []string{"adg", "adh", "adi", "aeg", "aeh", "aei", "afg", "afh", "afi", "bdg", "bdh", "bdi", "beg", "beh", "bei", "bfg", "bfh", "bfi", "cdg", "cdh", "cdi", "ceg", "ceh", "cei", "cfg", "cfh", "cfi"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := letterCombinations(tt.digits)
			if len(got) != len(tt.expected) {
				t.Errorf("letterCombinations() = %v, want %v", got, tt.expected)
				return
			}
			// Check if all elements are present (order doesn't matter)
			gotMap := make(map[string]bool)
			for _, s := range got {
				gotMap[s] = true
			}
			for _, s := range tt.expected {
				if !gotMap[s] {
					t.Errorf("letterCombinations() = %v, want %v", got, tt.expected)
					return
				}
			}
		})
	}
}
