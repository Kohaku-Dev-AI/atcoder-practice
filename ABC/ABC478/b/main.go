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
	n, v := nextInt(), nextInt()
	w := make([]int, n)
	for i := 0; i < n; i++ {
		w[i] = nextInt()
	}
	maxFun := 0

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				if i == j || j == k || k == i {
					continue
				}
				totalPrice := i + j + k + 3
				if totalPrice > v {
					continue
				}
				totalFun := w[i] + w[j] + w[k]
				if totalFun > maxFun {
					maxFun = totalFun
				}
			}
		}
	}
	fmt.Println(maxFun)
}
