package main

/*
メモ：
線形探索していい
*/

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func f(a int64, b int64) int64 {
	return a*a*a + a*a*b + a*b*b + b*b*b
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	N := int64(0)
	fmt.Fscanf(reader, "%d\n", &N)

	minX := int64(math.MaxInt64)
	for a, b := int64(1e6), int64(0); a >= b; a-- {
		for {
			x := f(a, b)
			if x >= N {
				if x < minX {
					minX = x
				}
				break
			}
			b++
		}
	}

	fmt.Println(minX)
}
