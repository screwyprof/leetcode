package mergesort

import "github.com/davecgh/go-spew/spew"

// merge merges nums2 into nums1 as one sorted array.
// The number of elements initialized in nums1 and nums2 are m and n respectively.
// You may assume that nums1 has enough space (size that is greater or equal to m + n)
// to hold additional elements from nums2.
func merge(nums1 []int, m int, nums2 []int, n int) {
	res := make([]int, m+n)

	spew.Dump(nums1[:m])
	spew.Dump(nums2)
	for i := 0; i < n; i++ {
		if nums1[i] > nums2[i] {
			res[i] = nums2[i]
			res[i+1] = nums1[i]
		} else {
			res[i] = nums1[i]
			res[i+1] = nums2[i]
		}
	}

	spew.Dump(res)
	copy(nums1[:], res[:])
}
