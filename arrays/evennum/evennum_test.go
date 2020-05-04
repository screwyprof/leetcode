package evennum

import "testing"

func TestFindNumbers(t *testing.T) {
	testCases := []struct {
		desc string
		nums []int
		want int
	}{
		{desc: "nil slice given, 0 returned", nums: nil, want: 0},
		{desc: "empty slice given, 0 returned", nums: []int{}, want: 0},
		{desc: "one odd number of digits given, 0 returned", nums: []int{0}, want: 0},
		//{desc: "one even number given of digits, 1 returned", nums: []int{11}, want: 1},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			got := findNumbers(tc.nums)
			assertEquals(t, tc.want, got)
		})
	}
}

func TestDigits(t *testing.T) {
	testCases := []struct {
		desc string
		num  int
		want int
	}{
		{desc: "0 given, 0 returned", num: 0, want: 1},
		{desc: "1 given, 1 returned", num: 1, want: 1},
		{desc: "27 given, 2 returned", num: 27, want: 2},
		{desc: "115 given, 3 returned", num: 115, want: 3},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			got := digits(tc.num)
			assertEquals(t, tc.want, got)
		})
	}
}

func assertEquals(t testing.TB, want, got int) {
	t.Helper()
	if want != got {
		t.Fatalf("want %d, got %d", want, got)
	}
}
