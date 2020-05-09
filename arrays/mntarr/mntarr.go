package mntarr

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	var peakReached bool
	var asc bool
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == arr[i+1] {
			return false
		}

		if !peakReached && arr[i] < arr[i+1] {
			asc = true
			continue
		}

		peakReached = true

		if !asc {
			return false
		}

		if arr[i] <= arr[i+1] {
			return false
		}
	}

	if !peakReached {
		return false
	}

	return true
}
