package mergesort

// merge merges nums2 into nums1 as one sorted array.
// The number of elements initialized in nums1 and nums2 are m and n respectively.
// You may assume that nums1 has enough space (size that is greater or equal to m + n)
// to hold additional elements from nums2.
func merge(nums1 []int, m int, nums2 []int, n int) {
	cnt := m + n - 1
	m--
	n--
	for n >= 0 {
		if m >= 0 && nums1[m] > nums2[n] {
			nums1[cnt] = nums1[m]
			m--
		} else {
			nums1[cnt] = nums2[n]
			n--
		}
		cnt--
	}
}
