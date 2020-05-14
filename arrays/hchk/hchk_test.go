package hchk

import "testing"

func TestHeightChecker(t *testing.T) {
	testCases := []struct {
		desc    string
		heights []int
		want    int
	}{
		{"nil slice given, nothing done", nil, 0},
		{"empty slice given, nothing done", []int{}, 0},
		{"ordered elements given, nothing done", []int{1, 2, 3, 4, 5}, 0},
		{"two unordered elements given, array re-arranged", []int{2, 1}, 2},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			got := heightChecker(tc.heights)
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
