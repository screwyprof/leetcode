package maxconsec

import "testing"

func TestFindMaxConsecutiveOnes(t *testing.T) {
	testCases := []struct {
		desc string
		nums []int
		want int
	}{
		{desc: "nil slice given, 0 returned", nums: nil, want: 0},
		{desc: "empty slice given, 0 returned", nums: []int{}, want: 0},
		{desc: "one '1' given, 1 returned", nums: []int{1}, want: 1},
		{desc: "one '0' given, 0 returned", nums: []int{0}, want: 0},
		{desc: "two consecutive 1's and no 0's given, 2 returned", nums: []int{1, 1}, want: 2},
		{desc: "two consecutive 1's followed by one 0 given, 2 returned", nums: []int{1, 1, 0}, want: 2},
		{desc: "two consecutive 1's following two 0's given, 2 returned", nums: []int{0, 0, 1, 1}, want: 2},
		{desc: "a few non consecutive figures given, 1 returned", nums: []int{1, 0, 1, 0}, want: 1},
		{desc: "two 1's windows given, one with 2 elements, second with 3, 3 returned ", nums: []int{1, 1, 0, 1, 1, 1}, want: 3},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			got := findMaxConsecutiveOnes(tc.nums)
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
