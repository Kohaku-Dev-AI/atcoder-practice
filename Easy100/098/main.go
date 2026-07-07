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
	a, b, c := nextInt(), nextInt(), nextInt()
	minValue := 0

	if a%2 == 0 || b%2 == 0 || c%2 == 0 {
		minValue = 0
	} else {
		max := max(a, b, c)
		if max == a {
			minValue = b * c
		} else if max == b {
			minValue = c * a
		} else {
			minValue = a * b
		}
	}
	fmt.Println(minValue)
}
