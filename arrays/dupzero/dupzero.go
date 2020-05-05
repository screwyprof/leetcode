package dupzero

// duplicateZeros duplicates each occurrence of zero, shifting the remaining elements to the right.
// Elements beyond the length of the original array are not written.
func duplicateZeros(arr []int) {
	n := len(arr)
	if n == 0 {
		return
	}

	for i := 0; i < n; i++ {
		if arr[i] == 0 {
			if i+1 < n {
				arr[i+1] = 0
			}
			i++
		}
	}
}
