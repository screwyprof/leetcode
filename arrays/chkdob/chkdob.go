package chkdup

func checkIfExist(arr []int) bool {
	hash := make(map[int]int, len(arr))
	for i := 0; i < len(arr); i++ {
		hash[arr[i]] = i
	}

	for i := range arr {
		if idx, ok := hash[2*arr[i]]; ok {
			if idx != i {
				return true
			}
		}
	}
	return false
}
