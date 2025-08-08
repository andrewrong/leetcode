package solution

// FourSumCount returns the number of tuples (i, j, k, l) such that nums1[i] + nums2[j] + nums3[k] + nums4[l] == 0
func FourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	// Implementation will be added later
	ab := make(map[int]int)
	cd := make(map[int]int)

	for _, a := range nums1 {
		for _, b := range nums2 {
			v, ok := ab[a+b]
			if !ok {
				ab[a+b] = 1
			} else {
				ab[a+b] = v + 1
			}
		}
	}

	for _, c := range nums3 {
		for _, d := range nums4 {
			v, ok := cd[c+d]
			if !ok {
				cd[c+d] = 1
			} else {
				cd[c+d] = v + 1
			}
		}
	}
	resultSum := 0
	for k, count := range ab {
		target := 0 - k
		v, ok := cd[target]
		if !ok {
			continue
		}
		resultSum += count * v
	}

	return resultSum
}
