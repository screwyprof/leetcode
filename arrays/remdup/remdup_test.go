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
