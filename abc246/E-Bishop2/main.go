package main

/*
メモ：
最短経路BFS
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
	y int
}

func search(N int, board [][]bool, q *Queue, visited [][]int, cost [][]int, Bx int, By int) int {
	for {
		if q.Empty() {
			return -1
		}
		src := q.Pop()
		visited[src.x][src.y]++
		if src.x == Bx && src.y == By {
			return cost[src.x][src.y]
		}
		checkFn := func(x int, y int) bool {
			if visited[x][y] > 0 {
				return false
			}
			if board[x][y] {
				return false
			}
			if cost[src.x][src.y] >= cost[x][y] {
				return false
			} else if cost[src.x][src.y]+1 < cost[x][y] {
				cost[x][y] = cost[src.x][src.y] + 1
				q.Push(&elem{
					x: x,
					y: y,
				})
			}
			return true
		}
		// 左上
		for x, y := src.x-1, src.y-1; x >= 0 && y >= 0; x, y = x-1, y-1 {
			if !checkFn(x, y) {
				break
			}
		}
		// 左下
		for x, y := src.x-1, src.y+1; x >= 0 && y < N; x, y = x-1, y+1 {
			if !checkFn(x, y) {
				break
			}
		}
		// 右上
		for x, y := src.x+1, src.y-1; x < N && y >= 0; x, y = x+1, y-1 {
			if !checkFn(x, y) {
				break
			}
		}
		// 右下
		for x, y := src.x+1, src.y+1; x < N && y < N; x, y = x+1, y+1 {
			if !checkFn(x, y) {
				break
			}
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	N, Ax, Ay, Bx, By := 0, 0, 0, 0, 0
	fmt.Fscanf(reader, "%d\n%d%d\n%d%d\n", &N, &Ax, &Ay, &Bx, &By)
	Ax--
	Ay--
	Bx--
	By--

	board := make([][]bool, N)
	for i := 0; i < N; i++ {
		board[i] = make([]bool, N)
		var s string
		fmt.Fscanf(reader, "%s\n", &s)
		for j := 0; j < N; j++ {
			if s[j] == '#' {
				board[i][j] = true
			}
		}
	}

	if (Ax+Ay)&1 != (Bx+By)&1 {
		fmt.Println(-1)
		return
	}

	visited := make([][]int, N)
	for i := 0; i < N; i++ {
		visited[i] = make([]int, N)
	}

	cost := make([][]int, N)
	for i := 0; i < N; i++ {
		cost[i] = make([]int, N)
		for j := 0; j < N; j++ {
			cost[i][j] = N*N + 1
		}
	}

	q := NewQueue(1500*1500 + 1)
	q.Push(&elem{
		x: Ax,
		y: Ay,
	})

	cost[Ax][Ay] = 0
	res := search(N, board, q, visited, cost, Bx, By)

	fmt.Println(res)
}
