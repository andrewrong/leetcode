package solution

import (
	"reflect"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	tests := []struct {
		name     string
		digits   string
		expected []string
	}{
		{
			name:     "Example 1",
			digits:   "23",
			expected: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			name:     "Empty string",
			digits:   "",
			expected: []string{},
		},
		{
			name:     "Single digit",
			digits:   "2",
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "Two digits with 4 letters",
			digits:   "79",
			expected: []string{"pw", "px", "py", "pz", "qw", "qx", "qy", "qz", "rw", "rx", "ry", "rz", "sw", "sx", "sy", "sz"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := letterCombinations(tt.digits)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("letterCombinations(%v) = %v, want %v", tt.digits, result, tt.expected)
			}
		})
	}
}