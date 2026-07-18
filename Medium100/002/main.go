package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	s, _ := reader.ReadString('\n')

	for len(s) > 0 && (s[len(s)-1] == '\n' || s[len(s)-1] == '\r') {
		s = s[:len(s)-1]
	}

	bCount := 0
	sumCount := 0
	for _, r := range s {
		if r == 'B' {
			bCount++
		} else if r == 'W' {
			sumCount += bCount
		}
	}
	fmt.Println(sumCount)
}
