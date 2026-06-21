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
	upperRow := make([]int, n)
	downRow := make([]int, n)
	for i := 0; i < n; i++ {
		upperRow[i] = nextInt()
	}
	for i := 0; i < n; i++ {
		downRow[i] = nextInt()
	}
	maxCount := 0
	for i := 0; i < n; i++ {
		currentCount := 0
		for j := 0; j <= i; j++ {
			currentCount += upperRow[j]
		}
		for k := i; k < n; k++ {
			currentCount += downRow[k]
		}
		if maxCount < currentCount {
			maxCount = currentCount
		}
	}
	fmt.Println(maxCount)
}
