package q

// Filter returns a new slice containing all elements of a for which the function f returns true.
// It iterates over each element of the input slice, applying the function f to it. If f returns true,
// the element is included in the resulting slice. The order of elements in the input slice is preserved
// in the output slice.
//
// This function is generic and can operate on slices of any type, as indicated by the type parameter T.
// The function f takes a single argument of the same type as the elements of the slice and returns a bool.
//
// Parameters:
//
//	a []T: The input slice containing elements of any type T.
//	f func(T) bool: A function that takes an element of type T and returns a bool. Elements for which f returns
//	true are included in the resulting slice.
//
// Returns:
//
//	[]T: A new slice containing all elements of the input slice for which the function f returns true.
//
// Example:
//
//	isEven := func(n int) bool { return n%2 == 0 }
//	numbers := []int{1, 2, 3, 4, 5}
//	evenNumbers := Filter(numbers, isEven)
//	fmt.Println(evenNumbers) // Output: [2 4]
func Filter[T any](a []T, f func(T) bool) []T {
	n := make([]T, 0, len(a))
	for _, e := range a {
		if f(e) {
			n = append(n, e)
		}
	}
	return n
}
