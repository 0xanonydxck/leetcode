package findlatestgroupofsizem

func findLatestStep(arr []int, m int) int {
	if len(arr) == m {
		return len(arr)
	}

	binary := make([]int, len(arr))
	for i := range binary {
		binary[i] = 1
	}

	for i := len(arr) - 1; i >= 0; i-- {
		index := arr[i] - 1
		binary[index] = 0

		if check(binary, m) {
			return i
		}
	}

	return -1
}

func check(binary []int, m int) bool {
	count := 0
	for _, item := range binary {
		if item == 1 {
			count++
		} else {
			if count == m {
				return true
			}

			count = 0
		}
	}

	if count == m {
		return true
	} else {
		return false
	}
}
