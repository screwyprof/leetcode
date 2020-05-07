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
