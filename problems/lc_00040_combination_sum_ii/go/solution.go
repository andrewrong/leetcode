package lc_00040_combination_sum_ii

import (
	"sort"
)

// combinationSum2 返回所有唯一组合，使得组合中的数字之和等于目标值
// 每个数字只能使用一次
func combinationSum2BySort(candidates []int, target int, startIdx int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}

	result := make([][]int, 0)

	for i := startIdx; i < len(candidates); i += 1 {
		if i > startIdx && candidates[i] == candidates[i-1] {
			continue
		}
		leave := target - candidates[i]
		if leave < 0 {
			break
		} else if leave == 0 {
			result = append(result, []int{candidates[i]})
			break
		} else {
			tmp := combinationSum2BySort(candidates, leave, i+1)
			for _, v := range tmp {
				result = append(result, append([]int{candidates[i]}, v...))
			}
		}
	}
	return result
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	oneAnswer := make([]int, 0)
	used := make([]bool, len(candidates))
	last := make([][]int, 0)

	combinationSumOther(candidates, target, 0, &oneAnswer, used, &last)

	return last
}

func combinationSumOther(candidates []int, target int, startIdx int, result *[]int, used []bool, last *[][]int) {
	if len(candidates) == 0 {
		return
	}

	if target == 0 {
		cp := make([]int, len(*result))
		copy(cp, *result)
		*last = append(*last, cp)
		return
	}

	for i := startIdx; i < len(candidates); i += 1 {
		leave := target - candidates[i]

		if i > startIdx && candidates[i] == candidates[i-1] && !used[i-1] {
			continue
		}

		if leave < 0 {
			return
		} else {
			*result = append(*result, candidates[i])
			used[i] = true
			combinationSumOther(candidates, leave, i+1, result, used, last)
			used[i] = false
			*result = (*result)[:len(*result)-1]
		}
	}
}
