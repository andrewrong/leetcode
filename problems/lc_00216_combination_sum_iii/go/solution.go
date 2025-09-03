package solution

func combinationSum3(k int, n int) [][]int {
	return combine(k, n, 1)
}

func combine(k int, n int, startIdx int) [][]int {
	result := make([][]int, 0)
	if k == 0 {
		return result
	}

	for i := startIdx; i <= 9; i += 1 {
		leave := n - i
		if leave < 0 {
			break
		}

		if k == 1 && leave == 0 {
			result = append(result, []int{i})
		} else {
			tmps := combine(k-1, leave, i+1)
			for _, v := range tmps {
				result = append(result, append(v, i))
			}
		}
	}

	return result
}
