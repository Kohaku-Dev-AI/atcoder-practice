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
	n, k, q := nextInt(), nextInt(), nextInt()
	a := make([]int, q)
	anserCount := make(map[int]int)
	for i := 0; i < q; i++ {
		a[i] = nextInt()
		anserCount[a[i]]++
	}
	for i := 1; i <= n; i++ {
		if value, ok := anserCount[i]; ok {
			if k-q+value > 0 {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		} else {
			if k-q > 0 {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		}
	}
}
