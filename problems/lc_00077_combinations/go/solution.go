package lc_00077_combinations

// 组合[1,n]
func combine(n int, k int) [][]int {
	if k == 0 {
		return [][]int{}
	}

	sum := make([][]int, 0)

	for i := n; i > 0; i -= 1 {
		if k == 1 {
			sum = append(sum, []int{i})
		} else {
			result := combine(i-1, k-1)
			for _, v := range result {
				sum = append(sum, append(v, i))
			}
		}
	}

	return sum
}
