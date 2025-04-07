package minimumpathsum

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minPathSum(grid [][]int) int {
	rowSize, colSize := len(grid), len(grid[0])
	result := make([][]int, rowSize)
	for i := 0; i < rowSize; i++ {
		result[i] = make([]int, colSize)
	}

	for row := 0; row < rowSize; row++ {
		for col := 0; col < colSize; col++ {
			leftIndex, topIndex := row-1, col-1
			if leftIndex < 0 && topIndex < 0 {
				result[row][col] = grid[row][col]
			} else if leftIndex < 0 {
				result[row][col] = result[row][topIndex] + grid[row][col]
			} else if topIndex < 0 {
				result[row][col] = result[leftIndex][col] + grid[row][col]
			} else {
				result[row][col] = Min(result[leftIndex][col], result[row][topIndex]) + grid[row][col]
			}
		}
	}

	return result[rowSize-1][colSize-1]
}
