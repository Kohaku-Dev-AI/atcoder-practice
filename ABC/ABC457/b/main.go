package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	matrix := make([][]int, n)

	for i := range matrix {
		var l int
		fmt.Scan(&l)
		matrix[i] = make([]int, l)

		for j := range matrix[i] {
			fmt.Scan(&matrix[i][j])
		}
	}

	var x, y int
	fmt.Scan(&x, &y)

	fmt.Println(matrix[x-1][y-1])
}
