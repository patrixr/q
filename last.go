package q

// Last returns the last element of a slice and a boolean indicating if the slice was non-empty.
// If the slice is empty, it returns the zero value of the element type and false.
func Last[E any](s []E) (E, bool) {
	if len(s) == 0 {
		var zero E
		return zero, false
	}
	return s[len(s)-1], true
}
