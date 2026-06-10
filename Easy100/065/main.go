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
	const maxBuf = 1024 * 1024
	sc.Buffer(make([]byte, bufio.MaxScanTokenSize), maxBuf)
	sc.Split(bufio.ScanWords)
	s := nextString()
	firstCount := 0
	lastCount := 0
	for i, r := range s {
		if string(r) == "A" && firstCount == 0 {
			firstCount = i + 1
		} else if string(r) == "Z" {
			lastCount = i + 1
		}
	}
	fmt.Println(lastCount - firstCount + 1)
}
