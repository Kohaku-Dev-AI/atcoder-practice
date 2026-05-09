package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	nextInt := func() int {
		sc.Scan()
		i, _ := strconv.Atoi(sc.Text())
		return i
	}

	n := nextInt()
	k := nextInt()
	matrix := make([][]int, n)

	for i := 0; i < n; i++ {
		l := nextInt()
		matrix[i] = make([]int, l)
		for j := 0; j < l; j++ {
			matrix[i][j] = nextInt()
		}
	}

	c := make([]int, n)
	for i := 0; i < n; i++ {
		c[i] = nextInt()
	}

	// --- ここからがリファクタリングの肝 ---
	
	// 割り算（あまり）の計算を簡単にするため、Kを「0番目始まり」にずらす
	k-- 

	for i := 0; i < n; i++ {
		// 現在の数列 Ai が、配列 B に追加する「要素の合計数」
		totalAdded := c[i] * len(matrix[i])

		// K がこの追加分より小さいなら、答えは絶対にこの Ai の中にある！
		if k < totalAdded {
			// Ai は何回も繰り返されているが、中身のパターンは len(matrix[i]) 周期。
			// したがって、K を数列の長さで割った「あまり」の位置が答えになる。
			ans := matrix[i][k % len(matrix[i])]
			fmt.Println(ans)
			return // 答えが見つかったら即終了
		}

		// この数列の中に無いなら、追加された要素数を K から引いて、次の数列へ進む
		k -= totalAdded
	}
}