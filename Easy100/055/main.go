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

func isPerfectSquare(n int) bool {
	root := int(math.Round(math.Sqrt(float64(n))))
	return root*root == n
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
		for j := i + 1; j < n; i++ {
			totalNumber := 0
			for k := 0; k < d; k++ {
				totalNumber += (x[i][k] - x[j][k]) * (x[i][k] - x[j][k])
			}
			if isPerfectSquare(totalNumber) {
				intCount++
			}
		}
	}
	fmt.Println(intCount)
}
