// leetcode/problems/lc_00020_valid_parentheses/go/solution_test.go
package main

import "testing"

func TestIsValid(t *testing.T) {
	testCases := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "Basic valid case",
			s:    "()",
			want: true,
		},
		{
			name: "All types valid",
			s:    "()[]{}",
			want: true,
		},
		{
			name: "Basic invalid case",
			s:    "(]",
			want: false,
		},
		{
			name: "Nested valid case",
			s:    "{[]}",
			want: true,
		},
		{
			name: "Nested invalid case",
			s:    "([)]",
			want: false,
		},
		{
			name: "Starts with closing bracket",
			s:    "]",
			want: false,
		},
		{
			name: "Odd number of brackets",
			s:    "([)",
			want: false,
		},
		{
			name: "Empty string is valid",
			s:    "",
			want: true,
		},
		{
			name: "Long valid string",
			s:    "(([]){})",
			want: true,
		},
		{
			name: "Long invalid string",
			s:    "(([]){})(",
			want: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := isValid(tc.s)
			if got != tc.want {
				t.Errorf("isValid(%q) = %v, want %v", tc.s, got, tc.want)
			}
		})
	}
}
