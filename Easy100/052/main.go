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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	b := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		b[i] = nextInt()
	}
	ans := b[0]

	for i := 1; i < n-1; i++ {
		ans += min(b[i-1], b[i])
	}
	ans += b[n-2]
	fmt.Println(ans)
}
