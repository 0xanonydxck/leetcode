package maximalsquare

func Min(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}

	result := nums[0]
	for i := range nums {
		if nums[i] < result {
			result = nums[i]
		}
	}

	return result
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	n, m := len(matrix), len(matrix[0])
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
	}

	maxSide := 0
	for i := 0; i < n; i++ {
		dp[i][0] = int(matrix[i][0] - '0')
		maxSide = Max(maxSide, dp[i][0])
	}

	for j := 0; j < m; j++ {
		dp[0][j] = int(matrix[0][j] - '0')
		maxSide = Max(maxSide, dp[0][j])
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = Min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
				maxSide = Max(maxSide, dp[i][j])
			}
		}
	}

	return maxSide * maxSide
}
