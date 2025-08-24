// leetcode/problems/lc_00232_implement_queue_using_stacks/go/solution.go
package main

type MyQueue struct {
	stackIn  []int
	stackOut []int
}

func Constructor() MyQueue {
	return MyQueue{
		stackIn:  make([]int, 0),
		stackOut: make([]int, 0),
	}
}

func (q *MyQueue) Push(x int) {
	q.stackIn = append(q.stackIn, x)
}

func (q *MyQueue) Pop() int {
	if len(q.stackOut) == 0 {
		inSize := len(q.stackIn)

		for i := inSize - 1; i >= 0; i -= 1 {
			q.stackOut = append(q.stackOut, q.stackIn[i])
		}
		q.stackIn = q.stackIn[:0]
	}

	element := q.stackOut[len(q.stackOut)-1]
	q.stackOut = q.stackOut[:len(q.stackOut)-1]
	return element
}

func (q *MyQueue) Peek() int {
	if len(q.stackOut) == 0 {
		inSize := len(q.stackIn)

		for i := inSize - 1; i >= 0; i -= 1 {
			q.stackOut = append(q.stackOut, q.stackIn[i])
		}
		q.stackIn = q.stackIn[:0]
	}

	element := q.stackOut[len(q.stackOut)-1]
	return element
}

func (q *MyQueue) Empty() bool {
	if len(q.stackIn) == 0 && len(q.stackOut) == 0 {
		return true
	}
	return false
}
