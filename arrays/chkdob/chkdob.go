package chkdup

func checkIfExist(arr []int) bool {
	hash := make(map[int]*struct{}, len(arr))
	for _, v := range arr {
		if hash[v*2] != nil {
			return true
		}
		if v&1 == 0 && hash[v>>1] != nil {
			return true
		}
		hash[v] = &struct{}{}
	}
	return false
}
