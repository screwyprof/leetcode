package binsearch

import "testing"

func TestSearch(t *testing.T) {
	testCases := []struct {
		desc     string
		haystack []int
		needle   int
		want     int
	}{
		{"nil slice given, not found", nil, 42, -1},
		{"empty slice given, not found", []int{}, 7, -1},
		{"odd array given, not found", []int{0, 2, 4, 6, 8, 10, 12, 14, 16}, 42, -1},
		{"even array given, not found", []int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18}, 42, -1},
		{"element found at the beginning", []int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18}, 0, 0},
		{"element found at the end", []int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18}, 18, 9},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			got := Search(tc.haystack, tc.needle)
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
