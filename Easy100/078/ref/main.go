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

	h := make([]int, n)
	for i := 0; i < n; i++ {
		h[i] = nextInt()
	}

	h[0]--

	for i := 1; i < n; i++ {
		if h[i]-1 >= h[i-1] {
			h[i]--
		} else if h[i] < h[i-1] {
			fmt.Println("No")
			return
		}
	}

	fmt.Println("Yes")
}
