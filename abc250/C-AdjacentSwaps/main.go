package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	N, Q := 0, 0
	fmt.Fscanf(reader, "%d%d\n", &N, &Q)
	a := make([]int, N+1)
	pos := make([]int, N+1)
	for i := 1; i <= N; i++ {
		a[i] = i
		pos[i] = i
	}
	for i := 0; i < Q; i++ {
		x := 0
		fmt.Fscanf(reader, "%d\n", &x)
		if pos[x] == N {
			a[pos[x]], a[pos[x]-1], pos[x], pos[a[pos[x]-1]] = a[pos[x]-1], a[pos[x]], pos[a[pos[x]-1]], pos[x]
			continue
		}
		a[pos[x]], a[pos[x]+1], pos[x], pos[a[pos[x]+1]] = a[pos[x]+1], a[pos[x]], pos[a[pos[x]+1]], pos[x]
	}
	for i := 1; i < N; i++ {
		fmt.Print(a[i], " ")
	}
	fmt.Println(a[N])
}
