package solution

import (
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {
	tests := []struct {
		input    []byte
		expected []byte
		name     string
	}{
		{
			input:    []byte{'h', 'e', 'l', 'l', 'o'},
			expected: []byte{'o', 'l', 'l', 'e', 'h'},
			name:     "example1",
		},
		{
			input:    []byte{'H'},
			expected: []byte{'H'},
			name:     "single_char",
		},
		{
			input:    []byte{},
			expected: []byte{},
			name:     "empty_string",
		},
		{
			input:    []byte{'a', 'b', 'c', 'd', 'e'},
			expected: []byte{'e', 'd', 'c', 'b', 'a'},
			name:     "odd_length",
		},
		{
			input:    []byte{'a', 'b', 'c', 'd'},
			expected: []byte{'d', 'c', 'b', 'a'},
			name:     "even_length",
		},
		{
			input:    []byte{'!', '@', '#', '\'', '%', '^', '&', '*'},
			expected: []byte{'*', '&', '^', '%', '\'', '#', '@', '!'},
			name:     "special_chars",
		},
		{
			input:    []byte{'A', 'B', 'a', 'b', '1', '2'},
			expected: []byte{'2', '1', 'b', 'a', 'B', 'A'},
			name:     "mixed_chars",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			input := make([]byte, len(test.input))
			copy(input, test.input)
			reverseString(input)
			if !reflect.DeepEqual(input, test.expected) {
				t.Errorf("reverseString(%v) = %v; expected %v", test.input, input, test.expected)
			}
		})
	}
}
