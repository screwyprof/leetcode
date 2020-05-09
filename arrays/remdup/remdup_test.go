package remdup

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	testCases := []struct {
		desc   string
		nums   []int
		length int
		want   []int
	}{
		{"nil slice given, 0 returned", nil, 0, nil},
		{"empty slice given, 0 returned", []int{}, 0, []int{}},
		{"one element given, 1 returned, nothing changed", []int{27}, 1, []int{27}},
		{"no duplicates given, nothing changed, original length returned", []int{27, 12, 42}, 3, []int{27, 12, 42}},
		{"one duplicate at the beginning given, duplicate removed, length changed", []int{27, 27, 42}, 2, []int{27, 42}},
		{"a few duplicates at the beginning given, duplicates removed, length changed", []int{27, 27, 27, 27, 42}, 2, []int{27, 42}},
		{"one duplicate at the end given, duplicate removed, length changed", []int{27, 42, 42}, 2, []int{27, 42}},
		{"a few duplicates at the end given, duplicates removed, length changed", []int{27, 42, 42, 42, 42}, 2, []int{27, 42}},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			got := removeDuplicates(tc.nums)
			assertEquals(t, tc.length, got)
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
