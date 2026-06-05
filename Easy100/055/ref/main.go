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

func isPerfectSquare(n int) bool {
	for q := 0; q*q <= n; q++ {
		if q*q == n {
			return true
		}
	}
	return false
}

func main() {
	sc.Split(bufio.ScanWords)
	n, d := nextInt(), nextInt()

	x := make([][]int, n)
	for i := 0; i < n; i++ {
		x[i] = make([]int, d)
		for j := 0; j < d; j++ {
			x[i][j] = nextInt()
		}
	}
	intCount := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			totalNumber := 0
			for k := 0; k < d; k++ {
				diff := x[i][k] - x[j][k]
				totalNumber += diff * diff
			}

			if isPerfectSquare(totalNumber) {
				intCount++
			}
		}
	}
	fmt.Println(intCount)
}
