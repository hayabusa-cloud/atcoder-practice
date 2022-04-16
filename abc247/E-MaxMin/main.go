package main

/*
メモ：
条件を満たす(L,R)は必ずX, Yは中に入る。
今zが来ました。最後に出たX, Yを記録している。
スライス[s1..., X, s2..., Y, s3..., z]を考えます。
zを必ず選ぶ(L, R)は何通ですかというのを考えます。
最も右に位置する[Y, X]範囲外のaを考えます。
aはs1の所にある場合は、Xからzまで選んだ上、aからXまでは自由に選べます。合計pos[X]-pos[a]通です。
aはs2以降の所にある場合は0通です。
*/

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var N, MAX, MIN int
	_, _ = fmt.Fscanf(reader, "%d%d%d\n", &N, &MAX, &MIN)
	a := make([]int, N)
	for i := 0; i < N; i++ {
		_, _ = fmt.Fscanf(reader, "%d", &a[i])
	}
	posMAX, posMIN, posBad := -1, -1, -1
	res := 0
	for i := 0; i < N; i++ {
		if a[i] < MIN {
			posBad = i
			continue
		}
		if a[i] > MAX {
			posBad = i
			continue
		}
		if a[i] == MIN {
			posMIN = i
		}
		if a[i] == MAX {
			posMAX = i
		}
		left, right := posMIN, posMAX
		if left > right {
			left, right = right, left
		}
		if left <= posBad && posBad <= right {
			continue
		}
		if posBad >= right {
			continue
		}
		if posBad < left {
			res += left - posBad
		}
	}
	fmt.Println(res)
}
