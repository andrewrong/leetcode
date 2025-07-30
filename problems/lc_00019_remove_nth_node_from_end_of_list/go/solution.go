package lc_00019_remove_nth_node_from_end_of_list

import kit "leetcode-go/kit/go"

// removeNthFromEnd removes the nth node from the end of the list.
// TODO: Implement solution here
func Solve(head *kit.ListNode, n int) *kit.ListNode {
	if head == nil || n <= 0 {
		return head
	}

	vNode := &kit.ListNode{
		Next: head,
	}

	fast := vNode
	slow := vNode

	// skip n = there are n nodes in the middle of fast/slow;
	for i := 0; i < n; i++ {
		fast = fast.Next
		if fast == nil {
			return head
		}
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next
	return vNode.Next
}
