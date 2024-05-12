package q

// Ellipsis truncates a string to a maximum length and appends an ellipsis ("...")
// if the original string exceeds the maximum length. If the string is shorter
// than or equal to the maximum length, it returns the original string without
// modification.
//
// Parameters:
//   - s: The original string to be truncated.
//   - max: The maximum allowed length of the string.
//
// Returns:
//
//	A string that is either the original string (if its length is less than
//	or equal to max) or a truncated version of the original string with an
//	ellipsis appended (if its length is greater than max).
func Ellipsis(s string, max int) string {
	if len(s) > max {
		return s[:max] + "..."
	}
	return s
}
