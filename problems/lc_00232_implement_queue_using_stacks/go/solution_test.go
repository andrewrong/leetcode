// leetcode/problems/lc_00232_implement_queue_using_stacks/go/solution_test.go
package main

import "testing"

func TestMyQueue(t *testing.T) {
	q := Constructor()

	// Initially, the queue should be empty
	if !q.Empty() {
		t.Error("Expected new queue to be empty")
	}

	// Push 1
	q.Push(1)
	if q.Empty() {
		t.Error("Queue should not be empty after push")
	}
	if q.Peek() != 1 {
		t.Errorf("Expected peek to be 1, got %d", q.Peek())
	}

	// Push 2
	q.Push(2)
	if q.Peek() != 1 {
		t.Errorf("Expected peek to be 1, got %d", q.Peek())
	}

	// Pop
	val := q.Pop()
	if val != 1 {
		t.Errorf("Expected popped value to be 1, got %d", val)
	}
	if q.Peek() != 2 {
		t.Errorf("Expected peek to be 2, got %d", q.Peek())
	}

	// Push 3
	q.Push(3)
	if q.Peek() != 2 {
		t.Errorf("Expected peek to be 2, got %d", q.Peek())
	}

	// Pop remaining elements
	if q.Pop() != 2 {
		t.Error("Expected popped value to be 2")
	}
	if q.Pop() != 3 {
		t.Error("Expected popped value to be 3")
	}
	if !q.Empty() {
		t.Error("Expected queue to be empty")
	}
}
