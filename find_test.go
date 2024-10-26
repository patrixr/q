package q

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestFindAll(t *testing.T) {
	// Test with an empty slice
	t.Run("TestEmptySlice", func(t *testing.T) {
		var nums []int
		result := FindAll(nums, func(n int, idx int) bool {
			return n > 0
		})
		assert.Empty(t, result)
	})

	// Test when all elements match the predicate
	t.Run("TestAllMatchPredicate", func(t *testing.T) {
		nums := []int{1, 2, 3, 4}
		result := FindAll(nums, func(n int, idx int) bool {
			return n > 0
		})
		assert.Equal(t, nums, result)
	})

	// Test when no elements match the predicate
	t.Run("TestNoMatchPredicate", func(t *testing.T) {
		nums := []int{1, 2, 3, 4}
		result := FindAll(nums, func(n int, idx int) bool {
			return n > 4
		})
		assert.Empty(t, result)
	})

	// Test when some elements match the predicate
	t.Run("TestSomeMatchPredicate", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		result := FindAll(nums, func(n int, idx int) bool {
			return n%2 == 0 // matches even numbers
		})
		expected := []int{2, 4}
		assert.Equal(t, expected, result)
	})

	// Test with strings
	t.Run("TestTypeSpecific", func(t *testing.T) {
		strs := []string{"apple", "banana", "cherry"}
		result := FindAll(strs, func(s string, idx int) bool {
			return len(s) > 5
		})
		expected := []string{"banana", "cherry"}
		assert.Equal(t, expected, result)
	})
}
