package main

import (
	"fmt"
)

func main() {
	var n, m, c int
	fmt.Scan(&n, &m, &c)

	b := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&b[i])
	}

	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&a[i][j])
		}
	}

	count := 0
	for i := 0; i < n; i++ {
		sum := 0
		for j := 0; j < m; j++ {
			sum += a[i][j] * b[j]
		}
		if sum+c > 0 {
			count++
		}
	}
	fmt.Println(count)
}
