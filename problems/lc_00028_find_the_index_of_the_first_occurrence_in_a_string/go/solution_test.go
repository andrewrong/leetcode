// leetcode/problems/lc_00028_find_the_index_of_the_first_occurrence_in_a_string/go/solution_test.go
package main

import "testing"

func TestStrStr(t *testing.T) {
	testCases := []struct {
		name     string
		haystack string
		needle   string
		want     int
	}{
		{
			name:     "Basic case, needle is in the middle",
			haystack: "sadbutsad",
			needle:   "sad",
			want:     0,
		},
		{
			name:     "Needle is not in haystack",
			haystack: "leetcode",
			needle:   "leeto",
			want:     -1,
		},
		{
			name:     "Needle is at the beginning",
			haystack: "hello",
			needle:   "he",
			want:     0,
		},
		{
			name:     "Needle is at the end",
			haystack: "hello",
			needle:   "lo",
			want:     3,
		},
		{
			name:     "Haystack and needle are the same",
			haystack: "a",
			needle:   "a",
			want:     0,
		},
		{
			name:     "Needle is an empty string",
			haystack: "abc",
			needle:   "",
			want:     0,
		},
		{
			name:     "Haystack is an empty string and needle is not",
			haystack: "",
			needle:   "a",
			want:     -1,
		},
		{
			name:     "Both are empty strings",
			haystack: "",
			needle:   "",
			want:     0,
		},
		{
			name:     "Needle is longer than haystack",
			haystack: "a",
			needle:   "aa",
			want:     -1,
		},
		{
			name:     "Complex case with repeated characters",
			haystack: "mississippi",
			needle:   "issip",
			want:     4,
		},
		{
			name:     "Another complex case",
			haystack: "ababcabcabababd",
			needle:   "ababd",
			want:     10,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := strStr(tc.haystack, tc.needle)
			if got != tc.want {
				t.Errorf("strStr(%q, %q) = %v, want %v", tc.haystack, tc.needle, got, tc.want)
			}
		})
	}
}
