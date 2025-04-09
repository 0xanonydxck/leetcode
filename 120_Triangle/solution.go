package triangle

import "fmt"

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}

	if len(triangle) == 1 {
		return triangle[0][0]
	}

	grid := make([]int, len(triangle[len(triangle)-1]))
	copy(grid, triangle[len(triangle)-1])

	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			grid[j] = Min(grid[j], grid[j+1]) + triangle[i][j]
		}
	}

	fmt.Println(grid)
	return grid[0]
}
