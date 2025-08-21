package solution

import (
	"testing"
)

func TestReverseStr(t *testing.T) {
	tests := []struct {
		input    string
		k        int
		expected string
		name     string
	}{
		{
			input:    "abcdefg",
			k:        2,
			expected: "bacdfeg",
			name:     "example1",
		},
		{
			input:    "abcd",
			k:        2,
			expected: "bacd",
			name:     "example2",
		},
		{
			input:    "abcdefgh",
			k:        9, // k greater than string length
			expected: "hgfedcba",
			name:     "k_greater_than_length",
		},
		{
			input:    "abcdefghijk",
			k:        2,
			expected: "bacdfeghjik",
			name:     "exact_multiple",
		},
		{
			input:    "",
			k:        2,
			expected: "",
			name:     "empty_string",
		},
		{
			input:    "a",
			k:        2,
			expected: "a",
			name:     "single_char",
		},
		{
			input:    "abcdefghij",
			k:        3,
			expected: "cbadefihgj",
			name:     "multiple_groups",
		},
		{
			input:    "abc",
			k:        4, // k greater than string length
			expected: "cba",
			name:     "k_greater_than_string_length",
		},
		{
			input:    "abcd",
			k:        4, // k equals string length
			expected: "dcba",
			name:     "k_equals_string_length",
		},
		{
			input:    "abcde",
			k:        4,
			expected: "dcbae",
			name:     "k_less_than_string_length",
		},
		{
			input:    "hyzqyljrnigxvdtneasepfahmtyhlohwxmkqcdfehybknvdmfrfvtbsovjbdhevlfxpdaovjgunjqlimjkfnqcqnajmebeddqsgl",
			k:        39,
			expected: "fdcqkmxwholhytmhafpesaentdvxginrjlyqzyhehybknvdmfrfvtbsovjbdhevlfxpdaovjgunjqllgsqddebemjanqcqnfkjmi",
			name:     "long_string",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := reverseStr(test.input, test.k)
			if result != test.expected {
				t.Errorf("reverseStr(%q, %d) = %q; expected %q", test.input, test.k, result, test.expected)
			}
		})
	}
}
