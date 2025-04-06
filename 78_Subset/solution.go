package subset

func subsets(nums []int) [][]int {
	result, size := [][]int{}, len(nums)
	var dfs func(level int, include bool, path []int)
	dfs = func(level int, include bool, path []int) {
		if include {
			result = append(result, append([]int{}, path...))
		}

		if level == size {
			return
		}

		dfs(level+1, true, append(path, nums[level]))
		dfs(level+1, false, path)
	}

	dfs(0, true, make([]int, 0))
	return result
}
