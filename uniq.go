package q

// Uniq removes duplicates from a slice of any comparable type and returns a slice of unique elements.
// It preserves the order of the first occurrence of each element.
//
// Parameters:
// - inputSlice: A slice of any comparable type (e.g., int, string) from which duplicates need to be removed.
//
// Returns:
// A new slice containing only the unique elements from the input slice, in the order of their first occurrence.
//
// Example:
//
//	numbers := []int{1, 2, 2, 3, 4, 4, 5}
//	uniqueNumbers := Uniq(numbers)
//	fmt.Println(uniqueNumbers) // Output: [1, 2, 3, 4, 5]
func Uniq[T comparable](inputSlice []T) []T {
	uniqueSlice := make([]T, 0, len(inputSlice))
	seen := make(map[T]bool, len(inputSlice))
	for _, element := range inputSlice {
		if !seen[element] {
			uniqueSlice = append(uniqueSlice, element)
			seen[element] = true
		}
	}
	return uniqueSlice
}
