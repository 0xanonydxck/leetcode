package genparentheses

func generateParenthesis(n int) []string {
	result := make([]string, 0)
	if n == 0 {
		return result
	}

	var dfs func(n, open, close int, path []rune)
	dfs = func(n, open, close int, path []rune) {
		if open == n && close == n {
			result = append(result, string(path))
			return
		}

		if open < close || open > n || close > n {
			return
		}

		dfs(n, open+1, close, append(path, '('))
		dfs(n, open, close+1, append(path, ')'))
	}

	dfs(n, 0, 0, make([]rune, 0))
	return result
}
