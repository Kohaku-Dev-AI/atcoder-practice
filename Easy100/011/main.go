package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	maxCount := 0
	maxNumber := 0
	for i := 1; i <= n; i++ {
		count := 0
		isActive := true
		for j := i; j > 0; j /= 2 {
			if j%2 == 0 {
				count++
			} else {
				isActive = false
			}

			if j == 1 || count > maxCount {
				maxCount = count
				maxNumber = i
			}

			if !isActive {
				break
			}
		}
	}
	fmt.Println(maxNumber)
}
