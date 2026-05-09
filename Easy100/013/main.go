package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if a == b && b == c && a%2 == 0 && b%2 == 0 && c%2 == 0 {
		fmt.Println(-1)
		return
	}
	ans := 0
	for a%2 == 0 && b%2 == 0 && c%2 == 0 {
		tempA := b/2 + c/2
		tempB := c/2 + a/2
		tempC := a/2 + b/2
		a = tempA
		b = tempB
		c = tempC
		ans++
	}
	fmt.Println(ans)
}
