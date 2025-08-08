package solution

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		s        string
		t        string
		expected bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
		{"", "", true},
		{"abc", "abcd", false},
	}

	for _, test := range tests {
		result := Solve(test.s, test.t)
		if result != test.expected {
			t.Errorf("IsAnagram(%q, %q) = %v; expected %v", test.s, test.t, result, test.expected)
		}
	}
}
