package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	H, W, R, C := 0, 0, 0, 0
	fmt.Fscanf(reader, "%d%d\n", &H, &W)
	fmt.Fscanf(reader, "%d%d\n", &R, &C)
	ans := 4
	if R <= 1 {
		ans--
	}
	if R >= H {
		ans--
	}
	if C <= 1 {
		ans--
	}
	if C >= W {
		ans--
	}
	fmt.Println(ans)
}
