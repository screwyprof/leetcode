package repgr

func replaceElements(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	arr[len(arr)-1] = -1
	return arr
}
