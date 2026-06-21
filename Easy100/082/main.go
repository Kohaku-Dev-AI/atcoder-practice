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
	q, h, s, d := nextInt(), nextInt(), nextInt(), nextInt()
	n := nextInt()
	min1L := min(4*q, 2*h, s)
	min2L := min(2*min1L, d)
	minValue := min2L*(n/2) + min1L*(n%2)
	fmt.Println(minValue)
}
