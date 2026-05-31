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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	sc.Split(bufio.ScanWords)
	a, b, _, k := nextInt(), nextInt(), nextInt(), nextInt()
	if k == 0 {
		fmt.Println(a - b)
	}
	if k%2 == 0 {
		if abs(a-b) > 1000000000000000000 {
			fmt.Println("Unfair")
		} else {
			fmt.Println(a - b)
		}
	} else {
		if abs(a-b) > 1000000000000000000 {
			fmt.Println("Unfair")
		} else {
			fmt.Println(b - a)
		}
	}
}
