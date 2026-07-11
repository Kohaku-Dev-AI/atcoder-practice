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
	a := make([]int, n)
	pairCount := 0
	for i := 0; i < n; i++ {
		a[i] = nextInt()
		if a[a[i]-1] == i+1 {
			pairCount++
		}
	}
	fmt.Println(pairCount)
}
