package palindromepartitioning

func partition(s string) [][]string {
	result := make([][]string, 0)
	if len(s) == 0 {
		return result
	}

	var dfs func(prefix string, path []string)
	dfs = func(prefix string, path []string) {
		if len(prefix) == 0 {
			result = append(result, append([]string{}, path...))
			return
		}

		for i := 0; i < len(prefix); i++ {
			sub, rest := prefix[:i+1], prefix[i+1:]
			if isPalindrome(sub) {
				dfs(rest, append(path, sub))
			}
		}
	}

	dfs(s, make([]string, 0))
	return result
}

func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}

	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}

		left++
		right--
	}

	return true
}
