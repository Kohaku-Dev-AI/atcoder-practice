package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	ans := 0
	sockets := 1

	for sockets < b {
		sockets += a - 1
		ans++
	}
	fmt.Println(ans)
}
