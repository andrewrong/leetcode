package solution

import (
	"reflect"
	"testing"
)

func TestIntersection(t *testing.T) {
	tests := []struct {
		nums1    []int
		nums2    []int
		expected []int
	}{
		{[]int{1, 2, 2, 1}, []int{2, 2}, []int{2}},
		{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{9, 4}},
		{[]int{}, []int{}, []int{}},
		{[]int{1, 2, 3}, []int{4, 5, 6}, []int{}},
	}

	for _, test := range tests {
		result := Intersection(test.nums1, test.nums2)
		
		// Since order doesn't matter, we need to compare as sets
		expectedMap := make(map[int]bool)
		resultMap := make(map[int]bool)
		
		for _, val := range test.expected {
			expectedMap[val] = true
		}
		
		for _, val := range result {
			resultMap[val] = true
		}
		
		if !reflect.DeepEqual(expectedMap, resultMap) {
			t.Errorf("Intersection(%v, %v) = %v; expected %v", test.nums1, test.nums2, result, test.expected)
		}
	}
}