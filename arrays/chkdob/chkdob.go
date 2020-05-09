package chkdup

func checkIfExist(arr []int) bool {
	hash := make(map[int]struct{}, len(arr))
	for i := 0; i < len(arr); i++ {
		hash[arr[i]] = struct{}{}
	}

	for i := range arr {
		if _, ok := hash[2*arr[i]]; ok {
			return true
		}
	}
	return false
}
