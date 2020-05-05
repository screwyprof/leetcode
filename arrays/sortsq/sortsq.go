package sortsq

func sortedSquares(a []int) []int {
	n := len(a)
	j := 0

	for j < n && a[j] < 0 {
		j++
	}

	i := j - 1

	res := make([]int, 0, len(a))
	for 0 <= i && j < n {
		nEl := a[i] * a[i]
		pEl := a[j] * a[j]
		if nEl < pEl {
			res = append(res, nEl)
			i -= 1
		} else {
			res = append(res, pEl)
			j += 1
		}
	}

	for i >= 0 {
		res = append(res, a[i]*a[i])
		i -= 1
	}

	for j < n {
		res = append(res, a[j]*a[j])
		j += 1
	}

	return res
}
