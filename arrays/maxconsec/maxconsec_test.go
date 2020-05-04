package maxconsec

import "testing"

func TestFindMaxConsecutiveOnes(t *testing.T) {
	t.Run("nil slice given, 0 returned", func(t *testing.T) {
		t.Parallel()

		got := findMaxConsecutiveOnes(nil)
		assertEquals(t, 0, got)
	})

	t.Run("empty slice given, 0 returned", func(t *testing.T) {
		t.Parallel()

		got := findMaxConsecutiveOnes([]int{})
		assertEquals(t, 0, got)
	})
}

func assertEquals(t testing.TB, want, got int) {
	t.Helper()
	if want != got {
		t.Fatalf("want %d, got %d", want, got)
	}
}
