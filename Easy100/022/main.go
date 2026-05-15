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
	x := nextInt()

	a := make([]int, n)
	for i := range a {
		a[i] = nextInt()
	}
	sort.Ints(a)
	happyCount := 0
	for i := 0; i < n; i++ {
		if x < a[i] {
			break
		}
		x -= a[i]
		happyCount++
	}
	if happyCount == n && x > 0 {
		happyCount--
	}
	fmt.Println(happyCount)
}
