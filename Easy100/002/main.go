package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	x := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x[i])
	}
	// 最小値を保持する変数
	minTotalPower := math.MaxInt64

	for p := 1; p <= 100; p++ {
		currentSum := 0
		for i := 0; i < n; i++ {
			diff := x[i] - p
			currentSum += diff * diff
		}
		if currentSum < minTotalPower {
			minTotalPower = currentSum
		}
	}

	fmt.Println(minTotalPower)
}
