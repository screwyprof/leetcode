package remdup

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	j := 0
	for i := 1; i < len(nums); i++ {
		if nums[j] == nums[i] {
			continue
		}
		j++
		// preserve the original data
		// in[i], in[j] = in[j], in[i]
		// only set what is required
		nums[j] = nums[i]
	}
	nums = nums[:j+1]
	return j + 1
}
