package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
	a := make([]int, n)
	b := make([]int, m)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}
	slices.Sort(a)
	for i := 0; i < m; i++ {
		b[i] = nextInt()
	}
	slices.Sort(b)
	count := 0
	i := 0
	j := 0
	for i < n && j < m {
		if a[i]*2 >= b[j] {
			count++
			i++
			j++
		} else {
			i++
		}
	}
	fmt.Println(count)
}
