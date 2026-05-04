package main

import (
	"fmt"
)

func main() {
	a := make([][]int, 3)
	for i := 0; i < 3; i++ {
		a[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			fmt.Scan(&a[i][j])
		}
	}
	var n int
	fmt.Scan(&n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&b[i])
	}

	marked := make([][]bool, 3)
	for i := range marked {
		marked[i] = make([]bool, 3)
	}

	for _, value := range b {
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if a[i][j] == value {
					marked[i][j] = true
				}
			}
		}
	}
	ans := isBinngo(marked)
	if ans == true {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func isBinngo(m [][]bool) bool {
	for i := 0; i < 3; i++ {
		if m[i][0] && m[i][1] && m[i][2] {
			return true
		}
		if m[0][i] && m[1][i] && m[2][i] {
			return true
		}
	}

	if m[0][0] && m[1][1] && m[2][2] {
		return true
	}
	if m[2][0] && m[1][1] && m[0][2] {
		return true
	}

	return false
}
