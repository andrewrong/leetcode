package solution

import (
	"testing"
)

func TestWiggleMaxLength(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 7, 4, 9, 2, 5}, 6},                   // -6,3,-5,7,-3           // All elements form a wiggle sequence
		{[]int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}, 7}, // Expected wiggle sequence
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 2},          // Strictly increasing sequence
		{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}, 2},          // Strictly decreasing sequence
		{[]int{1}, 1},                                  // Single element
		{[]int{1, 1, 1}, 1},                            // All elements same
		{[]int{}, 0},                                   // Empty array
		{[]int{1, 2}, 2},                               // Two different elements
		{[]int{2, 1}, 2},                               // Two different elements (decreasing)
		{[]int{1, 1, 2, 2, 3, 3}, 3},                   // Array with duplicates
	}

	for _, test := range tests {
		result := WiggleMaxLength(test.nums)
		if result != test.expected {
			t.Errorf("WiggleMaxLength(%v) = %d; expected %d", test.nums, result, test.expected)
		}
	}
}
