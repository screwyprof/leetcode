package sortsq

// sortedSquares returns an array of the squares of each number in sorted non-decreasing order.
// a is expected to be sorted in non-decreasing order.
//
// Example 1:
// Input: [-4,-1,0,3,10]
// Output: [0,1,9,16,100]
//
// Example 2:
// Input: [-7,-3,2,3,11]
// Output: [4,9,9,49,121]
//
// See Squares of a Sorted Array (https://leetcode.com/explore/learn/card/fun-with-arrays/521/introduction/3240/)
//
func sortedSquares(a []int) []int {
	n := len(a)
	l, r := 0, n-1
	res := make([]int, n)
	for pos := n - 1; pos >= 0; pos-- {
		if a[l]*-1 > a[r] {
			res[pos] = a[l] * a[l]
			l++
		} else {
			res[pos] = a[r] * a[r]
			r--
		}
	}

	return res
}
