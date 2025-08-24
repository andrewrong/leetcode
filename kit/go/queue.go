package kit

import "container/list"

type Queue[T any] interface {
	Size() int
	Enqueue(v T)
	Dequeue() (T, bool)
	IsEmpty() bool
}

type listQueue[T any] struct {
	elements *list.List
}

func NewQueue[T any]() Queue[T] {
	return &listQueue[T]{
		elements: list.New(),
	}
}

func (q *listQueue[T]) Size() int {
	return q.elements.Len()
}

func (q *listQueue[T]) Enqueue(v T) {
	q.elements.PushBack(v)
}

func (q *listQueue[T]) Dequeue() (T, bool) {
	if q.IsEmpty() {
		var zero T
		return zero, false
	}

	value := q.elements.Front()
	q.elements.Remove(value)
	return value.Value.(T), true
}

func (q *listQueue[T]) IsEmpty() bool {
	if q.elements.Len() == 0 {
		return true
	}
	return false
}
