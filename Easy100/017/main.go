package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	// 1. 具材を float64 のスライスで読み込む
	v := make([]float64, n)
	for i := 0; i < n; i++ {
		v[i] = float64(nextInt())
	}
	sort.Float64s(v)
	ans := (v[0] + v[1]) / 2.0

	for i := 2; i < n; i++ {
		ans = (ans + v[i]) / 2.0
	}

	fmt.Println(ans)
}
