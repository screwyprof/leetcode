package maxconsec

// findMaxConsecutiveOnes finds the maximum number of consecutive 1s in the given array.
// Example:
// Input: [1,1,0,1,1,1]
// Output: 3
// Explanation: The first two digits or the last three digits are consecutive 1s.
// 		The maximum number of consecutive 1s is 3.
//
// The input array will only contain 0 and 1.
//
// See Max Consecutive Ones (https://leetcode.com/explore/learn/card/fun-with-arrays/521/introduction/3238/)
//
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
