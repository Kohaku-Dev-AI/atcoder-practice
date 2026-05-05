package main

import (
	"fmt"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	x := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x[i])
	}

	dist := 0
	for i := 0; i < n; i++ {
		if x[i] >= k-x[i] {
			dist += (k - x[i]) * 2
		} else {
			dist += x[i] * 2
		}
	}
	fmt.Println(dist)
}
