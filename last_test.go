package q

import (
	"testing"
)

func TestLast(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		s := []int{}
		got, ok := Last(s)
		if ok {
			t.Errorf("Last(%v) = %v, %v; want _, false", s, got, ok)
		}
	})
	t.Run("non-empty", func(t *testing.T) {
		s := []int{1, 2, 3}
		got, ok := Last(s)
		if !ok {
			t.Errorf("Last(%v) = %v, %v; want _, true", s, got, ok)
		}
	})
}
