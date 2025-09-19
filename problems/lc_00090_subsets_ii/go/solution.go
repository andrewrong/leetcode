package solution

import (
	"sort"
)

// subsetsWithDup returns all possible subsets of the given integer array that might contain duplicates.
// It ensures that the solution set does not contain duplicate subsets.
func subsetsWithDup(nums []int) [][]int {
	// Sort the array to handle duplicates properly
	sort.Ints(nums)

	one := make([]int, 0)
	used := make([]bool, len(nums))
	all := make([][]int, 0)
	backtrack(nums, &one, &all, used)

	return all
}

func backtrack(nums []int, one *[]int, all *[][]int, used []bool) {
	cp := make([]int, len(*one))
	copy(cp, *one)
	*all = append(*all, cp)

	if len(nums) == 0 {
		return
	}

	for i := 0; i < len(nums); i += 1 {
		if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
			continue
		}

		*one = append(*one, nums[i])
		used[i] = true
		backtrack(nums[i+1:], one, all, used[i+1:])
		used[i] = false
		*one = (*one)[:len(*one)-1]
	}
}
