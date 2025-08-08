package solution

import (
	"testing"
)

func TestFourSumCount(t *testing.T) {
	tests := []struct {
		nums1    []int
		nums2    []int
		nums3    []int
		nums4    []int
		expected int
	}{
		{[]int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}, 2},
		{[]int{0}, []int{0}, []int{0}, []int{0}, 1},
		{[]int{}, []int{}, []int{}, []int{}, 0},
		{[]int{-1, -1}, []int{-1, 1}, []int{-1, 1}, []int{1, -1}, 6},
	}

	for _, test := range tests {
		result := FourSumCount(test.nums1, test.nums2, test.nums3, test.nums4)
		if result != test.expected {
			t.Errorf("FourSumCount(%v, %v, %v, %v) = %d; expected %d",
				test.nums1, test.nums2, test.nums3, test.nums4, result, test.expected)
		}
	}
}
