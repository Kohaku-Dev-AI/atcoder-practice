package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	sc.Split(bufio.ScanWords)
	n, t, a := nextInt(), nextInt(), nextInt()
	h := make([]int, n)
	for i := 0; i < n; i++ {
		h[i] = nextInt()
	}

	tFloat := float64(t)
	aFloat := float64(a)
	minDifferenceTemperature := math.Abs(tFloat - float64(h[0])*0.006 - aFloat)
	count := 1

	for i := 1; i < n; i++ {
		currentDifferenceTemperature := math.Abs(tFloat - float64(h[i])*0.006 - aFloat)
		if currentDifferenceTemperature < minDifferenceTemperature {
			minDifferenceTemperature = currentDifferenceTemperature
			count = i + 1
		}
	}
	fmt.Println(count)
}
