package mergesort

// merge merges nums2 into nums1 as one sorted array.
// The number of elements initialized in nums1 and nums2 are m and n respectively.
// You may assume that nums1 has enough space (size that is greater or equal to m + n)
// to hold additional elements from nums2.
func merge(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m-1:], nums2[:])
}
