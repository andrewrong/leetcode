package lc_00039_combination_sum

import (
	"sort"
	"strconv"
)

type SliceKey struct {
	data []int
	Key  string
}

func NewSliceKey(data []int) SliceKey {
	sort.Ints(data)

	key := ""
	for _, v := range data {
		key += strconv.FormatInt(int64(v), 10)
	}
	return SliceKey{
		data: data,
		Key:  key,
	}
}

func (s *SliceKey) GetKey() string {
	return s.Key
}

func (s *SliceKey) GetData() []int {
	return s.data
}

// combinationSum 返回所有唯一组合，使得组合中的数字之和等于目标值
// 每个数字可以无限次使用
func combinationBySort(candidates []int, target int, startIdx int) [][]int {
	result := make([][]int, 0)

	if len(candidates) == 0 {
		return result
	}

	for i := startIdx; i < len(candidates); i += 1 {
		leave := target - candidates[i]
		if leave < 0 {
			break
		} else if leave == 0 {
			result = append(result, []int{candidates[i]})
			break
		} else {
			tmp := combinationBySort(candidates, leave, i)
			for _, v := range tmp {
				result = append(result, append(v, candidates[i]))
			}
		}
	}

	return result
}

func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}

	sort.Ints(candidates)
	result := combinationBySort(candidates, target, 0)

	removeSame := make(map[string]*SliceKey)
	for _, v := range result {
		sliceKey := NewSliceKey(v)
		if _, ok := removeSame[sliceKey.Key]; !ok {
			removeSame[sliceKey.Key] = &sliceKey
		}
	}

	theLst := make([][]int, 0)
	for _, v := range removeSame {
		theLst = append(theLst, v.GetData())
	}

	return theLst
}
