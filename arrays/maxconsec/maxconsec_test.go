package maxconsec

import "testing"

func TestFindMaxConsecutiveOnes(t *testing.T) {
	t.Run("nil slice given, 0 returned", func(t *testing.T) {
		got := findMaxConsecutiveOnes(nil)
		assertEquals(t, 0, got)
	})
}

func assertEquals(t testing.TB, want, got int) {
	t.Helper()
	if want != got {
		t.Fatalf("want %d, got %d", want, got)
	}
}
