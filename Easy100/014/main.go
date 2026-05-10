package main

import (
	"fmt"
)

func main() {
	var n, k int64 // 10^18を扱うためint64を使います
	fmt.Scan(&n, &k)

	if k == 0 {
		fmt.Println(n)
		return
	}

	// nをkで割った「余り」を求める
	// これにより、引き算を何兆回も繰り返すプロセスを一瞬でスキップできます
	n = n % k

	// 「kを引く前の値(n)」と「kを引いた後の値の絶対値(k-n)」の小さい方が答え
	if k-n < n {
		fmt.Println(k - n)
	} else {
		fmt.Println(n)
	}
}
