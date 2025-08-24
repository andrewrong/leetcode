// leetcode/problems/lc_00459_repeated_substring_pattern/go/solution_test.go
package main

import "testing"

func TestRepeatedSubstringPattern(t *testing.T) {
	testCases := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "Basic true case",
			s:    "abab",
			want: true,
		},
		{
			name: "Basic false case",
			s:    "aba",
			want: false,
		},
		{
			name: "Longer true case",
			s:    "abcabcabcabc",
			want: true,
		},
		{
			name: "Single character repeated",
			s:    "aaaa",
			want: true,
		},
		{
			name: "Single character not repeated",
			s:    "a",
			want: false,
		},
		{
			name: "Complex false case",
			s:    "abac",
			want: false,
		},
		{
			name: "True case with longer substring",
			s:    "ababab",
			want: true,
		},
		{
			name: "False case with similar pattern",
			s:    "ababac",
			want: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := repeatedSubstringPattern(tc.s)
			if got != tc.want {
				t.Errorf("repeatedSubstringPattern(%q) = %v, want %v", tc.s, got, tc.want)
			}
		})
	}
}
