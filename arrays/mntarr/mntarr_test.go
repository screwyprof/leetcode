package mntarr

import "testing"

func TestCheckIfExist(t *testing.T) {
	testCases := []struct {
		desc string
		arr  []int
		want bool
	}{
		{"nil slice given, not valid", nil, false},
		{"empty slice given, not valid", []int{}, false},
		{"less than three elements given, not valid", []int{27, 42}, false},
		{"two same elements in a row given, not valid", []int{7, 12, 12}, false},
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
