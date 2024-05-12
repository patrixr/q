package q

// Map applies a function to each element of a slice, returning a new slice of the results.
// The function f is applied to each element of the input slice a, and a new slice of the resulting type M is returned.
// This allows for transforming slices of one type into slices of another type through the provided function.
//
// Parameters:
// a - The input slice of type []T.
// f - The function to apply to each element of a. The function must take a value of type T and return a value of type M.
//
// Returns:
// A new slice of type []M containing the results of applying f to each element of a.
func Map[T any, M any](a []T, f func(T) M) []M {
	n := make([]M, len(a))
	for i, e := range a {
		n[i] = f(e)
	}
	return n
}
