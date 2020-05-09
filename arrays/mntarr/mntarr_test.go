package mntarr

import "testing"

func TestCheckIfExist(t *testing.T) {
	testCases := []struct {
		desc string
		arr  []int
		want bool
	}{
		{"nil slice given, false returned", nil, false},
		{"empty slice given, false returned", []int{}, false},
		{"less than three elements given, false returned", []int{27, 42}, false},
		{"two same elements in a row given, false returned", []int{7, 12, 12}, false},
		{"valid mountain array given, true returned", []int{0, 3, 2, 1}, true},
		{"monotonically ascending array given, false returned", []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, false},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			got := validMountainArray(tc.arr)
			assertEquals(t, tc.want, got)
		})
	}
}

func assertEquals(t testing.TB, want, got bool) {
	t.Helper()
	if want != got {
		t.Fatalf("want %v, got %v", want, got)
	}
}
