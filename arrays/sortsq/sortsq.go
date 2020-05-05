package sortsq

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
