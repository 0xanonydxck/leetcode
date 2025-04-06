package palindromepartitioning

import "testing"

func TestPartition(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected [][]string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: [][]string{},
		},
		{
			name:     "single character",
			input:    "a",
			expected: [][]string{{"a"}},
		},
		{
			name:     "two different characters",
			input:    "ab",
			expected: [][]string{{"a", "b"}},
		},
		{
			name:     "two same characters",
			input:    "aa",
			expected: [][]string{{"a", "a"}, {"aa"}},
		},
		{
			name:     "example case",
			input:    "aab",
			expected: [][]string{{"a", "a", "b"}, {"aa", "b"}},
		},
		{
			name:     "longer palindrome",
			input:    "racecar",
			expected: [][]string{{"r", "a", "c", "e", "c", "a", "r"}, {"r", "a", "cec", "a", "r"}, {"r", "aceca", "r"}, {"racecar"}},
		},
		{
			name:  "complex palindrome",
			input: "ababbbabbaba",
			expected: [][]string{
				{"a", "b", "a", "b", "b", "b", "a", "b", "b", "a", "b", "a"},
				{"a", "b", "a", "b", "b", "b", "a", "b", "b", "aba"},
				{"a", "b", "a", "b", "b", "b", "a", "b", "bab", "a"},
				{"a", "b", "a", "b", "b", "b", "a", "bb", "a", "b", "a"},
				{"a", "b", "a", "b", "b", "b", "a", "bb", "aba"},
				{"a", "b", "a", "b", "b", "b", "abba", "b", "a"},
				{"a", "b", "a", "b", "b", "bab", "b", "a", "b", "a"},
				{"a", "b", "a", "b", "b", "bab", "b", "aba"},
				{"a", "b", "a", "b", "b", "bab", "bab", "a"},
				{"a", "b", "a", "b", "b", "babbab", "a"},
				{"a", "b", "a", "b", "bb", "a", "b", "b", "a", "b", "a"},
				{"a", "b", "a", "b", "bb", "a", "b", "b", "aba"},
				{"a", "b", "a", "b", "bb", "a", "b", "bab", "a"},
				{"a", "b", "a", "b", "bb", "a", "bb", "a", "b", "a"},
				{"a", "b", "a", "b", "bb", "a", "bb", "aba"},
				{"a", "b", "a", "b", "bb", "abba", "b", "a"},
				{"a", "b", "a", "b", "bbabb", "a", "b", "a"},
				{"a", "b", "a", "b", "bbabb", "aba"},
				{"a", "b", "a", "bb", "b", "a", "b", "b", "a", "b", "a"},
				{"a", "b", "a", "bb", "b", "a", "b", "b", "aba"},
				{"a", "b", "a", "bb", "b", "a", "b", "bab", "a"},
				{"a", "b", "a", "bb", "b", "a", "bb", "a", "b", "a"},
				{"a", "b", "a", "bb", "b", "a", "bb", "aba"},
				{"a", "b", "a", "bb", "b", "abba", "b", "a"},
				{"a", "b", "a", "bb", "bab", "b", "a", "b", "a"},
				{"a", "b", "a", "bb", "bab", "b", "aba"},
				{"a", "b", "a", "bb", "bab", "bab", "a"},
				{"a", "b", "a", "bb", "babbab", "a"},
				{"a", "b", "a", "bbb", "a", "b", "b", "a", "b", "a"},
				{"a", "b", "a", "bbb", "a", "b", "b", "aba"},
				{"a", "b", "a", "bbb", "a", "b", "bab", "a"},
				{"a", "b", "a", "bbb", "a", "bb", "a", "b", "a"},
				{"a", "b", "a", "bbb", "a", "bb", "aba"},
				{"a", "b", "a", "bbb", "abba", "b", "a"},
				{"a", "b", "abbba", "b", "b", "a", "b", "a"},
				{"a", "b", "abbba", "b", "b", "aba"},
				{"a", "b", "abbba", "b", "bab", "a"},
				{"a", "b", "abbba", "bb", "a", "b", "a"},
				{"a", "b", "abbba", "bb", "aba"},
				{"a", "bab", "b", "b", "a", "b", "b", "a", "b", "a"},
				{"a", "bab", "b", "b", "a", "b", "b", "aba"},
				{"a", "bab", "b", "b", "a", "b", "bab", "a"},
				{"a", "bab", "b", "b", "a", "bb", "a", "b", "a"},
				{"a", "bab", "b", "b", "a", "bb", "aba"},
				{"a", "bab", "b", "b", "abba", "b", "a"},
				{"a", "bab", "b", "bab", "b", "a", "b", "a"},
				{"a", "bab", "b", "bab", "b", "aba"},
				{"a", "bab", "b", "bab", "bab", "a"},
				{"a", "bab", "b", "babbab", "a"},
				{"a", "bab", "bb", "a", "b", "b", "a", "b", "a"},
				{"a", "bab", "bb", "a", "b", "b", "aba"},
				{"a", "bab", "bb", "a", "b", "bab", "a"},
				{"a", "bab", "bb", "a", "bb", "a", "b", "a"},
				{"a", "bab", "bb", "a", "bb", "aba"},
				{"a", "bab", "bb", "abba", "b", "a"},
				{"a", "bab", "bbabb", "a", "b", "a"},
				{"a", "bab", "bbabb", "aba"},
				{"a", "babbbab", "b", "a", "b", "a"},
				{"a", "babbbab", "b", "aba"},
				{"a", "babbbab", "bab", "a"},
				{"aba", "b", "b", "b", "a", "b", "b", "a", "b", "a"},
				{"aba", "b", "b", "b", "a", "b", "b", "aba"},
				{"aba", "b", "b", "b", "a", "b", "bab", "a"},
				{"aba", "b", "b", "b", "a", "bb", "a", "b", "a"},
				{"aba", "b", "b", "b", "a", "bb", "aba"},
				{"aba", "b", "b", "b", "abba", "b", "a"},
				{"aba", "b", "b", "bab", "b", "a", "b", "a"},
				{"aba", "b", "b", "bab", "b", "aba"},
				{"aba", "b", "b", "bab", "bab", "a"},
				{"aba", "b", "b", "babbab", "a"},
				{"aba", "b", "bb", "a", "b", "b", "a", "b", "a"},
				{"aba", "b", "bb", "a", "b", "b", "aba"},
				{"aba", "b", "bb", "a", "b", "bab", "a"},
				{"aba", "b", "bb", "a", "bb", "a", "b", "a"},
				{"aba", "b", "bb", "a", "bb", "aba"},
				{"aba", "b", "bb", "abba", "b", "a"},
				{"aba", "b", "bbabb", "a", "b", "a"},
				{"aba", "b", "bbabb", "aba"},
				{"aba", "bb", "b", "a", "b", "b", "a", "b", "a"},
				{"aba", "bb", "b", "a", "b", "b", "aba"},
				{"aba", "bb", "b", "a", "b", "bab", "a"},
				{"aba", "bb", "b", "a", "bb", "a", "b", "a"},
				{"aba", "bb", "b", "a", "bb", "aba"},
				{"aba", "bb", "b", "abba", "b", "a"},
				{"aba", "bb", "bab", "b", "a", "b", "a"},
				{"aba", "bb", "bab", "b", "aba"},
				{"aba", "bb", "bab", "bab", "a"},
				{"aba", "bb", "babbab", "a"},
				{"aba", "bbb", "a", "b", "b", "a", "b", "a"},
				{"aba", "bbb", "a", "b", "b", "aba"},
				{"aba", "bbb", "a", "b", "bab", "a"},
				{"aba", "bbb", "a", "bb", "a", "b", "a"},
				{"aba", "bbb", "a", "bb", "aba"},
				{"aba", "bbb", "abba", "b", "a"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := partition(tt.input)
			if !equalStringSlices(got, tt.expected) {
				t.Errorf("partition() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// Helper function to compare two string slice slices regardless of order
func equalStringSlices(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}
	used := make([]bool, len(b))
	for _, sliceA := range a {
		found := false
		for j, sliceB := range b {
			if !used[j] && equalSlice(sliceA, sliceB) {
				used[j] = true
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// Helper function to compare two string slices
func equalSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
