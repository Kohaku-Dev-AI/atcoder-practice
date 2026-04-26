package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	counts := make(map[int]bool)
	for i := 0; i < n; i++ {
		var d int
		fmt.Scan(&d)
		counts[d] = true
	}

	fmt.Println(len(counts))
}
