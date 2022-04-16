package main

/*
メモ：
再帰
*/

import (
	"bufio"
	"fmt"
	"os"
)

func f(n int) {
	if n == 1 {
		fmt.Print(n)
		return
	}
	f(n - 1)
	fmt.Print(" ", n, " ")
	f(n - 1)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	N := 0
	fmt.Fscanf(reader, "%d\n", &N)
	f(N)
}
