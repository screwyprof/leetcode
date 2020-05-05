package mergesort

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	testCases := []struct {
		desc  string
		nums1 []int
		nums2 []int
		m     int
		n     int
		want  []int
	}{
		{
			desc:  "num1 and num 2 with the same one element, two elements merged",
			nums1: []int{7, 0},
			nums2: []int{7},
			m:     1,
			n:     1,
			want:  []int{7, 7},
		},
		{
			desc:  "num1 with one element, num2 with one element, num2's element is bigger, elements merged",
			nums1: []int{7, 0},
			nums2: []int{42},
			m:     1,
			n:     1,
			want:  []int{7, 42},
		},
		{
			desc:  "num1 with one element, num2 with one element, num1's element is bigger, elements merged",
			nums1: []int{42, 0},
			nums2: []int{7},
			m:     1,
			n:     1,
			want:  []int{7, 42},
		},
		{
			desc:  "arbitrary arrays given, elements merged",
			nums1: []int{1, 2, 3, 0, 0, 0},
			nums2: []int{2, 5, 6},
			m:     3,
			n:     3,
			want:  []int{1, 2, 2, 3, 5, 6},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			merge(tc.nums1, tc.m, tc.nums2, tc.n)
			assertEquals(t, tc.want, tc.nums1)
		})
	}
}

func assertEquals(t testing.TB, want, got []int) {
	t.Helper()
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v, got %v", want, got)
	}
}
