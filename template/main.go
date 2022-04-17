package main

/*
メモ：
*/

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	N := 0
	fmt.Fscanf(reader, "%d\n", &N)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscanf(reader, "%d", &A[i])
	}
}
