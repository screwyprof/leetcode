package binsearch

func Search(haystack []int, needle int) int {
	l := 0
	r := len(haystack)

	for l < r {
		h := int(uint(l+r) >> 1)
		if haystack[h] == needle {
			return h
		}

		if haystack[h] > needle {
			r = h
		} else {
			l = h + 1
		}
	}

	return -1
}
