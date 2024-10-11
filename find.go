package q

func Find[T comparable](array []T, predicate func (T, int) bool) (bool, T, int) {
	for i, item := range array {
		if predicate(item, i) {
			return true, item, i
		}
	}
	var empty T
	return false, empty, -1
}
