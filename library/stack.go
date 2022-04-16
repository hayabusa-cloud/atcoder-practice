package library

type Stack[T any] struct {
	val []T
	top int
}

func NewStack[T any](cap int) *Stack[T] {
	return &Stack[T]{
		val: make([]T, cap),
		top: 0,
	}
}

func (s *Stack[T]) Empty() bool {
	return s.top == 0
}
func (s *Stack[T]) Push(x T) {
	s.val[s.top] = x
	s.top++
}
func (s *Stack[T]) Pop() T {
	s.top--
	return s.val[s.top]
}
func (s *Stack[T]) Top() T {
	return s.val[s.top-1]
}
