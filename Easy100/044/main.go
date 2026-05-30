package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	if len(s) == 1 {
		fmt.Println(0)
		return
	}

	startWithZeroErrors := 0
	startWithOneErrors := 0

	for i := 0; i < len(s); i++ {
		char := s[i]

		if i%2 == 0 {
			if char != '0' {
				startWithZeroErrors++
			}
			if char != '1' {
				startWithOneErrors++
			}
		} else {
			if char != '1' {
				startWithZeroErrors++
			}
			if char != '0' {
				startWithOneErrors++
			}
		}
	}

	if startWithZeroErrors < startWithOneErrors {
		fmt.Println(startWithZeroErrors)
	} else {
		fmt.Println(startWithOneErrors)
	}
}
