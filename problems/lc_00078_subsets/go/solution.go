package solution

import "sort"

// Subsets returns all possible subsets (the power set) of the given array
func innerSubset(nums []int) [][]int {
	result := make([][]int, 0)

	if len(nums) == 0 {
		result = append(result, []int{})
		return result
	}

	if len(nums) == 1 {
		result = append(result, []int{nums[0]})
		result = append(result, []int{})
		return result
	}

	tmp := innerSubset(nums[1:])
	for _, v := range tmp {
		result = append(result, append([]int{nums[0]}, v...))
	}
	result = append(result, tmp...)

	return result
}

func Subsets(nums []int) [][]int {
	sort.Ints(nums)
	one := make([]int, 0)
	all := make([][]int, 0)
	innerSubset2(nums, &one, &all)

	return all
}

func innerSubset2(nums []int, one *[]int, all *[][]int) {
	cp := make([]int, len(*one))
	copy(cp, *one)
	*all = append(*all, cp)

	if len(nums) == 0 {
		return
	}

	for i := 0; i < len(nums); i += 1 {
		*one = append(*one, nums[i])
		innerSubset2(nums[i+1:], one, all)
		*one = (*one)[:len(*one)-1]
	}
}
