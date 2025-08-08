package lc_00015_3sum

import "sort"

// 使用双指针法, 三个数字求和，本身比二数求和要麻烦很多，而且题目要求的是结果不要重复；所以整个解题的方式准备使用双指针的解决方式:
// 整体的思路是：外层进行一层循环，因为是进行了排序的，所以可以通过前后移动来保证是否能得到要求的结果;
func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	sort.Ints(nums)
	size := len(nums)

	if size > 0 && nums[0] > 0 {
		return result
	}

	for i := 0; i < size-2; i += 1 {
		// 去重
		if i > 0 && (nums[i] == nums[i-1]) {
			continue
		}

		first := i + 1
		second := size - 1
		target := 0 - nums[i]

		for first < second {
			//去重，前后一起去重

			if (nums[first] + nums[second]) > target {
				second -= 1
			} else if nums[first]+nums[second] < target {
				first += 1
			} else {
				tmp := []int{
					nums[i],
					nums[first],
					nums[second],
				}
				result = append(result, tmp)

				for first < second && (nums[first] == nums[first+1]) {
					first += 1
				}
				for first < second && (nums[second] == nums[second-1]) {
					second -= 1
				}

				first += 1
				second -= 1
			}
		}
	}
	return result
}
