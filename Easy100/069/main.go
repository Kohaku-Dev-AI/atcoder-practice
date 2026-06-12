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
	const maxBuf = 512 * 1024
	sc.Buffer(make([]byte, maxBuf), maxBuf)
	s := nextString()

	prevS := ""
	current := ""
	totalCount := 0

	for i := 0; i < len(s); i++ {

		current += string(s[i])

		if current != prevS {
			totalCount++
			prevS = current
			current = ""
		}
	}

	fmt.Println(totalCount)
}
