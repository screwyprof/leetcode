package mntarr

func validMountainArray(arr []int) bool {
	n := len(arr) - 1

	l := 0
	r := n

	for l < n && arr[l] < arr[l+1] {
		l++
	}

	for r > 0 && arr[r-1] > arr[r] {
		r--
	}

	return l > 0 && l == r && r < n
}
