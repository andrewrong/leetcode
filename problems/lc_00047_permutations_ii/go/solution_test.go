package lc_00047_permutations_ii

import (
	"reflect"
	"sort"
	"testing"
)

func TestPermuteUnique(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 1, 2},
			expected: [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}},
		},
		{
			name:     "Example 2",
			nums:     []int{1, 2, 1, 1},
			expected: [][]int{{1, 1, 1, 2}, {1, 1, 2, 1}, {1, 2, 1, 1}, {2, 1, 1, 1}},
		},
		{
			name:     "Single Element",
			nums:     []int{1},
			expected: [][]int{{1}},
		},
		{
			name:     "All Same Elements",
			nums:     []int{1, 1, 1},
			expected: [][]int{{1, 1, 1}},
		},
		{
			name:     "Empty Array",
			nums:     []int{},
			expected: [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := permuteUnique(tt.nums)
			sort.Slice(result, func(i, j int) bool {
				for k := range result[i] {
					if result[i][k] != result[j][k] {
						return result[i][k] < result[j][k]
					}
				}
				return false
			})
			sort.Slice(tt.expected, func(i, j int) bool {
				for k := range tt.expected[i] {
					if tt.expected[i][k] != tt.expected[j][k] {
						return tt.expected[i][k] < tt.expected[j][k]
					}
				}
				return false
			})
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("permuteUnique(%v) = %v, want %v", tt.nums, result, tt.expected)
			}
		})
	}
}
