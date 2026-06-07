package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)
	s := nextString()

	if s[0] != 'A' {
		fmt.Println("WA")
		return
	}

	targetRange := s[2 : len(s)-1]

	firstC := strings.Index(targetRange, "C")
	lastC := strings.LastIndex(targetRange, "C")

	if firstC == -1 || firstC != lastC {
		fmt.Println("WA")
		return
	}

	posA := 0
	posC := 2 + firstC

	for i := 0; i < len(s); i++ {
		if i == posA || i == posC {
			continue
		}
		if s[i] < 'a' || s[i] > 'z' {
			fmt.Println("WA")
			return
		}
	}

	fmt.Println("AC")
}
