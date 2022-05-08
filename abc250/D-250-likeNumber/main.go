package main

import (
	"bufio"
	"fmt"
	"os"
)

var pTable = []uint64{2}

func main() {
	pre()
	reader := bufio.NewReader(os.Stdin)

	N := uint64(0)
	fmt.Fscanf(reader, "%d\n", &N)

	cnt := 0
	for i := 0; i < len(pTable)-1; i++ {
		p := pTable[i]
		for j := i + 1; j < len(pTable); j++ {
			q := pTable[j]
			if q*q*q > N/p {
				break
			}
			cnt++
		}
	}
	fmt.Println(cnt)
}

func pre() {
	for p := 3; p < 793700; p += 2 {
		flag := true
		for i := 3; i*i <= p; i += 2 {
			if p%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			pTable = append(pTable, uint64(p))
		}
	}
}
