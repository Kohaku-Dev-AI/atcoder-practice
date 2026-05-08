package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if a == b && b == c {
		if a%2 == 0 {
			fmt.Println(-1)
		} else {
			fmt.Println(0)
		}
		return
	}

	ans := 0
	for (a|b|c)%2 == 0 {
		a, b, c = (b+c)/2, (a+c)/2, (a+b)/2
		ans++
	}
	fmt.Println(ans)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
