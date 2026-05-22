package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)
	s := nextString()

	appeared := make(map[rune]bool)
	for _, char := range s {
		if appeared[char] {
			fmt.Println("no")
			return
		}
		appeared[char] = true
	}

	fmt.Println("yes")
}
