package solution

// Intersection finds the intersection of two integer slices
func Intersection(nums1 []int, nums2 []int) []int {
	// Implementation will be added later
	result := make(map[int]struct{})
	record := make(map[int]struct{})
	for _, element := range nums1 {
		record[element] = struct{}{}
	}

	for _, element := range nums2 {
		if _, ok := record[element]; ok {
			result[element] = struct{}{}
		}
	}

	r := make([]int, 0)
	for k := range result {
		r = append(r, k)
	}

	return r
}
