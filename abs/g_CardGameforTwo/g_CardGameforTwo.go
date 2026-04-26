package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	var n int
	fmt.Scan(&n)

	lists := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&lists[i])
	}

	slices.SortFunc(lists, func(a, b int) int {
		return cmp.Compare(b, a)
	})

	var alice int
	var bob int

	for j := 0; j < n; j += 2 {
		alice += lists[j]
	}
	for k := 1; k < n; k += 2 {
		bob += lists[k]
	}
	fmt.Println(alice - bob)
}
