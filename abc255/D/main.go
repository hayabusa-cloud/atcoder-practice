package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	N, Q := 0, 0
	fmt.Fscanf(reader, "%d%d\n", &N, &Q)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscanf(reader, "%d", &A[i])
	}
	fmt.Fscanln(reader)

	sort.Ints(A)
	l, h, min, max := 0, 0, A[0], A[N-1]
	for i := 0; i < N; i++ {
		l += A[i] - min
		h += max - A[i]
	}

	for q := 0; q < Q; q++ {
		X := 0
		fmt.Fscanf(reader, "%d\n", &X)

		if X <= min {
			fmt.Println(l + (min-X)*N)
			continue
		}
		if X >= max {
			fmt.Println(h + (X-max)*N)
			continue
		}

		i := sort.SearchInts(A, X)

		res1 := l - (X-min)*(N-i) + (X-min)*i
		res2 := h + (max-X)*(N-1) - (max-X)*i

		if res1 < res2 {
			fmt.Println(res1)
		} else {
			fmt.Println(res2)
		}
	}
}
