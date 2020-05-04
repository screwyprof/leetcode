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
		{desc: "one element given, 0 returned", nums: []int{1}, want: 0},
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
