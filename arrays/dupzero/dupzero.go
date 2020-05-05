package dupzero

// duplicateZeros duplicates each occurrence of zero, shifting the remaining elements to the right.
// Elements beyond the length of the original array are not written.
func duplicateZeros(arr []int) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			copy(arr[i+1:], arr[i:])
			i++
		}
	}
}
