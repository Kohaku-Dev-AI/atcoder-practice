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

	// 答えを格納するためのスライスを用意しておく
	ans := make([]int, n)

	for i := 1; i <= n; i++ {
		a := nextInt()
		ans[a-1] = i
	}
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", ans[i])
	}
	fmt.Println()
}
