package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	R, C, A := 0, 0, [4]int{}
	fmt.Fscanf(reader, "%d%d\n", &R, &C)
	fmt.Fscanf(reader, "%d%d\n", &A[0], &A[1])
	fmt.Fscanf(reader, "%d%d\n", &A[2], &A[3])
	fmt.Println(A[R*2-2+C-1])
}
