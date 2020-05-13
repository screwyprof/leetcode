package sortpar

import (
	"reflect"
	"testing"
)

func TestSortArrayByParity(t *testing.T) {
	testCases := []struct {
		desc string
		arr  []int
		want []int
	}{
		{"nil slice given, nothing done", nil, nil},
		{"empty slice given, nothing done", []int{}, []int{}},
		{"one element given, nothing done", []int{0}, []int{0}},
		{"all event elements given, nothing done", []int{0, 42, 30}, []int{0, 42, 30}},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			got := sortArrayByParity(tc.arr)
			assertEquals(t, tc.want, got)
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
