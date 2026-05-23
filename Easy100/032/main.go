package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	const maxBufSize = 20 * 1024 * 1024 // 20MB
	sc.Buffer(make([]byte, bufio.MaxScanTokenSize), maxBufSize)

	sc.Split(bufio.ScanWords)
	n := nextInt()
	s := nextString()

	x := 0
	maxCount := 0

	for i := 0; i < n; i++ {
		if string(s[i]) == "I" {
			x++
			if x > maxCount {
				maxCount = x
			}
		} else if string(s[i]) == "D" {
			x--
		}
	}
	fmt.Println(maxCount)
}
