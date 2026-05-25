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
	const maxBufSize = 20 * 1024 * 1024
	sc.Buffer(make([]byte, bufio.MaxScanTokenSize), maxBufSize)
	sc.Split(bufio.ScanWords)
	s := nextString()
	alphabetMap := make(map[rune]bool)
	for _, char := range s {
		alphabetMap[char] = true
	}

	for i := 'a'; i <= 'z'; i++ {
		if !alphabetMap[i] {
			fmt.Println(string(i))
			return
		}
	}
	fmt.Println("None")
}
