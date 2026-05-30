package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Scan(&input)

	charCounts := make(map[string]int)
	for _, r := range input {
		char := string(r)
		charCounts[char]++
	}
	isBeatiful := true
	for _, value := range charCounts {
		if value%2 != 0 {
			isBeatiful = false
			break
		}
	}
	if isBeatiful {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
