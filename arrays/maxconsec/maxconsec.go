package maxconsec

// findMaxConsecutiveOnes finds the maximum number of consecutive 1s in the given array.
func findMaxConsecutiveOnes(nums []int) int {
	var (
		res    int
		window []int
	)
	for _, n := range nums {
		if n == 1 {
			window = append(window, 1)
		} else {
			res = max(res, len(window))
			window = nil
		}
	}
	res = max(res, len(window))
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
