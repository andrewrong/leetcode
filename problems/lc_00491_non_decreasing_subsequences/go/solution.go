package lc_00491_non_decreasing_subsequences

// findSubsequences returns all the different possible non-decreasing subsequences
// of the given array with at least two elements.

func findSubsequences(nums []int) [][]int {
	one := make([]int, 0)
	all := make([][]int, 0)
	// used := make([]bool, len(nums))
	backtracker(nums, 0, &one, &all)

	return all
}

func backtracker(nums []int, startIdx int, one *[]int, all *[][]int) {
	if len(*one) > 1 {
		cp := make([]int, len(*one))
		copy(cp, *one)
		*all = append(*all, cp)
	}

	if startIdx >= len(nums) {
		return
	}

	set := make(map[int]struct{})
	for i := startIdx; i < len(nums); i += 1 {
		if _, ok := set[nums[i]]; ok {
			continue
		}
		set[nums[i]] = struct{}{}

		if len(*one) == 0 {
			*one = append(*one, nums[i])
		} else {
			if nums[i] < (*one)[len(*one)-1] {
				continue
			}
			*one = append(*one, nums[i])
		}
		backtracker(nums, i+1, one, all)
		*one = (*one)[:len(*one)-1]
	}
}
