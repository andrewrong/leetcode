package lc_00203_remove_element_in_list

import (
	"reflect"
	"testing"
)

// Helper function to create a linked list from a slice
func createList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	head := &ListNode{Val: vals[0]}
	current := head
	for i := 1; i < len(vals); i++ {
		current.Next = &ListNode{Val: vals[i]}
		current = current.Next
	}
	return head
}

// Helper function to convert a linked list to a slice
func listToSlice(head *ListNode) []int {
	var result []int = make([]int, 0)
	current := head
	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}
	return result
}

func TestSolve(t *testing.T) {
	testCases := []struct {
		name  string // 测试用例的描述
		input []int  // 输入链表的值
		val   int    // 要移除的值
		want  []int  // 期望的结果链表
	}{
		{
			name:  "Example 1",
			input: []int{1, 2, 6, 3, 4, 5, 6},
			val:   6,
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "Empty list",
			input: []int{},
			val:   1,
			want:  []int{},
		},
		{
			name:  "Remove all elements",
			input: []int{7, 7, 7, 7},
			val:   7,
			want:  []int{},
		},
		{
			name:  "Remove from beginning",
			input: []int{1, 1, 2, 3, 4},
			val:   1,
			want:  []int{2, 3, 4},
		},
		{
			name:  "Remove from end",
			input: []int{1, 2, 3, 4, 4},
			val:   4,
			want:  []int{1, 2, 3},
		},
		{
			name:  "Remove from middle",
			input: []int{1, 2, 3, 2, 4},
			val:   2,
			want:  []int{1, 3, 4},
		},
		{
			name:  "Single node to remove",
			input: []int{5},
			val:   5,
			want:  []int{},
		},
		{
			name:  "Single node to keep",
			input: []int{5},
			val:   3,
			want:  []int{5},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create linked list from input slice
			head := createList(tc.input)

			// Call the function to test
			got := Solve(head, tc.val)

			// Convert result linked list to slice for comparison
			gotSlice := listToSlice(got)

			// Compare the result with expected
			if !reflect.DeepEqual(gotSlice, tc.want) {
				t.Errorf("Solve(%v, %v) = %v; want %v", tc.input, tc.val, gotSlice, tc.want)
			}
		})
	}
}
