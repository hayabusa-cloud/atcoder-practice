package library

type Queue[T any] struct {
	val         []T
	front, tail int
}

func NewQueue[T any](cap int) *Queue[T] {
	return &Queue[T]{
		val:   make([]T, cap),
		front: 0,
		tail:  0,
	}
}

func (q *Queue[T]) Empty() bool {
	return q.front == q.tail
}
func (q *Queue[T]) Push(x T) {
	q.val[q.tail] = x
	q.tail++
}
func (q *Queue[T]) Pop() T {
	ret := q.val[q.front]
	q.front++
	return ret
}
func (q *Queue[T]) Front() T {
	return q.val[q.front]
}
