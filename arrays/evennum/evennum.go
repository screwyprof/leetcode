package evennum

// findNumbers returns how many of nums's elements contain an even number of digits.
func findNumbers(nums []int) int {
	var evens int
	for _, num := range nums {
		if digits(num)%2 == 0 {
			evens++
		}
	}
	return evens
}

func digits(n int) int {
	if n == 0 {
		return 1
	}
	var count int
	for n != 0 {
		n /= 10
		count++
	}
	return count
}
