package main

import (
	"fmt"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	totalDist := 0
	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		minDist := x
		if k-x < x {
			minDist = k - x
		}
		totalDist += minDist * 2
	}
	fmt.Println(totalDist)
}
