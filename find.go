package q

import "regexp"

// Find searches for an element in the array that satisfies the predicate function.
// It returns a boolean indicating whether such an element was found, the element itself,
// and the index of the element in the array. If no element is found, it returns false,
// the zero value of the element type, and -1.
//
// The function is generic and works with any type T that is comparable.
//
// Parameters:
//   - array: A slice of elements of type T to search through.
//   - predicate: A function that takes an element of type T and its index, and returns
//     a boolean indicating whether the element satisfies the condition.
//
// Returns:
// - bool: True if an element satisfying the predicate is found, otherwise false.
// - T: The element that satisfies the predicate, or the zero value of type T if not found.
// - int: The index of the element in the array, or -1 if not found.
func Find[T comparable](array []T, predicate func(T, int) bool) (bool, T, int) {
	for i, item := range array {
		if predicate(item, i) {
			return true, item, i
		}
	}
	var empty T
	return false, empty, -1
}

// Eq returns a predicate function that checks if an element is equal to the specified value.
//
// Parameters:
//   - val: The value to compare each element against.
//
// Returns:
//   - func(T, int) bool: A predicate function that takes an element of type T and its index,
//     and returns true if the element is equal to val, otherwise false.
//
// Example usage with Find:
//
//	array := []int{1, 2, 3, 4, 5}
//	found, element, index := q.Find(array, q.Eq(3))
//	// found: true, element: 3, index: 2
func Eq[T comparable](val T) func(T, int) bool {
	return func(it T, _ int) bool {
		return it == val
	}
}

// Match returns a predicate function that checks if an element matches the specified regular expression.
// The function returns true if the element matches the regular expression, otherwise false.
func Match(rexp string) func(string, int) bool {
	r := regexp.MustCompile(rexp)
	return func(it string, _ int) bool {
		return r.MatchString(it)
	}
}
