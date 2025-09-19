package lc_00046_permutations

// permute returns all possible permutations of the given array nums.
func permute(nums []int) [][]int {
	one := make([]int, 0)
	all := make([][]int, 0)
	used := make([]bool, len(nums))

	back(nums, &one, &all, used)
	return all
}

func back(nums []int, one *[]int, all *[][]int, used []bool) {
	if len(nums) == 0 {
		return
	}

	if len(*one) == len(nums) {
		cp := make([]int, len(*one))
		copy(cp, *one)
		*all = append(*all, cp)
		return
	}

	for i := 0; i < len(nums); i += 1 {
		if used[i] {
			continue
		}
		*one = append(*one, nums[i])
		used[i] = true
		back(nums, one, all, used)
		used[i] = false
		*one = (*one)[:len(*one)-1]
	}
}
