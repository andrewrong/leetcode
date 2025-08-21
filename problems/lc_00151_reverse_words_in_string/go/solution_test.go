package solution

import (
	"testing"
)

func TestReverseWords(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		name     string
	}{
		{
			input:    "the sky is blue",
			expected: "blue is sky the",
			name:     "basic",
		},
		{
			input:    "  hello world  ",
			expected: "world hello",
			name:     "extra_spaces",
		},
		{
			input:    "a good   example",
			expected: "example good a",
			name:     "multiple_spaces",
		},
		{
			input:    "word",
			expected: "word",
			name:     "single_word",
		},
		{
			input:    "",
			expected: "",
			name:     "empty_string",
		},
		{
			input:    "  ",
			expected: "",
			name:     "only_spaces",
		},
		{
			input:    "a",
			expected: "a",
			name:     "single_char",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := reverseWords(test.input)
			if result != test.expected {
				t.Errorf("reverseWords(%q) = %q; expected %q", test.input, result, test.expected)
			}
		})
	}
}