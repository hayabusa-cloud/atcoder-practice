package main

/*
メモ：
Greedy
*/

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	N, K, X := 0, 0, 0
	fmt.Fscanf(reader, "%d%d%d\n", &N, &K, &X)
	A := make([]int, N)
	sum := 0
	for i := 0; i < N; i++ {
		fmt.Fscanf(reader, "%d", &A[i])
		sum += A[i]
	}

	save := 0
	for i := 0; i < N && K > 0; i++ {
		if A[i] >= X {
			n := A[i] / X
			if n > K {
				n = K
			}
			K -= n
			A[i] = A[i] - n*X
			save += n * X
		}
	}
	if K == 0 {
		fmt.Println(sum - save)
		return
	}

	sort.Slice(A, func(i, j int) bool {
		return A[i] > A[j]
	})

	for i := 0; i < N && K > 0; i++ {
		save += A[i]
		A[i] = 0
		K--
	}
	fmt.Println(sum - save)
}
