package sortsq

// sortedSquares returns an array of the squares of each number in sorted non-decreasing order.
// a is expected to be sorted in non-decreasing order.
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

// sortedSquares returns an array of the squares of each number in sorted non-decreasing order.
// It sorts elements in place.
// a is expected to be sorted in non-decreasing order.
func sortedSquares2(a []int) []int {
	for i := 0; i < len(a); i++ {
		a[i] = a[i] * a[i]
		key := a[i]
		j := i - 1
		for j >= 0 && a[j] > key {
			a[j+1] = a[j]
			j--
		}
		a[j+1] = key
	}
	return a
}
