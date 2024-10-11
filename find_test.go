package q

import (
	"testing"
)

func TestFind(t *testing.T) {
	t.Run("FindInt", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5}
		predicate := func(item int, index int) bool {
			return item == 3
		}
		found, item, index := Find(array, predicate)
		if !found || item != 3 || index != 2 {
			t.Errorf("Expected to find 3 at index 2, but got found=%v, item=%v, index=%v", found, item, index)
		}
	})

	t.Run("FindString", func(t *testing.T) {
		array := []string{"apple", "banana", "cherry"}
		predicate := func(item string, index int) bool {
			return item == "banana"
		}
		found, item, index := Find(array, predicate)
		if !found || item != "banana" || index != 1 {
			t.Errorf("Expected to find 'banana' at index 1, but got found=%v, item=%v, index=%v", found, item, index)
		}
	})

	t.Run("NotFound", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5}
		predicate := func(item int, index int) bool {
			return item == 6
		}
		found, item, index := Find(array, predicate)
		if found || item != 0 || index != -1 {
			t.Errorf("Expected not to find any item, but got found=%v, item=%v, index=%v", found, item, index)
		}
	})

	t.Run("EmptyArray", func(t *testing.T) {
		array := []int{}
		predicate := func(item int, index int) bool {
			return item == 1
		}
		found, item, index := Find(array, predicate)
		if found || item != 0 || index != -1 {
			t.Errorf("Expected not to find any item in empty array, but got found=%v, item=%v, index=%v", found, item, index)
		}
	})

	t.Run("FindWithIndex", func(t *testing.T) {
		array := []int{10, 20, 30, 40}
		predicate := func(item int, index int) bool {
			return index == 2 // Looking for the item at index 2
		}
		found, item, index := Find(array, predicate)
		if !found || item != 30 || index != 2 {
			t.Errorf("Expected to find item at index 2 (30), but got found=%v, item=%v, index=%v", found, item, index)
		}
	})
}
