package maxconsec

// findMaxConsecutiveOnes finds the maximum number of consecutive 1s in the given array.
func findMaxConsecutiveOnes(nums []int) int {
	var (
		res int
		cnt int
	)
	for _, n := range nums {
		if n == 0 {
			cnt = 0
		} else {
			cnt++
			res = max(res, cnt)
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
