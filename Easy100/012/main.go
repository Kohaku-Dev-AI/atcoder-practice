package main

import (
	"fmt"
)

func main() {
	var k, n int
	fmt.Scan(&k, &n)

	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	maxdist := 0
	for i := 1; i < n; i++ {
		dist := a[i] - a[i-1]
		if dist >= maxdist {
			maxdist = dist
		}
	}
	dist := a[0] + k - a[n-1]
	if dist >= maxdist {
		maxdist = dist
	}

	ans := k - maxdist
	fmt.Println(ans)
}
