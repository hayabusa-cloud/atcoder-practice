package main

/*
メモ：シミュレート
*/

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var A, B, K int64 = 0, 0, 0
	fmt.Fscanf(reader, "%d%d%d\n", &A, &B, &K)
	cnt := 0
	for A < B {
		A *= K
		cnt++
	}
	fmt.Println(cnt)
}
