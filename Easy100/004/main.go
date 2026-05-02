package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	// 見つからなかった場合の初期値
	ans := -1
	for x := 1; x <= n; x++ {
		taxIncluded := (x * 108) / 100
		if taxIncluded == n {
			ans = x
			break
		}
	}

	if ans != -1 {
		fmt.Println(ans)
	} else {
		fmt.Println(":(")
	}
}
