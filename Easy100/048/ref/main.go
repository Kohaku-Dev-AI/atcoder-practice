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

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()

	prevHeight := nextInt()
	maxCount := 0
	currentCount := 0

	for i := 1; i < n; i++ {
		currentHeight := nextInt()
		if currentHeight <= prevHeight {
			currentCount++
			maxCount = max(maxCount, currentCount)
		} else {
			currentCount = 0
		}
		prevHeight = currentHeight
	}
	fmt.Println(maxCount)
}
