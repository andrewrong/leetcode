package lc_00206_reverse_linked_list

import kit "leetcode-go/kit/go"

// TODO: Implement solution here
func Solve(head *kit.ListNode) *kit.ListNode {
	if head == nil {
		return nil
	}

	first := head
	var second *kit.ListNode = nil

	for first != nil {
		tmp := first.Next
		first.Next = second
		second = first
		first = tmp
	}

	return second
}

func SolveListRecursion(head *kit.ListNode) *kit.ListNode {
	return recursion(nil, head)
}

func recursion(prev, cur *kit.ListNode) *kit.ListNode {
	if cur == nil {
		return prev
	}

	tmp := cur.Next
	cur.Next = prev
	return recursion(cur, tmp)
}
