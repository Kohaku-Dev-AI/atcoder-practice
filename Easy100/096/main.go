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

	totalCount := 1
	oddNumberCount := 1

	for i := 0; i < n; i++ {
		a := nextInt()

		totalCount *= 3

		if a%2 == 0 {
			oddNumberCount *= 2
		} else {
			oddNumberCount *= 1
		}
	}

	fmt.Println(totalCount - oddNumberCount)
}
