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
		{desc: "nil slice given, 0 returned", arr: nil, want: nil},
		{desc: "empty slice given, 0 returned", arr: []int{}, want: []int{}},
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
