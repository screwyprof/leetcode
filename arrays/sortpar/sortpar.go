package sortpar

func sortArrayByParity(arr []int) []int {
	pos := 0
	for i := range arr {
		if arr[i]&1 == 0 {
			arr[pos], arr[i] = arr[i], arr[pos]
			pos++
		}
	}
	return arr
}
