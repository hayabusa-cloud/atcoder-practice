package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	N, K := 0, 0
	fmt.Fscanf(reader, "%d%d\n", &N, &K)
	A := make([]int, K)
	X := make([]int, N+1)
	Y := make([]int, N+1)
	for i := 0; i < K; i++ {
		fmt.Fscanf(reader, "%d", &A[i])
	}
	fmt.Fscanln(reader)
	for i := 1; i <= N; i++ {
		fmt.Fscanf(reader, "%d%d\n", &X[i], &Y[i])
	}

	max := float64(0)
	for i := 1; i <= N; i++ {
		x, y := float64(X[i]), float64(Y[i])
		min := math.MaxFloat32
		for k := 0; k < K; k++ {
			lx, ly := float64(X[A[k]]), float64(Y[A[k]])
			dis2 := (lx-x)*(lx-x) + (ly-y)*(ly-y)
			if dis2 < min {
				min = dis2
			}
		}
		if min > max {
			max = min
		}
	}
	fmt.Println(math.Sqrt(max))
}
