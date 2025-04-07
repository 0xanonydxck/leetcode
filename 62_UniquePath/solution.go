package uniquepath

/**
 * m -> number of rows
 * n -> number of colums
 */
func uniquePaths(m int, n int) int {
	grid := make([][]int, m)
	for i := 0; i < m; i++ {
		grid[i] = make([]int, n)
	}

	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			top := row - 1
			left := col - 1

			if top < 0 || left < 0 {
				grid[row][col] = 1
			} else {
				grid[row][col] = grid[top][col] + grid[row][left]
			}
		}
	}

	return grid[m-1][n-1]
}
