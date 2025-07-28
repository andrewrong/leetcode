package lc_00203_remove_element_in_list

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

func Solve(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	// create virtual head
	virtualHeadNode := &ListNode{
		Val:  math.MaxInt,
		Next: head,
	}

	defer func() {
		virtualHeadNode = nil
	}()

	tmp := virtualHeadNode
	for tmp.Next != nil {
		if tmp.Next.Val != val {
			tmp = tmp.Next
			continue
		}
		// delete tmp.Next
		tmp.Next = tmp.Next.Next
	}

	return virtualHeadNode.Next
}
