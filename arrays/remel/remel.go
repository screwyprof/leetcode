package remel

// removeElement removes all instances of  val in-place and returns the new length.
func removeElement(nums []int, val int) int {
	l := 0
	for _, n := range nums {
		if n != val {
			nums[l] = n
			l++
		}
	}
	return l
}
