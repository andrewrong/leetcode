package kit

import (
	"testing"
)

// TestIntegerQueue tests the generic queue implementation with the `int` type.
func TestIntegerQueue(t *testing.T) {
	t.Run("NewQueue", func(t *testing.T) {
		q := NewQueue[int]()
		if q == nil {
			t.Fatal("NewQueue[int]() returned nil")
		}
		if !q.IsEmpty() {
			t.Error("Expected new queue to be empty")
		}
		if q.Size() != 0 {
			t.Errorf("Expected new queue size to be 0, got %d", q.Size())
		}
	})

	t.Run("Enqueue and Size", func(t *testing.T) {
		q := NewQueue[int]()
		q.Enqueue(10)
		if q.IsEmpty() {
			t.Error("Queue should not be empty after enqueue")
		}
		if q.Size() != 1 {
			t.Errorf("Expected size to be 1, got %d", q.Size())
		}
		q.Enqueue(20)
		if q.Size() != 2 {
			t.Errorf("Expected size to be 2, got %d", q.Size())
		}
	})

	t.Run("Dequeue", func(t *testing.T) {
		q := NewQueue[int]()
		q.Enqueue(10)
		q.Enqueue(20)

		val, ok := q.Dequeue()
		if !ok {
			t.Error("Expected Dequeue to succeed, but it failed")
		}
		if val != 10 {
			t.Errorf("Expected dequeued value to be 10 (FIFO), got %d", val)
		}
		if q.Size() != 1 {
			t.Errorf("Expected size to be 1 after dequeue, got %d", q.Size())
		}

		val, ok = q.Dequeue()
		if !ok {
			t.Error("Expected Dequeue to succeed, but it failed")
		}
		if val != 20 {
			t.Errorf("Expected dequeued value to be 20, got %d", val)
		}
		if !q.IsEmpty() {
			t.Error("Queue should be empty after dequeuing all elements")
		}
	})

	t.Run("FIFO Order", func(t *testing.T) {
		q := NewQueue[int]()
		elements := []int{1, 2, 3, 4, 5}
		for _, elem := range elements {
			q.Enqueue(elem)
		}

		for i := 0; i < len(elements); i++ {
			val, ok := q.Dequeue()
			if !ok {
				t.Fatalf("Dequeue failed unexpectedly for element %d", elements[i])
			}
			if val != elements[i] {
				t.Errorf("Expected dequeued value %d, but got %d", elements[i], val)
			}
		}
	})

	t.Run("Operations on Empty Queue", func(t *testing.T) {
		q := NewQueue[int]()
		val, ok := q.Dequeue()
		if ok {
			t.Error("Expected Dequeue on empty queue to fail, but it succeeded")
		}
		if val != 0 {
			t.Errorf("Expected dequeued value from empty queue to be the zero value (0), got %d", val)
		}
	})
}

// TestStringQueue tests the generic queue implementation with the `string` type.
func TestStringQueue(t *testing.T) {
	t.Run("Enqueue and Dequeue", func(t *testing.T) {
		q := NewQueue[string]()
		q.Enqueue("hello")
		q.Enqueue("world")

		val, ok := q.Dequeue()
		if !ok || val != "hello" {
			t.Errorf(`Expected to dequeue "hello", got %q (ok: %v)`, val, ok)
		}

		val, ok = q.Dequeue()
		if !ok || val != "world" {
			t.Errorf(`Expected to dequeue "world", got %q (ok: %v)`, val, ok)
		}
	})

	t.Run("Operations on Empty Queue", func(t *testing.T) {
		q := NewQueue[string]()
		val, ok := q.Dequeue()
		if ok {
			t.Error("Expected Dequeue on empty queue to fail")
		}
		if val != "" { // Zero value for string is ""
			t.Errorf(`Expected zero value "", got %q`, val)
		}
	})
}
