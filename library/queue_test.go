package library

import "testing"

func TestQueue(t *testing.T) {
	q := NewQueue[int](4)
	if !q.Empty() {
		t.Fail()
		return
	}
	q.Push(0)
	if q.Pop() != 0 {
		t.Fail()
		return
	}
	q.Push(1)
	q.Push(2)
	if q.Pop() != 1 {
		t.Fail()
		return
	}
	if q.Front() != 2 {
		t.Fail()
		return
	}
	q.Push(3)
	if q.Empty() {
		t.Fail()
		return
	}
	if q.Pop() != 2 {
		t.Fail()
		return
	}
	if q.Pop() != 3 {
		t.Fail()
		return
	}
	if !q.Empty() {
		t.Fail()
		return
	}
}
