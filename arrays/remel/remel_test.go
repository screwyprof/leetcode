package remel

import (
	"reflect"
	"sort"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	testCases := []struct {
		desc   string
		nums   []int
		val    int
		length int
		want   []int
	}{
		{"nil slice given, 0 returned", nil, 0, 0, nil},
		{"empty slice given, 0 returned", []int{}, 0, 0, []int{}},
		{"val not found, nothing changed", []int{1, 2, 3}, 0, 3, []int{1, 2, 3}},
		{"val found at the beginning, val removed", []int{42, 6, 19}, 42, 2, []int{6, 19}},
		{"2 val's found at the beginning, val's removed", []int{42, 42, 19}, 42, 1, []int{19}},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			got := removeElement(tc.nums, tc.val)
			assertEquals(t, tc.length, got)
			sort.Ints(tc.nums)
			assertEquals(t, tc.want, tc.nums[:tc.length])
		})
	}
}

func assertEquals(t testing.TB, want, got interface{}) {
	t.Helper()
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v, got %v", want, got)
	}
}
