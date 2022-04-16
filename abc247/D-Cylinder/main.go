package main

/*
メモ：
キュー
*/

import (
	"bufio"
	"fmt"
	"os"
)

type Queue struct {
	val         []*elem
	front, tail int
}

func NewQueue(cap int) *Queue {
	return &Queue{
		val:   make([]*elem, cap),
		front: 0,
		tail:  0,
	}
}

func (q *Queue) Empty() bool {
	return q.front == q.tail
}
func (q *Queue) Push(x *elem) {
	q.val[q.tail] = x
	q.tail++
}
func (q *Queue) Pop() *elem {
	ret := q.val[q.front]
	q.front++
	return ret
}
func (q *Queue) Front() *elem {
	return q.val[q.front]
}

type elem struct {
	x int
	c int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	q := NewQueue(200000)
	var Q, op, x, c int
	_, _ = fmt.Fscanf(reader, "%d\n", &Q)
	for i := 0; i < Q; i++ {
		_, _ = fmt.Fscanf(reader, "%d", &op)
		if op == 1 {
			_, _ = fmt.Fscanf(reader, "%d%d\n", &x, &c)
			q.Push(&elem{
				x: x,
				c: c,
			})
		}
		if op == 2 {
			_, _ = fmt.Fscanf(reader, "%d\n", &c)
			sum := int64(0)
			for c > 0 {
				f := q.Front()
				if c >= f.c {
					sum += int64(f.c) * int64(f.x)
					q.Pop()
					c -= f.c
				} else {
					sum += int64(c) * int64(f.x)
					f.c -= c
					c = 0
				}
			}
			_, _ = fmt.Println(sum)
		}
	}
}
