package lc_00206_reverse_linked_list

import (
	kit "leetcode-go/kit/go"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Helper function to create a linked list from a slice
func createList(vals []int) *kit.ListNode {
	if len(vals) == 0 {
		return nil
	}

	head := &kit.ListNode{Val: vals[0]}
	curr := head
	for i := 1; i < len(vals); i++ {
		curr.Next = &kit.ListNode{Val: vals[i]}
		curr = curr.Next
	}
	return head
}

// Helper function to convert a linked list to a slice
func listToSlice(head *kit.ListNode) []int {
	var result []int
	curr := head
	for curr != nil {
		result = append(result, curr.Val)
		curr = curr.Next
	}
	return result
}

func TestReverseList(t *testing.T) {
	// Test case 1: Normal list [1,2,3,4,5] -> [5,4,3,2,1]
	head1 := createList([]int{1, 2, 3, 4, 5})
	expected1 := []int{5, 4, 3, 2, 1}
	result1 := SolveListRecursion(head1)
	assert.Equal(t, expected1, listToSlice(result1))

	// Test case 2: Single node [1] -> [1]
	head2 := createList([]int{1})
	expected2 := []int{1}
	result2 := SolveListRecursion(head2)
	assert.Equal(t, expected2, listToSlice(result2))

	// Test case 3: Empty list [] -> []
	head3 := createList([]int{})
	var expected3 []int
	result3 := SolveListRecursion(head3)
	assert.Equal(t, expected3, listToSlice(result3))

	// Test case 4: Two nodes [1,2] -> [2,1]
	head4 := createList([]int{1, 2})
	expected4 := []int{2, 1}
	result4 := SolveListRecursion(head4)
	assert.Equal(t, expected4, listToSlice(result4))
}
