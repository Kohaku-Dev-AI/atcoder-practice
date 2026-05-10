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
	d := make([]int, n)

	for i := 0; i < n; i++ {
		d[i] = nextInt()
	}
	k := 0
	sort.Ints(d)
	if d[n/2-1] != d[n/2] {
		k = d[n/2] - d[n/2-1]
	} else {
		k = 0
	}

	fmt.Println(k)
}
