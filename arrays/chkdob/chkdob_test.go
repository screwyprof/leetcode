package chkdup

import "testing"

func TestCheckIfExist(t *testing.T) {
	testCases := []struct {
		desc string
		arr  []int
		want bool
	}{
		{"nil slice given, not found", nil, false},
		{"empty slice given, not found", []int{}, false},
		{"one element given, not found", []int{27}, false},
		{"two elements given, not found", []int{27, 42}, false},
		{"arbitrary elements given, not found", []int{3, 1, 7, 11}, false},
		{"two elements given, the double comes first, found", []int{84, 42}, true},
		{"two elements given, the double comes last, found", []int{42, 84}, true},
		{"a few elements given, the double comes before the value, found", []int{10, 2, 5, 3}, true},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			got := checkIfExist(tc.arr)
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
