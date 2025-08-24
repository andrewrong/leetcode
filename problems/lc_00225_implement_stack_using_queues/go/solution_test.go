// leetcode/problems/lc_00225_implement_stack_using_queues/go/solution_test.go
package main

import "testing"

func TestMyStack(t *testing.T) {
	s := Constructor()

	// Initially, the stack should be empty
	if !s.Empty() {
		t.Error("Expected new stack to be empty")
	}

	// Push 1
	s.Push(1)
	if s.Empty() {
		t.Error("Stack should not be empty after push")
	}
	if s.Top() != 1 {
		t.Errorf("Expected top to be 1, got %d", s.Top())
	}

	// Push 2
	s.Push(2)
	if s.Top() != 2 {
		t.Errorf("Expected top to be 2, got %d", s.Top())
	}

	// Pop
	val := s.Pop()
	if val != 2 {
		t.Errorf("Expected popped value to be 2, got %d", val)
	}
	if s.Top() != 1 {
		t.Errorf("Expected top to be 1, got %d", s.Top())
	}

	// Push 3
	s.Push(3)
	if s.Top() != 3 {
		t.Errorf("Expected top to be 3, got %d", s.Top())
	}

	// Pop remaining elements
	if s.Pop() != 3 {
		t.Error("Expected popped value to be 3")
	}
	if s.Pop() != 1 {
		t.Error("Expected popped value to be 1")
	}
	if !s.Empty() {
		t.Error("Expected stack to be empty")
	}
}
