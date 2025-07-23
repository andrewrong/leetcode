package lcc_00027_remove_element

import (
	"reflect"
	"sort"
	"testing"
)

func TestSolve(t *testing.T) {
	testCases := []struct {
		name     string // 测试用例的描述
		nums     []int  // 输入数组
		val      int    // 要移除的值
		wantK    int    // 期望返回的长度 k
		wantNums []int  // 期望数组的前 k 个元素 (顺序无关)
	}{
		{
			name:     "Example 1",
			nums:     []int{3, 2, 2, 3},
			val:      3,
			wantK:    2,
			wantNums: []int{2, 2},
		},
		{
			name:     "Example 2",
			nums:     []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:      2,
			wantK:    5,
			wantNums: []int{0, 1, 3, 0, 4},
		},
		{
			name:     "No element to remove",
			nums:     []int{1, 2, 3, 4, 5},
			val:      6,
			wantK:    5,
			wantNums: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "All elements to remove",
			nums:     []int{4, 4, 4, 4},
			val:      4,
			wantK:    0,
			wantNums: []int{},
		},
		{
			name:     "Empty input slice",
			nums:     []int{},
			val:      1,
			wantK:    0,
			wantNums: []int{},
		},
		{
			name:     "Remove from the beginning",
			nums:     []int{5, 5, 1, 2, 3},
			val:      5,
			wantK:    3,
			wantNums: []int{1, 2, 3},
		},
		{
			name:     "Remove from the end",
			nums:     []int{1, 2, 3, 5, 5},
			val:      5,
			wantK:    3,
			wantNums: []int{1, 2, 3},
		},
	}

	for _, tc := range testCases {
		// 使用 t.Run 可以为每个测试用例创建子测试，方便管理和查看
		t.Run(tc.name, func(t *testing.T) {
			// 由于 removeElement 是原地修改数组，为了不影响其他测试用例，
			// 我们需要为每个用例创建一个输入的副本。
			numsCopy := make([]int, len(tc.nums))
			copy(numsCopy, tc.nums)

			// 调用待测试的函数
			gotK := Solve(numsCopy, tc.val)

			// 1. 验证返回的长度 k 是否正确
			if gotK != tc.wantK {
				t.Errorf("removeElement(%v, %v) returned k = %v; want %v", tc.nums, tc.val, gotK, tc.wantK)
				return // 如果长度不正确，后续的数组内容比较也无意义了
			}

			// 2. 验证数组内容是否正确 (由于顺序无关，先排序再比较)
			// 获取函数处理后实际得到的前 k 个元素
			gotNums := numsCopy[:gotK]

			// 对实际结果和期望结果都进行排序
			sort.Ints(gotNums)
			sort.Ints(tc.wantNums)

			// 使用 reflect.DeepEqual 比较两个排序后的切片
			if !reflect.DeepEqual(gotNums, tc.wantNums) {
				// 为了日志清晰，我们显示原始输入
				t.Errorf("removeElement(%v, %v) modified nums to %v; want first %v elements to be a permutation of %v", tc.nums, tc.val, numsCopy, tc.wantK, tc.wantNums)
			}
		})
	}
}
