package lc_00491_non_decreasing_subsequences

import (
	"reflect"
	"testing"
)

func TestFindSubsequences(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{
			name: "Example 1",
			nums: []int{4, 6, 7, 7},
			expected: [][]int{
				{4, 6}, {4, 6, 7}, {4, 6, 7, 7}, {4, 7}, {4, 7, 7},
				{6, 7}, {6, 7, 7}, {7, 7},
			},
		},
		{
			name:     "Example 2",
			nums:     []int{4, 4, 3, 2, 1},
			expected: [][]int{{4, 4}},
		},
		{
			name:     "Empty slice",
			nums:     []int{},
			expected: [][]int{},
		},
		{
			name:     "Single element",
			nums:     []int{1},
			expected: [][]int{},
		},
		{
			name:     "Two same elements",
			nums:     []int{1, 1},
			expected: [][]int{{1, 1}},
		},
		{
			name:     "Decreasing sequence",
			nums:     []int{5, 4, 3, 2, 1},
			expected: [][]int{},
		},
		{
			name: "Mixed sequence",
			nums: []int{1, 2, 1, 2},
			expected: [][]int{
				{1, 2}, {1, 2, 2}, {1, 1}, {1, 1, 2},
				{2, 2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findSubsequences(tt.nums)
			if !equal(result, tt.expected) {
				t.Errorf("findSubsequences(%v) = %v, want %v", tt.nums, result, tt.expected)
			}
		})
	}
}

// equal checks if two slices of slices are equal (ignoring order)
func equal(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	// Convert to map for comparison
	aMap := make(map[string]int)
	bMap := make(map[string]int)

	for _, sub := range a {
		key := toString(sub)
		aMap[key]++
	}

	for _, sub := range b {
		key := toString(sub)
		bMap[key]++
	}

	return reflect.DeepEqual(aMap, bMap)
}

// toString converts a slice to a string for map key
func toString(s []int) string {
	result := ""
	for _, v := range s {
		result += string(rune(v+'0')) + ","
	}
	return result
}
