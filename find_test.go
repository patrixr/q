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
		found, item := Find(array, predicate)
		if !found || item != 3 {
			t.Errorf("Expected to find 3, but got found=%v, item=%v", found, item)
		}
	})

	t.Run("FindString", func(t *testing.T) {
		array := []string{"apple", "banana", "cherry"}
		predicate := func(item string, index int) bool {
			return item == "banana"
		}
		found, item := Find(array, predicate)
		if !found || item != "banana" {
			t.Errorf("Expected to find 'banana', but got found=%v, item=%v", found, item)
		}
	})

	t.Run("NotFound", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5}
		predicate := func(item int, index int) bool {
			return item == 6
		}
		found, item := Find(array, predicate)
		if found || item != 0 {
			t.Errorf("Expected not to find any item, but got found=%v, item=%v", found, item)
		}
	})

	t.Run("EmptyArray", func(t *testing.T) {
		array := []int{}
		predicate := func(item int, index int) bool {
			return item == 1
		}
		found, item := Find(array, predicate)
		if found || item != 0 {
			t.Errorf("Expected not to find any item in empty array, but got found=%v, item=%v", found, item)
		}
	})

	t.Run("FindWithIndex", func(t *testing.T) {
		array := []int{10, 20, 30, 40}
		predicate := func(item int, index int) bool {
			return index == 2 // Looking for the item at index 2
		}
		found, item := Find(array, predicate)
		if !found || item != 30 {
			t.Errorf("Expected to find item at index 2 (30), but got found=%v, item=%v", found, item)
		}
	})
}
