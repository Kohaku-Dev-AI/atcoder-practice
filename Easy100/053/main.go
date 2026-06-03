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
	n, m := nextInt(), nextInt()
	roadCount := make(map[int]int, n+1)
	for i := 0; i < m; i++ {
		a := nextInt()
		b := nextInt()
		roadCount[a]++
		roadCount[b]++
	}
	for i := 1; i < n+1; i++ {
		fmt.Printf("%d\n", roadCount[i])
	}
}
