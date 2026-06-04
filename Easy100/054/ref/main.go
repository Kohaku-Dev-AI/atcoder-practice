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
	var n int
	fmt.Scan(&n)

	used := make(map[string]bool)

	prevWord := nextString()
	used[prevWord] = true

	for i := 1; i < n; i++ {
		currentWord := nextString()
		if prevWord[len(prevWord)-1] != currentWord[0] {
			fmt.Println("No")
			return
		}

		if used[currentWord] {
			fmt.Println("No")
			return
		}
		used[currentWord] = true
		prevWord = currentWord
	}
	fmt.Println("Yes")
}
