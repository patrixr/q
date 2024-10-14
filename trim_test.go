package q

import (
	"testing"
)

func TestTrimIndent(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "\t\tThis is a test.\n\t\tWith some indentation.\n\t\tAnd another line.",
			expected: "This is a test.\nWith some indentation.\nAnd another line.",
		},
		{
			input:    "    Leading spaces.\n    Another line with spaces.\n    Last line.",
			expected: "Leading spaces.\nAnother line with spaces.\nLast line.",
		},
		{
			input:    "No indentation at all.\nJust normal lines.",
			expected: "No indentation at all.\nJust normal lines.",
		},
		{
			input:    "\n\n\n",
			expected: "\n\n\n", // Expecting the same empty lines
		},
		{
			input:    "\t\t\t   \n\t\tAnother line with tabs.\n\t\t\t\nYet another line.",
			expected: "\n    Another line with tabs.\n\nYet another line.",
		},
		{
			input:    "   Mixed indents.\n\t\tTab indented line.\n    Another mixed line.",
			expected: "Mixed indents.\n Tab indented line.\n Another mixed line.",
		},
		{
			input:    "",
			expected: "",
		},
		{
			input: `
				START
				   INDENT
				END
		    `,
			expected: "\nSTART\n   INDENT\nEND\n",
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := TrimIndent(test.input)
			if result != test.expected {
				t.Errorf("Expected:\n%s\nGot:\n%s", test.expected, result)
			}
		})
	}
}
