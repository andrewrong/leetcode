package lc_00019_remove_nth_node_from_end_of_list

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

func TestRemoveNthFromEnd(t *testing.T) {
	// Test case 1: Normal case [1,2,3,4,5], n=2 -> [1,2,3,5]
	head1 := createList([]int{1, 2, 3, 4, 5})
	expected1 := []int{1, 2, 3, 5}
	result1 := Solve(head1, 2)
	assert.Equal(t, expected1, listToSlice(result1))

	// Test case 2: Remove head [1], n=1 -> []
	head2 := createList([]int{1})
	var expected2 []int
	result2 := Solve(head2, 1)
	assert.Equal(t, expected2, listToSlice(result2))

	// Test case 3: Remove last node [1,2], n=1 -> [1]
	head3 := createList([]int{1, 2})
	expected3 := []int{1}
	result3 := Solve(head3, 1)
	assert.Equal(t, expected3, listToSlice(result3))

	// Test case 4: Remove first node [1,2], n=2 -> [2]
	head4 := createList([]int{1, 2})
	expected4 := []int{2}
	result4 := Solve(head4, 2)
	assert.Equal(t, expected4, listToSlice(result4))

	// Test case 5: Multiple nodes [1,2,3], n=2 -> [1,3]
	head5 := createList([]int{1, 2, 3})
	expected5 := []int{1, 3}
	result5 := Solve(head5, 2)
	assert.Equal(t, expected5, listToSlice(result5))
}
