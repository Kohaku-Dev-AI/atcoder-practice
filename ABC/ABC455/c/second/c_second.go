package main

import (
	"fmt"
	"slices"
)

func main() {
	var n, k int

	if _, err := fmt.Scan(&n, &k); err != nil {
		return
	}

	counts := make(map[int]int)
	totalsum := 0

	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		totalsum += a
		counts[a]++
	}

	reductions := make([]int, 0, len(counts))
	for val, count := range counts {
		reductions = append(reductions, val*count)
	}

	slices.Sort(reductions)
	slices.Reverse(reductions)

	for i := 0; i < k && i < len(reductions); i++ {
		totalsum -= reductions[i]
	}
	fmt.Println(totalsum)
}
