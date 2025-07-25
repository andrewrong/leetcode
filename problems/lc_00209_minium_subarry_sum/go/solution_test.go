package lcc_00209_minium_subarry_sum

import (
	"testing"
)

func TestSolve(t *testing.T) {
	testCases := []struct {
		name     string // 测试用例的描述
		target   int    // 目标值
		nums     []int  // 输入数组
		wantSize int    // 期望返回的最小子数组长度
	}{
		{
			name:     "Example 1",
			target:   7,
			nums:     []int{2, 3, 1, 2, 4, 3},
			wantSize: 2,
		},
		{
			name:     "Example 2",
			target:   4,
			nums:     []int{1, 4, 4},
			wantSize: 1,
		},
		{
			name:     "Example 3",
			target:   11,
			nums:     []int{1, 1, 1, 1, 1, 1, 1, 1},
			wantSize: 0,
		},
		{
			name:     "Target equals sum of all elements",
			target:   10,
			nums:     []int{1, 2, 3, 4},
			wantSize: 4,
		},
		{
			name:     "Target greater than sum of all elements",
			target:   15,
			nums:     []int{1, 2, 3, 4},
			wantSize: 0,
		},
		{
			name:     "Single element equals target",
			target:   5,
			nums:     []int{5},
			wantSize: 1,
		},
		{
			name:     "Single element less than target",
			target:   5,
			nums:     []int{3},
			wantSize: 0,
		},
		{
			name:     "Empty input slice",
			target:   1,
			nums:     []int{},
			wantSize: 0,
		},
		{
			name:     "Large target with small elements",
			target:   100,
			nums:     []int{1, 2, 3, 4, 5},
			wantSize: 0,
		},
		{
			name:     "Exact match at the beginning",
			target:   6,
			nums:     []int{3, 3, 1, 2, 4},
			wantSize: 2,
		},
	}

	for _, tc := range testCases {
		// 使用 t.Run 可以为每个测试用例创建子测试，方便管理和查看
		t.Run(tc.name, func(t *testing.T) {
			// 调用待测试的函数
			gotSize := Solve(tc.target, tc.nums)

			// 验证返回的长度是否正确
			if gotSize != tc.wantSize {
				t.Errorf("Solve(%d, %v) = %d; want %d", tc.target, tc.nums, gotSize, tc.wantSize)
			}
		})
	}
}