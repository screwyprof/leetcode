package sortsq

func sortedSquares(a []int) []int {
	n := len(a)
	l, r := 0, n-1
	pos, res := n-1, make([]int, n)

	for l <= r {
		if a[l]*-1 > a[r] {
			res[pos] = a[l] * a[l]
			l++
		} else {
			res[pos] = a[r] * a[r]
			r--
		}

		pos--
	}

	return res
}
