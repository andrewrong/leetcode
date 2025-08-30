// leetcode/problems/lc_00347_top_k_frequent_elements/go/solution_test.go
package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{
			name: "Example 1",
			nums: []int{1, 1, 1, 2, 2, 3},
			k:    2,
			want: []int{1, 2},
		},
		{
			name: "Example 2",
			nums: []int{1},
			k:    1,
			want: []int{1},
		},
		{
			name: "Example 3",
			nums: []int{4, 1, -1, 2, -1, 2, 3},
			k:    2,
			want: []int{-1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := topKFrequent(tt.nums, tt.k)
			sort.Ints(got)
			sort.Ints(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topKFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}
