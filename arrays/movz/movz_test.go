package movz

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	testCases := []struct {
		desc string
		arr  []int
		want []int
	}{
		{"nil slice given, nothing done", nil, nil},
		{"empty slice given, nothing done", []int{}, []int{}},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			moveZeroes(tc.arr)
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
