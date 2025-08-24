// leetcode/problems/lc_00150_evaluate_reverse_polish_notation/go/solution_test.go
package main

import "testing"

func TestEvalRPN(t *testing.T) {
	testCases := []struct {
		name   string
		tokens []string
		want   int
	}{
		{
			name:   "Example 1",
			tokens: []string{"2", "1", "+", "3", "*"},
			want:   9,
		},
		{
			name:   "Example 2",
			tokens: []string{"4", "13", "5", "/", "+"},
			want:   6,
		},
		{
			name:   "Example 3",
			tokens: []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			want:   22,
		},
		{
			name:   "Test with negative numbers",
			tokens: []string{"-1", "1", "*", "-1", "+"},
			want:   -2,
		},
		{
			name:   "Test with subtraction",
			tokens: []string{"5", "2", "-"},
			want:   3,
		},
		{
			name:   "Test with single number",
			tokens: []string{"42"},
			want:   42,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := evalRPN(tc.tokens)
			if got != tc.want {
				t.Errorf("evalRPN(%v) = %v, want %v", tc.tokens, got, tc.want)
			}
		})
	}
}
