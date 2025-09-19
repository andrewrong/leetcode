package solution

import (
	"reflect"
	"sort"
	"testing"
)

// Helper function to check if two slices of slices are equal (ignoring order)
func equalIgnoreOrder(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	// Sort each slice for comparison
	sort.Slice(a, func(i, j int) bool {
		for k := 0; k < len(a[i]) && k < len(a[j]); k++ {
			if a[i][k] != a[j][k] {
				return a[i][k] < a[j][k]
			}
		}
		return len(a[i]) < len(a[j])
	})

	sort.Slice(b, func(i, j int) bool {
		for k := 0; k < len(b[i]) && k < len(b[j]); k++ {
			if b[i][k] != b[j][k] {
				return b[i][k] < b[j][k]
			}
		}
		return len(b[i]) < len(b[j])
	})

	// Compare element by element
	for i := range a {
		if !reflect.DeepEqual(a[i], b[i]) {
			return false
		}
	}

	return true
}

func TestSubsetsWithDup(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected [][]int
	}{
		{
			name:  "Example 1",
			input: []int{1, 2, 2},
			expected: [][]int{
				{},
				{1},
				{1, 2},
				{1, 2, 2},
				{2},
				{2, 2},
			},
		},
		{
			name:  "Example 2",
			input: []int{0},
			expected: [][]int{
				{},
				{0},
			},
		},
		{
			name:  "Empty array",
			input: []int{},
			expected: [][]int{
				{},
			},
		},
		{
			name:  "All duplicates",
			input: []int{1, 1, 1},
			expected: [][]int{
				{},
				{1},
				{1, 1},
				{1, 1, 1},
			},
		},
		{
			name:  "No duplicates",
			input: []int{1, 2, 3},
			expected: [][]int{
				{},
				{1},
				{1, 2},
				{1, 2, 3},
				{1, 3},
				{2},
				{2, 3},
				{3},
			},
		},
		{
			name:  "Negative numbers with duplicates",
			input: []int{-1, 1, 2, 2},
			expected: [][]int{
				{},
				{-1},
				{-1, 1},
				{-1, 1, 2},
				{-1, 1, 2, 2},
				{-1, 2},
				{-1, 2, 2},
				{1},
				{1, 2},
				{1, 2, 2},
				{2},
				{2, 2},
			},
		},
		{
			name:     "xxx",
			input:    []int{4, 4, 4, 1, 4},
			expected: [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := subsetsWithDup(tt.input)
			if !equalIgnoreOrder(result, tt.expected) {
				t.Errorf("subsetsWithDup(%v) = %v, expected %v", tt.input, result, tt.expected)
			}
		})
	}
}
