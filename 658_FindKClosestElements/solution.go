package findkclosestelements

const (
	intSize = 32 << (^uint(0) >> 63) // 32 or 64
	MaxInt  = 1<<(intSize-1) - 1     // MaxInt32 or MaxInt64 depending on intSize.
)

func Abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

func findStartIndex(arr []int, target int) (int, int) {
	result, min := -1, MaxInt
	for i := range arr {
		diff := Abs(target - arr[i])
		if diff < min {
			result = i
			min = diff
		}
	}

	return result, min
}

func isReachK(start, end, k int) bool {
	size := end - start + 1
	return size == k
}

func findClosestElements(arr []int, k int, target int) []int {
	size := len(arr)
	if size == 0 {
		return make([]int, 0)
	} else if size <= k {
		return arr
	} else if arr[0] > target {
		return arr[0:k]
	} else if arr[size-1] < target {
		return arr[size-k:]
	}

	startIndex, targetDiff := findStartIndex(arr, target)
	start, end := startIndex, startIndex
	for {
		for i := start - 1; i >= 0; i-- {
			if isReachK(start, end, k) {
				break
			} else {
				diff := Abs(arr[i] - target)
				if diff != targetDiff {
					break
				}

				start--
			}
		}

		for i := end + 1; i < len(arr); i++ {
			if isReachK(start, end, k) {
				break
			} else {
				diff := Abs(arr[i] - target)
				if diff != targetDiff {
					break
				}
				end++
			}
		}

		if isReachK(start, end, k) {
			break
		}
		targetDiff++
	}

	return arr[start : end+1]
}
