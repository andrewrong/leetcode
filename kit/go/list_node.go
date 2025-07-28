package kit

import "math"

// ListNode is the definition for a singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList() *ListNode {
	return &ListNode{
		Val:  math.MaxInt,
		Next: nil,
	}
}
