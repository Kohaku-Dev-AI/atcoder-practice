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
	n, d, x := nextInt(), nextInt(), nextInt()
	a := make([]int, n)
	for i := range a {
		a[i] = nextInt()
	}
	eatCount := 0
	for i := 0; i < n; i++ {
		eatCount++
		for j := 1; j*a[i]+1 <= d; j++ {
			eatCount++
		}
	}
	x = x + eatCount
	fmt.Println(x)
}
