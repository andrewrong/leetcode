package solution

import (
	"reflect"
	"testing"
)

func TestSubsets(t *testing.T) {
	tests := []struct {
		nums     []int
		expected [][]int
	}{
		{
			[]int{1, 2, 3},
			[][]int{
				{},
				{1},
				{2},
				{1, 2},
				{3},
				{1, 3},
				{2, 3},
				{1, 2, 3},
			},
		},
		{
			[]int{0},
			[][]int{
				{},
				{0},
			},
		},
		{
			[]int{},
			[][]int{
				{},
			},
		},
		{
			[]int{1, 2, 2},
			[][]int{},
		},
	}

	for _, test := range tests {
		result := Subsets(test.nums)
		if !equalSubsets(result, test.expected) {
			t.Errorf("Subsets(%v) = %v; expected %v", test.nums, result, test.expected)
		}
	}
}

// equalSubsets checks if two slices of slices contain the same subsets (order independent)
func equalSubsets(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	// Convert to maps for easier comparison
	mapA := make(map[string]bool)
	mapB := make(map[string]bool)

	for _, subset := range a {
		// Sort subset for consistent key
		sorted := make([]int, len(subset))
		copy(sorted, subset)
		// Simple sorting for small arrays
		for i := 0; i < len(sorted); i++ {
			for j := i + 1; j < len(sorted); j++ {
				if sorted[i] > sorted[j] {
					sorted[i], sorted[j] = sorted[j], sorted[i]
				}
			}
		}

		key := ""
		for _, v := range sorted {
			key += string(rune(v))
		}
		mapA[key] = true
	}

	for _, subset := range b {
		// Sort subset for consistent key
		sorted := make([]int, len(subset))
		copy(sorted, subset)
		// Simple sorting for small arrays
		for i := 0; i < len(sorted); i++ {
			for j := i + 1; j < len(sorted); j++ {
				if sorted[i] > sorted[j] {
					sorted[i], sorted[j] = sorted[j], sorted[i]
				}
			}
		}

		key := ""
		for _, v := range sorted {
			key += string(rune(v))
		}
		mapB[key] = true
	}

	return reflect.DeepEqual(mapA, mapB)
}
