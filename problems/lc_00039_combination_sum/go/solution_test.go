package lc_00039_combination_sum

import (
	"reflect"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	tests := []struct {
		name       string
		candidates []int
		target     int
		expected   [][]int
	}{
		{
			name:       "示例1",
			candidates: []int{2, 3, 6, 7},
			target:     7,
			expected:   [][]int{{2, 2, 3}, {7}},
		},
		{
			name:       "示例2",
			candidates: []int{2, 3, 5},
			target:     8,
			expected:   [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		{
			name:       "示例3",
			candidates: []int{2},
			target:     1,
			expected:   [][]int{},
		},
		{
			name:       "无解",
			candidates: []int{1},
			target:     3,
			expected:   [][]int{{1, 1, 1}},
		},
		{
			name:       "空数组",
			candidates: []int{},
			target:     3,
			expected:   [][]int{},
		},
		{
			name:       "目标值为0",
			candidates: []int{1, 2, 3},
			target:     0,
			expected:   [][]int{},
		},
		{
			name:       "完全匹配",
			candidates: []int{1, 2, 3},
			target:     6,
			expected:   [][]int{{1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 2}, {1, 1, 2, 2}, {2, 2, 2}, {1, 1, 1, 3}, {1, 2, 3}, {3, 3}},
		},
		{
			name:       "重复元素",
			candidates: []int{2, 2, 3, 7},
			target:     7,
			expected:   [][]int{{2, 2, 3}, {7}},
		},
		{
			name:       "大量重复元素",
			candidates: []int{1, 2, 3, 4},
			target:     10,
			expected:   [][]int{{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1, 1, 1, 2}, {1, 1, 1, 1, 1, 1, 1, 3}, {1, 1, 1, 1, 1, 1, 2, 2}, {1, 1, 1, 1, 1, 2, 2, 1}, {1, 1, 1, 1, 1, 2, 3}, {1, 1, 1, 1, 1, 3, 2}, {1, 1, 1, 1, 1, 4, 1}, {1, 1, 1, 1, 2, 2, 1, 1}, {1, 1, 1, 1, 2, 2, 2}, {1, 1, 1, 1, 2, 3, 1}, {1, 1, 1, 1, 3, 2, 1}, {1, 1, 1, 1, 4, 1, 1}, {1, 1, 1, 2, 2, 1, 1, 1}, {1, 1, 1, 2, 2, 2, 1}, {1, 1, 1, 2, 3, 1, 1}, {1, 1, 1, 3, 2, 1, 1}, {1, 1, 1, 4, 1, 1, 1}, {1, 1, 2, 2, 1, 1, 1, 1}, {1, 1, 2, 2, 2, 1, 1}, {1, 1, 2, 3, 1, 1, 1}, {1, 1, 3, 2, 1, 1, 1}, {1, 1, 4, 1, 1, 1, 1}, {1, 2, 2, 1, 1, 1, 1, 1}, {1, 2, 2, 2, 1, 1, 1}, {1, 2, 3, 1, 1, 1, 1}, {1, 3, 2, 1, 1, 1, 1}, {1, 4, 1, 1, 1, 1, 1}, {2, 2, 1, 1, 1, 1, 1, 1}, {2, 2, 2, 1, 1, 1, 1}, {2, 2, 3, 1, 1, 1}, {2, 3, 2, 1, 1, 1}, {2, 4, 1, 1, 1, 1}, {3, 2, 1, 1, 1, 1}, {3, 3, 1, 1, 1}, {3, 4, 1, 1}, {4, 2, 1, 1, 1}, {4, 3, 1, 1}, {4, 4, 1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := combinationSum(tt.candidates, tt.target)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("combinationSum(%v, %d) = %v, expected %v", tt.candidates, tt.target, result, tt.expected)
			}
		})
	}
}
