package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

var sc = bufio.NewScanner(os.Stdin)

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)
	s := nextString()
	slength := len(s)
	if s[0] != 'A' {
		fmt.Println("WA")
		return
	}

	cCount := 0
	for i := 2; i < slength-1; i++ {
		if s[i] == 'C' {
			cCount++
		}
	}
	if cCount != 1 {
		fmt.Println("WA")
		return
	}

	for _, r := range s {
		if r != 'A' && r != 'C' {
			if !unicode.IsLower(r) {
				fmt.Println("WA")
				return
			}
		}
	}
	fmt.Println("AC")
}
