package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	N, A, B := 0, 0, 0
	fmt.Fscanf(reader, "%d%d%d\n", &N, &A, &B)
	flag := false
	for i := 1; i <= A; i++ {
		flag = !flag
		for in := 1; in <= N; in++ {
			flag = !flag
			for j := 1; j <= B; j++ {
				for jn := 1; jn <= N; jn++ {
					if flag {
						writer.WriteByte('.')
					} else {
						writer.WriteByte('#')
					}
				}
			}
			writer.WriteByte('\n')
		}
		writer.WriteByte('\n')
	}
}
