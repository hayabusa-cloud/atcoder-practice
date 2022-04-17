package main

/*
メモ：スキャン
*/

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var S string
	fmt.Fscanf(reader, "%s\n", &S)
	flag := make([]bool, 10)
	for i := 0; i < 9; i++ {
		flag[S[i]-'0'] = true
	}
	for i := 0; i < 10; i++ {
		if !flag[i] {
			fmt.Printf("%c\n", '0'+i)
			return
		}
	}
}
