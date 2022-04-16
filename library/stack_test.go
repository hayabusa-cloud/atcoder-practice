package library

import "testing"

func TestStack(t *testing.T) {
	s := NewStack[int](3)
	if !s.Empty() {
		t.Fail()
		return
	}
	s.Push(0)
	if s.Pop() != 0 {
		t.Fail()
		return
	}
	s.Push(1)
	s.Push(2)
	if s.Pop() != 2 {
		t.Fail()
		return
	}
	if s.Top() != 1 {
		t.Fail()
		return
	}
	s.Push(3)
	if s.Empty() {
		t.Fail()
		return
	}
	if s.Pop() != 3 {
		t.Fail()
		return
	}
	if s.Pop() != 1 {
		t.Fail()
		return
	}
	if !s.Empty() {
		t.Fail()
		return
	}
}
