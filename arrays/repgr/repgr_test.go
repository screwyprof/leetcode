package repgr

import (
	"reflect"
	"testing"
)

func TestReplaceElements(t *testing.T) {
	testCases := []struct {
		desc string
		arr  []int
		want []int
	}{
		{"nil slice given, nothing done", nil, nil},
		{"empty slice given, nothing done", []int{}, []int{}},
		{"one element given, [-1] returned", []int{42}, []int{-1}},
		{"two elements given, [the max of two, -1] returned", []int{7, 42}, []int{42, -1}},
		{"arbitrary elements given, elements replaced", []int{17, 18, 5, 4, 6, 1}, []int{18, 6, 6, 6, 1, -1}},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			got := replaceElements(tc.arr)
			assertEquals(t, tc.want, got)
		})
	}
}

func assertEquals(t testing.TB, want, got []int) {
	t.Helper()
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v, got %v", want, got)
	}
}
