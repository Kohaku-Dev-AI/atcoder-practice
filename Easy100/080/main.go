package main

import (
	"fmt"
	"math"
)

func main() {
	var n int64
	fmt.Scan(&n)

	start := int64(math.Sqrt(float64(n)))

	for i := start; i >= 1; i-- {
		if n%i == 0 {
			j := n / i
			ans := (i - 1) + (j - 1)
			fmt.Println(ans)
			return
		}
	}
}
