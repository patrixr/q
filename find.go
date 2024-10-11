package q

func Find[T comparable](array []T, predicate func (T, int) bool) (bool, T) {
	for i, item := range array {
		if predicate(item, i) {
			return true, item
		}
	}
	var empty T
	return false, empty
}
