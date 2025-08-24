package kit

type Stack[T any] interface {
	Pop() (T, bool)
	Push(value T)
	Peek() (T, bool)
	Size() int
}

type ArrayStack[T any] struct {
	data []T
}

func NewStack[T any]() Stack[T] {
	return &ArrayStack[T]{
		data: make([]T, 0),
	}
}

func (s *ArrayStack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}

	size := len(s.data)

	value := s.data[size-1]
	s.data = s.data[:size-1]
	return value, true
}

func (s *ArrayStack[T]) Push(value T) {
	s.data = append(s.data, value)
}

func (s *ArrayStack[T]) Peek() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}

	value := s.data[len(s.data)-1]
	return value, true
}

func (s *ArrayStack[T]) IsEmpty() bool {
	if len(s.data) == 0 {
		return true
	}

	return false
}

func (s *ArrayStack[T]) Size() int {
	return len(s.data)
}
