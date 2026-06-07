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
	defaultTime := make([]int, n)
	for i := 0; i < n; i++ {
		defaultTime[i] = nextInt()
	}
	m := nextInt()
	for i := 0; i < m; i++ {
		p := nextInt()
		x := nextInt()
		temp := make([]int, len(defaultTime))
		copy(temp, defaultTime)
		if temp[p-1] != x {
			temp[p-1] = x
		}
		totalValue := 0
		for j := 0; j < n; j++ {
			totalValue += temp[j]
		}
		fmt.Println(totalValue)
	}
}
