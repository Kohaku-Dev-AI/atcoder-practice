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

var (
	n         int
	x, y      []float64
	used      []bool
	path      []int
	totalDist float64
	count     float64
)

func main() {
	sc.Split(bufio.ScanWords)
	n = nextInt()
	x = make([]float64, n)
	y = make([]float64, n)
	for i := 0; i < n; i++ {
		x[i] = float64(nextInt())
		y[i] = float64(nextInt())
	}

	used = make([]bool, n)
	path = make([]int, 0, n)

	dfs(0)
	ans := totalDist / count
	fmt.Printf("%.10f\n", ans)
}

func dfs(depth int) {
	if depth == n {
		count++
		dist := 0.0
		for i := 0; i < n-1; i++ {
			p1 := path[i]
			p2 := path[i+1]
			dx := x[p1] - x[p2]
			dy := y[p1] - y[p2]
			dist += math.Sqrt(dx*dx + dy*dy)
		}
		totalDist += dist
		return
	}

	for i := 0; i < n; i++ {
		if !used[i] {
			used[i] = true
			path = append(path, i)

			dfs(depth + 1)

			path = path[:len(path)-1]
			used[i] = false
		}
	}
}
