package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	X, A, D, N := int64(0), int64(0), int64(0), int64(0)
	fmt.Fscanf(reader, "%d%d%d%d\n", &X, &A, &D, &N)

	l, r := A, A+(N-1)*D
	if l > r {
		l, r = r, l
		D = -D
	}

	if X >= r {
		fmt.Println(X - r)
	} else if X <= l {
		fmt.Println(l - X)
	} else {
		r := (X - l + D) % D
		if r < D-r {
			fmt.Println(r)
		} else {
			fmt.Println(D - r)
		}
	}
}
