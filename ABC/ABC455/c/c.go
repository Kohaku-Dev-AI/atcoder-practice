package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	// ここに処理を書く

	lists := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&lists[i])
	}

	for a := 0; a < k; a++ {
		counts := make(map[int]int)
		var maxKey int

		for _, list := range lists {
			counts[list]++

			maxSumInt := -1
			for K, V := range counts {
				if P := K * V; P > maxSumInt {
					maxKey = K
				}
			}
		}
		for b := 0; b < n; b++ {
			if lists[b] == maxKey {
				lists[b] = 0
			}
		}
	}

	sum := 0
	for c := 0; c < n; c++ {
		sum += lists[c]
	}
	fmt.Println(sum)
}
