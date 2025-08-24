// leetcode/problems/lc_00225_implement_stack_using_queues/go/solution.go
package main

import "container/list"

type MyStack struct {
	queue *list.List
}

func Constructor() MyStack {
	stack := MyStack{
		queue: list.New(),
	}
	return stack
}

func (s *MyStack) Push(x int) {
	s.queue.PushBack(x)
}

func (s *MyStack) Pop() int {
	size := s.queue.Len()
	for i := 0; i < size-1; i += 1 {
		ele := s.queue.Front()
		s.queue.Remove(ele)
		s.queue.PushBack(ele.Value.(int))
	}
	ele := s.queue.Front()
	s.queue.Remove(ele)
	return ele.Value.(int)
}

func (s *MyStack) Top() int {
	value := s.queue.Back()
	return value.Value.(int)
}

func (s *MyStack) Empty() bool {
	if s.queue.Len() == 0 {
		return true
	}
	return false
}
