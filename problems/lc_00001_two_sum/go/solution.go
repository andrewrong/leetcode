package solution

// TwoSum finds two numbers in the array that add up to the target
func TwoSum(nums []int, target int) []int {
	// Implementation will be added later
	cache := make(map[int]int)
	for index, element := range nums {
		cache[element] = index
	}

	for index, element := range nums {
		other := target - element

		preIndex, ok := cache[other]
		if ok && index != preIndex {
			return []int{
				min(preIndex, index),
				max(preIndex, index),
			}
		}
	}

	return []int{}

}
