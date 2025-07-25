package lc_00059_spiral_matrix_2

import (
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	testCases := []struct {
		name     string
		n        int
		expected [][]int
	}{
		{
			name: "n=1",
			n:    1,
			expected: [][]int{
				{1},
			},
		},
		{
			name: "n=2",
			n:    2,
			expected: [][]int{
				{1, 2},
				{4, 3},
			},
		},
		{
			name: "n=3",
			n:    3,
			expected: [][]int{
				{1, 2, 3},
				{8, 9, 4},
				{7, 6, 5},
			},
		},
		{
			name: "n=4",
			n:    4,
			expected: [][]int{
				{1, 2, 3, 4},
				{12, 13, 14, 5},
				{11, 16, 15, 6},
				{10, 9, 8, 7},
			},
		},
		{
			name: "n=5",
			n:    5,
			expected: [][]int{
				{1, 2, 3, 4, 5},
				{16, 17, 18, 19, 6},
				{15, 24, 25, 20, 7},
				{14, 23, 22, 21, 8},
				{13, 12, 11, 10, 9},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Solve(tc.n)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Solve(%d) = %v; expected %v", tc.n, result, tc.expected)
			}
		})
	}
}