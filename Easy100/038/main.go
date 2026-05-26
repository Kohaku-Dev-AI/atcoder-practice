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
	a := make([]int, 5)
	modA := make([]int, 5)
	quotient := make([]int, 5)
	for i := 0; i < 5; i++ {
		a[i] = nextInt()
		modA[i] = a[i] % 10
		if modA[i] != 0 {
			quotient[i] = a[i]/10 + 1
		} else {
			quotient[i] = a[i] / 10
		}
	}
	minA := 10
	totalTime := 0
	for i := 0; i < 5; i++ {
		if modA[i] < minA && modA[i] != 0 {
			minA = modA[i]
		}
		totalTime += quotient[i] * 10
	}
	fmt.Println(totalTime + minA - 10)
}
