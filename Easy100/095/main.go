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

	maxVal := max(a, max(b, c))
	rem := (maxVal - a) + (maxVal - b) + (maxVal - c)

	if rem%2 == 0 {
		fmt.Println(rem / 2)
	} else {
		fmt.Println((rem + 3) / 2)
	}
}