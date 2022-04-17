package main

/*
メモ：二分探索
*/

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	N, Q := 0, 0
	fmt.Fscanf(reader, "%d\n", &N)
	A := make([]int, N+1)
	for i := 1; i <= N; i++ {
		fmt.Fscanf(reader, "%d", &A[i])
	}
	fmt.Fscanf(reader, "\n")

	idx := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		idx[i] = make([]int, 0)
	}

	for i := 1; i <= N; i++ {
		a := A[i]
		idx[a] = append(idx[a], i)
	}

	fmt.Fscanf(reader, "%d\n", &Q)
	for i := 0; i < Q; i++ {
		L, R, X := 0, 0, 0
		fmt.Fscanf(reader, "%d%d%d\n", &L, &R, &X)
		if len(idx[X]) < 1 {
			fmt.Println(0)
			continue
		}
		l := sort.Search(len(idx[X]), func(i int) bool {
			return idx[X][i] >= L
		})
		r := sort.Search(len(idx[X]), func(i int) bool {
			return idx[X][i] > R
		})
		fmt.Println(r - l)
	}
}
