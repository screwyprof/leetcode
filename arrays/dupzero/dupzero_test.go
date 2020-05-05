package dupzero

import (
	"reflect"
	"testing"
)

func TestDuplicateZeros(t *testing.T) {
	testCases := []struct {
		desc string
		arr  []int
		want []int
	}{
		{"nil slice given, nothing done", nil, nil},
		{"empty slice given, nothing done", []int{}, []int{}},
		{"one zero element given, nothing done", []int{0}, []int{0}},
		{"one non-zero element given, nothing done", []int{1}, []int{1}},
		{"only zero elements given, nothing done", []int{0, 0, 0}, []int{0, 0, 0}},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			duplicateZeros(tc.arr)
			assertEquals(t, tc.want, tc.arr)
		})
	}
}

func assertEquals(t testing.TB, want, got []int) {
	t.Helper()
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v, got %v", want, got)
	}
}
