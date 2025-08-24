package kit

import (
	"testing"
)

// TestIntegerStack tests the generic stack implementation with the `int` type.
func TestIntegerStack(t *testing.T) {
	t.Run("NewStack", func(t *testing.T) {
		s := NewStack[int]()
		if s == nil {
			t.Fatal("NewStack[int]() returned nil")
		}

		as, ok := s.(*ArrayStack[int])
		if !ok {
			t.Fatal("NewStack[int]() did not return a concrete *ArrayStack[int]")
		}

		if !as.IsEmpty() {
			t.Error("Expected new stack to be empty")
		}
		if as.Size() != 0 {
			t.Errorf("Expected new stack size to be 0, got %d", as.Size())
		}
	})

	t.Run("Push and Size", func(t *testing.T) {
		s := NewStack[int]()
		s.Push(10)
		as := s.(*ArrayStack[int])

		if as.IsEmpty() {
			t.Error("Stack should not be empty after push")
		}
		if as.Size() != 1 {
			t.Errorf("Expected size to be 1, got %d", as.Size())
		}

		s.Push(20)
		if as.Size() != 2 {
			t.Errorf("Expected size to be 2, got %d", as.Size())
		}
	})

	t.Run("Peek", func(t *testing.T) {
		s := NewStack[int]()
		s.Push(10)
		s.Push(20)

		val, ok := s.Peek()
		if !ok {
			t.Error("Expected Peek to succeed, but it failed")
		}
		if val != 20 {
			t.Errorf("Expected peeked value to be 20, got %d", val)
		}

		as := s.(*ArrayStack[int])
		if as.Size() != 2 {
			t.Error("Peek should not remove the element from the stack")
		}
	})

	t.Run("Pop", func(t *testing.T) {
		s := NewStack[int]()
		s.Push(10)
		s.Push(20)

		val, ok := s.Pop()
		if !ok {
			t.Error("Expected Pop to succeed, but it failed")
		}
		if val != 20 {
			t.Errorf("Expected popped value to be 20, got %d", val)
		}

		as := s.(*ArrayStack[int])
		if as.Size() != 1 {
			t.Errorf("Expected size to be 1 after pop, got %d", as.Size())
		}

		val, ok = s.Pop()
		if !ok {
			t.Error("Expected Pop to succeed, but it failed")
		}
		if val != 10 {
			t.Errorf("Expected popped value to be 10, got %d", val)
		}

		if !as.IsEmpty() {
			t.Error("Stack should be empty after popping all elements")
		}
	})

	t.Run("Operations on Empty Stack", func(t *testing.T) {
		s := NewStack[int]()

		val, ok := s.Pop()
		if ok {
			t.Error("Expected Pop on empty stack to fail, but it succeeded")
		}
		if val != 0 {
			t.Errorf("Expected popped value from empty stack to be the zero value (0), got %d", val)
		}

		val, ok = s.Peek()
		if ok {
			t.Error("Expected Peek on empty stack to fail, but it succeeded")
		}
		if val != 0 {
			t.Errorf("Expected peeked value from empty stack to be the zero value (0), got %d", val)
		}
	})
}

// TestStringStack tests the generic stack implementation with the `string` type.
func TestStringStack(t *testing.T) {
	t.Run("Push and Pop", func(t *testing.T) {
		s := NewStack[string]()
		s.Push("hello")
		s.Push("world")

		val, ok := s.Pop()
		if !ok || val != "world" {
			t.Errorf(`Expected to pop "world", got %q (ok: %v)`, val, ok)
		}

		val, ok = s.Pop()
		if !ok || val != "hello" {
			t.Errorf(`Expected to pop "hello", got %q (ok: %v)`, val, ok)
		}
	})

	t.Run("Peek", func(t *testing.T) {
		s := NewStack[string]()
		s.Push("alpha")

		val, ok := s.Peek()
		if !ok || val != "alpha" {
			t.Errorf(`Expected to peek "alpha", got %q (ok: %v)`, val, ok)
		}
	})

	t.Run("Operations on Empty Stack", func(t *testing.T) {
		s := NewStack[string]()

		val, ok := s.Pop()
		if ok {
			t.Error("Expected Pop on empty stack to fail")
		}
		if val != "" { // Zero value for string is ""
			t.Errorf(`Expected zero value "", got %q`, val)
		}
	})
}