package solution

import (
	"reflect"
	"testing"
)

func TestFindContentChildren(t *testing.T) {
	tests := []struct {
		g        []int
		s        []int
		expected int
	}{
		{[]int{1, 2, 3}, []int{1, 1}, 1},
		{[]int{1, 2}, []int{1, 2, 3}, 2},
		{[]int{1, 2, 7, 8}, []int{1, 3, 5, 9}, 3},
		{[]int{1, 2, 3}, []int{3}, 1},
		{[]int{}, []int{1, 2, 3}, 0},
		{[]int{1, 2, 3}, []int{}, 0},
		{[]int{10, 9, 8, 7}, []int{5, 6, 7, 8}, 2},
	}

	for _, test := range tests {
		result := FindContentChildren(test.g, test.s)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("FindContentChildren(%v, %v) = %d; expected %d", test.g, test.s, result, test.expected)
		}
	}
}