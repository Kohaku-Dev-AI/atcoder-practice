package main

import (
	"fmt"
)

func main() {
	var n, y int
	fmt.Scan(&n, &y)

	ans1, ans2, ans3 := -1, -1, -1

	for i := 0; i <= n; i++ {
		for j := 0; j <= n-i; j++ {
			k := n - i - j

			if 10000*i+5000*j+1000*k == y {
				ans1, ans2, ans3 = i, j, k
				break
			}
		}
		if ans1 != -1 {
			break
		}
	}
	fmt.Printf("%d %d %d\n", ans1, ans2, ans3)
}
