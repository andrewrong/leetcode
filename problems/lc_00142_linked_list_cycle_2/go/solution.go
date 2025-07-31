package lc_00142_linked_list_cycle_2

import kit "leetcode-go/kit/go"

// detectCycle detects the node where the cycle in a linked list begins.
// TODO: Implement solution here
func Solve(head *kit.ListNode) *kit.ListNode {
	if head == nil {
		return nil
	}
	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if slow == fast {
			// encounter each other
			index1 := head
			index2 := slow

			for index1 != index2 {
				index1 = index1.Next
				index2 = index2.Next
			}

			return index1
		}
	}
	return nil
}
