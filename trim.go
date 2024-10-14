package q

import (
	"fmt"
	"strings"
	"unicode"
)

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
	lines := strings.Split(strings.ReplaceAll(s, "\t", "  "), "\n")

	if len(lines) == 0 {
		return s
	}

	// Determine the minimum indent
	leftmostIndent := -1

	trimmedLines := Map(lines, func(l string) indentRef {
		ref := processLineIndent(l)

		if len(ref.line) == 0 {
			return ref
		}

		if leftmostIndent < 0 || ref.indent < leftmostIndent {
			leftmostIndent = ref.indent
		}

		return ref
	})

	var builder strings.Builder

	fmt.Println("leftmostIndent", leftmostIndent)

	if leftmostIndent < 0 {
		leftmostIndent = 0
	}

	for i, ref := range trimmedLines {
		fmt.Println(ref.indent, ref.line)

		if len(ref.line) > 0 {
			builder.WriteString(strings.Repeat(" ", ref.indent-leftmostIndent))
			builder.WriteString(ref.line)
		}

		if i < len(trimmedLines)-1 {
			builder.WriteRune('\n')
		}
	}

	return builder.String()
}

// Alias for TrimIndent
func Paragraph(text string) string {
	return TrimIndent(text)
}

type indentRef struct {
	line   string
	indent int
}

func processLineIndent(line string) indentRef {
	trimmed := strings.TrimSpace(line)

	if trimmed == "" {
		return indentRef{"", 0}
	}

	count := 0
	for _, r := range line {
		if !unicode.IsSpace(r) {
			break
		}
		if r == '\t' {
			count += 2
		} else {
			count++
		}
	}
	return indentRef{trimmed, count}
}
