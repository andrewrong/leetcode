package lc_00142_linked_list_cycle_2

import (
	kit "leetcode-go/kit/go"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Helper function to create a cycle in the linked list
// pos is the index of the node that the tail connects to (0-indexed)
// If pos is -1, there is no cycle
func createListWithCycle(vals []int, pos int) *kit.ListNode {
	if len(vals) == 0 {
		return nil
	}

	head := &kit.ListNode{Val: vals[0]}
	curr := head
	nodes := []*kit.ListNode{head}

	for i := 1; i < len(vals); i++ {
		curr.Next = &kit.ListNode{Val: vals[i]}
		curr = curr.Next
		nodes = append(nodes, curr)
	}

	if pos >= 0 && pos < len(nodes) {
		curr.Next = nodes[pos] // Create cycle
	}

	return head
}

func TestDetectCycle(t *testing.T) {
	// Test case 1: List with cycle [3,2,0,-4], pos=1
	// The tail connects to node index 1 (node with value 2)
	// 3 2 0 -4
	head1 := createListWithCycle([]int{3, 2, 0, -4}, 1)
	result1 := Solve(head1)
	// Verify that result1 is the node with value 2 (the start of the cycle)
	assert.NotNil(t, result1)
	assert.Equal(t, 2, result1.Val)

	// Test case 2: List with cycle [1,2], pos=0
	// The tail connects to node index 0 (node with value 1)
	head2 := createListWithCycle([]int{1, 2}, 0)
	result2 := Solve(head2)
	// Verify that result2 is the node with value 1 (the start of the cycle)
	assert.NotNil(t, result2)
	assert.Equal(t, 1, result2.Val)

	// Test case 3: List without cycle [1]
	head3 := createListWithCycle([]int{1}, -1)
	result3 := Solve(head3)
	assert.Nil(t, result3)

	// Test case 4: Empty list
	var head4 *kit.ListNode
	result4 := Solve(head4)
	assert.Nil(t, result4)

	// Test case 5: List with cycle [1,2,3], pos=2
	// The tail connects to node index 2 (node with value 3)
	head5 := createListWithCycle([]int{1, 2, 3}, 2)
	result5 := Solve(head5)
	// Verify that result5 is the node with value 3 (the start of the cycle)
	assert.NotNil(t, result5)
	assert.Equal(t, 3, result5.Val)

	// Test case 6: Single node with cycle
	head6 := createListWithCycle([]int{1}, 0)
	result6 := Solve(head6)
	// Verify that result6 is the node with value 1
	assert.NotNil(t, result6)
	assert.Equal(t, 1, result6.Val)
}
