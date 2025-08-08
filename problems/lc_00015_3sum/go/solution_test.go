package lc_00015_3sum

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{
			name:     "示例1",
			nums:     []int{-1, 0, 1, 2, -1, -4},
			expected: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name:     "全零数组",
			nums:     []int{0, 0, 0},
			expected: [][]int{{0, 0, 0}},
		},
		{
			name:     "无解数组",
			nums:     []int{0, 1, 1},
			expected: [][]int{},
		},
		{
			name:     "空数组",
			nums:     []int{},
			expected: [][]int{},
		},
		{
			name:     "单元素数组",
			nums:     []int{0},
			expected: [][]int{},
		},
		{
			name:     "两元素数组",
			nums:     []int{0, 1},
			expected: [][]int{},
		},
		{
			name:     "包含重复元素",
			nums:     []int{-2, 0, 0, 2, 2},
			expected: [][]int{{-2, 0, 2}},
		},
		{
			name:     "全为正数",
			nums:     []int{1, 2, 3},
			expected: [][]int{},
		},
		{
			name:     "全为负数",
			nums:     []int{-1, -2, -3},
			expected: [][]int{},
		},
		{
			name:     "多个重复元素",
			nums:     []int{-1, -1, -1, 0, 0, 0, 1, 1, 1, 2, 2, 2},
			expected: [][]int{{-1, -1, 2}, {-1, 0, 1}, {0, 0, 0}},
		},
		{
			name:     "大量重复零",
			nums:     []int{0, 0, 0, 0, 0, 0},
			expected: [][]int{{0, 0, 0}},
		},
		{
			name:     "重复元素导致多个解",
			nums:     []int{-2, -2, -1, -1, 0, 0, 1, 1, 2, 2},
			expected: [][]int{{-2, 0, 2}, {-2, 1, 1}, {-1, -1, 2}, {-1, 0, 1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := threeSum(tt.nums)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("threeSum(%v) = %v, expected %v", tt.nums, result, tt.expected)
			}
		})
	}
}
