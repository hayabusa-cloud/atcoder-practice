package main

/*
メモ：サーチ+DP
*/

import (
	"bufio"
	"fmt"
	"os"
)

const mod int64 = 998244353

var N, M, K int64 = 0, 0, 0

var dp [][]int64

func DFS(i int64, K int64) int64 {
	if i == 0 {
		return 1
	}
	sum := int64(0)
	for a := int64(1); a <= M; a++ {
		if a+i-1 > K {
			break
		}
		if dp[i-1][K-a] == 0 {
			dp[i-1][K-a] = DFS(i-1, K-a) % mod
		}
		sum = (sum + dp[i-1][K-a]) % mod

	}
	dp[i][K] = sum
	return sum
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscanf(reader, "%d%d%d\n", &N, &M, &K)

	dp = make([][]int64, N+1)
	for i := int64(0); i <= N; i++ {
		dp[i] = make([]int64, K+1)
	}

	fmt.Println(DFS(N, K))
}
