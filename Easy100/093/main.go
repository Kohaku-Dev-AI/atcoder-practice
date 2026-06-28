package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt64() int64 {
	sc.Scan()
	i, _ := strconv.ParseInt(sc.Text(), 10, 64)
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	a := make([]int64, n) 
	var totalLength int64 = 0

	for i := 0; i < n; i++ {
		a[i] = nextInt64()
		totalLength += a[i]
	}

	var currentLength int64 = 0
	var minCost int64 = 1 << 62 

	for i := 0; i < n-1; i++ {
		currentLength += a[i]
		rightLength := totalLength - currentLength

		diff := currentLength - rightLength
		if diff < 0 {
			diff = -diff
		}

		if diff < minCost {
			minCost = diff
		}
	}

	fmt.Println(minCost)
}