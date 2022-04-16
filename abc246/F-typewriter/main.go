package main

/*
メモ：
DFS+包除原理
*/

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func power(b, n, mod int64) int64 {
	e := int64(1)
	for n > 0 {
		if n&1 == 1 {
			e = e * b % mod
		}
		b = b * b % mod
		n >>= 1
	}
	return e
}

func solve(a uint32, L int, mod int64) int64 {
	if a == 0 {
		return 0
	}
	return power(int64(bits.OnesCount32(a)), int64(L), mod)
}

func DFS(a []uint32, depth int, mask uint32, count int, L int) int64 {
	const mod = 998244353
	if depth >= len(a) {
		if count == 0 {
			return 0
		}
		if count&1 == 1 {
			return solve(mask, L, mod)
		}
		return mod - solve(mask, L, mod)
	}
	return (DFS(a, depth+1, mask, count, L) + DFS(a, depth+1, mask&a[depth], count+1, L)) % mod
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	N, L := 0, 0
	fmt.Fscanf(reader, "%d%d\n", &N, &L)
	a := make([]uint32, N)
	for i := 0; i < N; i++ {
		var s string
		fmt.Fscanf(reader, "%s\n", &s)
		for j := 0; j < len(s); j++ {
			a[i] |= 1 << (s[j] - 'a')
		}
	}

	res := DFS(a, 0, (1<<27)-1, 0, L)
	fmt.Println(res)
}
