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
		{"only even elements given, nothing done", []int{0, 42, 30}, []int{0, 42, 30}},
		{"only odd elements given, nothing done", []int{1, 7, 27}, []int{1, 7, 27}},
		{"even elements followed by odd elements given, nothing done", []int{42, 30, 1, 7}, []int{42, 30, 1, 7}},
		{"an odd element followed by an even given, elements swapped", []int{7, 42}, []int{42, 7}},
		{"arbitrary array with odd and even elements given, elements rearranged", []int{3, 1, 2, 4}, []int{2, 4, 3, 1}},
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
