package lcc_00027_remove_element

// solve is the function for the problem.
// Remember to rename it and make it public (e.g. Solve) if you want to call it from main.go

// 这个题目是原地移除和val一样的值，最后返回数组的前面的k个不包含val的数据
func Solve(nums []int, val int) int {
	size := 0
	prev := -1

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			size += 1

			// prev是表示下一个可以被插入的val
			if prev != -1 {
				nums[prev] = nums[i]
				prev += 1
			} else {
				nums[0] = nums[i]
				prev = 1
			}
		}
	}

	return size
}
