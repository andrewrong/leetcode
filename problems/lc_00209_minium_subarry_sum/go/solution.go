package lcc_00209_minium_subarry_sum

func Solve(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	min_size := len(nums) + 1
	current_size := 0
	sum := 0
	firstIndex := 0

	for _, data := range nums {
		sum += data
		current_size++

		for sum >= target {
			if min_size > current_size {
				min_size = current_size
			}

			sum -= nums[firstIndex]
			firstIndex++
			current_size--
		}
	}

	if min_size == len(nums)+1 {
		return 0
	}
	return min_size
}
