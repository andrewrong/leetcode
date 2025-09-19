package lc_00040_combination_sum_ii

import (
	"reflect"
	"testing"
)

func TestCombinationSum2(t *testing.T) {
	tests := []struct {
		name     string
		candidates []int
		target   int
		expected [][]int
	}{
		{
			name:     "示例1",
			candidates: []int{10, 1, 2, 7, 6, 1, 5},
			target:   8,
			expected: [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}},
		},
		{
			name:     "示例2",
			candidates: []int{2, 5, 2, 1, 2},
			target:   5,
			expected: [][]int{{1, 2, 2}, {5}},
		},
		{
			name:     "有解",
			candidates: []int{2, 3, 6, 7},
			target:   9,
			expected: [][]int{{2, 7}, {3, 6}},
		},
		{
			name:     "空数组",
			candidates: []int{},
			target:   3,
			expected: [][]int{},
		},
		{
			name:     "单元素匹配",
			candidates: []int{3},
			target:   3,
			expected: [][]int{{3}},
		},
		{
			name:     "重复元素",
			candidates: []int{1, 1, 1, 2, 2},
			target:   3,
			expected: [][]int{{1, 1, 1}, {1, 2}},
		},
		{
			name:     "目标值为0",
			candidates: []int{1, 2, 3},
			target:   0,
			expected: [][]int{},
		},
		{
			name:     "包含重复元素的复杂情况",
			candidates: []int{1, 1, 1, 1, 2, 2, 3},
			target:   4,
			expected: [][]int{{1, 1, 1, 1}, {1, 1, 2}, {1, 3}, {2, 2}},
		},
		{
			name:     "全部元素都参与",
			candidates: []int{1, 2, 3, 4},
			target:   10,
			expected: [][]int{{1, 2, 3, 4}},
		},
		{
			name:     "目标值为1",
			candidates: []int{1, 2, 3},
			target:   1,
			expected: [][]int{{1}},
		},
		{
			name:     "无解情况2",
			candidates: []int{1, 2, 3},
			target:   7,
			expected: [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := combinationSum2(tt.candidates, tt.target)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("combinationSum2(%v, %d) = %v, expected %v", tt.candidates, tt.target, result, tt.expected)
			}
		})
	}
}