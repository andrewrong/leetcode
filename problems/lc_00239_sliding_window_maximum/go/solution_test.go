// leetcode/problems/lc_00239_sliding_window_maximum/go/solution_test.go
package main

import (
	"reflect"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	testCases := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{
			name: "Example 1",
			nums: []int{1, 3, -1, -3, 5, 3, 6, 7},
			k:    3,
			want: []int{3, 3, 5, 5, 6, 7},
		},
		{
			name: "Test case with k = 1",
			nums: []int{1, 2, 3, 4, 5},
			k:    1,
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Test case with k = size of nums",
			nums: []int{5, 4, 3, 2, 1},
			k:    5,
			want: []int{5},
		},
		{
			name: "Test case with decreasing numbers",
			nums: []int{9, 8, 7, 6, 5},
			k:    2,
			want: []int{9, 8, 7, 6},
		},
		{
			name: "Test case with all same numbers",
			nums: []int{4, 4, 4, 4, 4},
			k:    4,
			want: []int{4, 4},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := maxSlidingWindow(tc.nums, tc.k)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("maxSlidingWindow(%v, %d) = %v, want %v", tc.nums, tc.k, got, tc.want)
			}
		})
	}
}
