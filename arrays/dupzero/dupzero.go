package dupzero

// duplicateZeros duplicates each occurrence of zero, shifting the remaining elements to the right.
// Elements beyond the length of the original array are not written.
func duplicateZeros(arr []int) {
	var i, shifts int
	for i = 0; shifts+i < len(arr); i++ {
		if arr[i] == 0 {
			shifts++
		}
	}

	for i = i - 1; shifts > 0; i-- {
		if i+shifts < len(arr) {
			arr[i+shifts] = arr[i]
		}

		if arr[i] == 0 {
			shifts--
			arr[i+shifts] = arr[i]
		}
	}
}
