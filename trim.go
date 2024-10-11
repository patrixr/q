package q

import "strings"

// TrimIndent removes the common leading whitespace from each line in a given string.
// It first determines the minimum indentation level across all non-empty lines,
// then removes that amount of leading whitespace from each line and trims any
// additional leading or trailing whitespace.
//
// Parameters:
//
//	s (string): The input string with potential leading indentation.
//
// Returns:
//
//	string: A new string with the common leading indentation removed from each line.
//
// Example:
//
//	input := "    line1\n    line2\n    line3"
//	output := TrimIndent(input)
//	// output will be:
//	// "line1\nline2\nline3"
func TrimIndent(s string) string {
	lines := strings.Split(s, "\n")

	if len(lines) == 0 {
		return s
	}

	// Determine the minimum indent
	minIndent := len(lines[0])
	for _, line := range lines[1:] {
		if len(line) == 0 {
			continue
		}
		indent := len(line) - len(strings.TrimLeft(line, " \t"))
		if indent < minIndent {
			minIndent = indent
		}
	}

	// Remove the minimum indent and trim whitespace from each line
	for i := range lines {
		lines[i] = strings.TrimSpace(lines[i][minIndent:])
	}

	return strings.Join(lines, "\n")
}

// Alias for TrimIndent
func Paragraph(text string) string {
	return TrimIndent(text)
}
