package main

/*
メモ：
１カード＝１E、数字がVのグラフを考えます。
まずFindUnionやBFSでサイクルを求める。
サイクルCiを絞って考えます。答えをans[i]とします。
全体の答えはans[0]*ans[1]*...*ans[k]になります。kはサイクルの数です。
そしてans[i]を求めます。
len=Lのサイクルを考えます。通数をDP[L]とします。
場合ア：L番目のEを選ぶ。DP[L-1]通をカウントしていいです。
場合イ：L番目のEを選ばない。L-1番目のEは必ず選ぶのでDP[L-2]通をカウントします。
*/

import (
	"bufio"
	"fmt"
	"os"
)

func find(parent []int, i int) int {
	if parent[i] == i {
		return i
	}
	parent[i] = find(parent, parent[i])
	return parent[i]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	N := 0
	fmt.Fscanf(reader, "%d\n", &N)
	P := make([]int, N)
	Q := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscanf(reader, "%d", &P[i])
		P[i]--
	}
	fmt.Fscanln(reader)
	for i := 0; i < N; i++ {
		fmt.Fscanf(reader, "%d", &Q[i])
		Q[i]--
	}
	fmt.Fscanln(reader)

	parent := make([]int, N)
	for i := 0; i < N; i++ {
		parent[i] = i
	}

	for i := 0; i < N; i++ {
		parent[Q[i]] = find(parent, parent[P[i]])
	}

	m := make(map[int]int64)
	for i := 0; i < N; i++ {
		m[i] = 0
	}
	for i := 0; i < N; i++ {
		m[find(parent, i)]++
	}

	M := int64(998244353)
	dp := make([]int64, N+2)
	dp[0] = 1
	dp[1] = 1
	dp[2] = 3
	for i := 3; i <= N; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % M
	}

	result := int64(1)
	for _, v := range m {
		result = result * dp[v] % M
	}

	fmt.Println(result)
}
