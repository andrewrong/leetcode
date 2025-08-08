package solution

import (
	"testing"
)

func TestFindMinArrowShots(t *testing.T) {
	tests := []struct {
		points   [][]int
		expected int
	}{
		{[][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}, 2},
		{[][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}, 4},
		{[][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}, 2},
		{[][]int{}, 0},
	}

	for _, test := range tests {
		result := FindMinArrowShots(test.points)
		if result != test.expected {
			t.Errorf("FindMinArrowShots(%v) = %d; expected %d", test.points, result, test.expected)
		}
	}
}