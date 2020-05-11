package sortsq

import (
	"reflect"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	t.Run("test case #1", func(t *testing.T) {
		want := []int{0, 1, 9, 16, 100}
		got := sortedSquares([]int{-4, -1, 0, 3, 10})
		assertEqual(t, want, got)
	})

	t.Run("test case #2", func(t *testing.T) {
		want := []int{4, 9, 9, 49, 121}
		got := sortedSquares([]int{-7, -3, 2, 3, 11})
		assertEqual(t, want, got)
	})
}

func assertEqual(t testing.TB, want, got []int) {
	t.Helper()
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v, got %v", want, got)
	}
}

var result []int

func BenchmarkSortedSquares(b *testing.B) {
	var res []int
	for i := 0; i < b.N; i++ {
		res = sortedSquares([]int{-7, -3, 2, 3, 11})
	}
	result = res
}

func BenchmarkSortedSquares2(b *testing.B) {
	var res []int
	for i := 0; i < b.N; i++ {
		res = sortedSquares2([]int{-7, -3, 2, 3, 11})
	}
	result = res
}
