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
		{"non zero elements given, nothing done", []int{42, 27, 100}, []int{42, 27, 100}},
		{"two elements, last zero given, nothing done", []int{42, 0}, []int{42, 0}},
		{"two elements, first zero, second overwritten with zero", []int{0, 42}, []int{0, 0}},
		{"three elements, first zero, second overwritten with zero", []int{0, 42, 7}, []int{0, 0, 42}},
		{"an arbitrary array with zeros given, zeros duplicated", []int{1, 0, 2, 3, 0, 4, 5, 0}, []int{1, 0, 0, 2, 3, 0, 0, 4}},
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
