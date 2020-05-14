package hchk

import "sort"

func heightChecker(heights []int) int {
	sorted := append(heights[:0:0], heights...)
	sort.Ints(sorted)

	var cnt int
	for i := range heights {
		if heights[i] != sorted[i] {
			cnt++
		}
	}

	return cnt
}
