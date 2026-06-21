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
	a := make([]int, n+1)
	b := make([]int, n)
	for i := 0; i < n+1; i++ {
		a[i] = nextInt()
	}
	for i := 0; i < n; i++ {
		b[i] = nextInt()
	}
	maxCount := 0
	for i := 0; i < n; i++ {
		if b[i] <= a[i] {
			maxCount += b[i]
		} else {
			maxCount += a[i]
			b[i] -= a[i]
			if b[i] <= a[i+1] {
				maxCount += b[i]
				a[i+1] -= b[i]
			} else {
				maxCount += a[i+1]
				a[i+1] = 0
			}
		}
	}
	fmt.Println(maxCount)
}
