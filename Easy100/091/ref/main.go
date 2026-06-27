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

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()

	x := make([]float64, n)
	y := make([]float64, n)

	for i := 0; i < n; i++ {
		x[i] = float64(nextInt())
		y[i] = float64(nextInt())
	}

	totalDistance := 0.0

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dx := x[i] - x[j]
			dy := y[i] - y[j]
			dist := math.Sqrt(dx*dx + dy*dy)
			totalDistance += dist
		}
	}

	ans := totalDistance * 2.0 / float64(n)

	fmt.Printf("%.10f\n", ans)
}
