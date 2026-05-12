package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)

	maxCount := 0
	currentCount := 0

	for _, char := range s {
		if char == 'A' || char == 'C' || char == 'G' || char == 'T' {
			currentCount++
			if currentCount > maxCount {
				maxCount = currentCount
			}
		} else {
			currentCount = 0
		}
	}

	fmt.Println(maxCount)
}
