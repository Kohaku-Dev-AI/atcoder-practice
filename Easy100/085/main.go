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
	d, n := nextInt(), nextInt()

	if n == 100 {
		n = 101
	}

	result := 0
	if d == 0 {
		result = n
	} else if d == 1 {
		result = n * 100
	} else if d == 2 {
		result = n * 100 * 100
	}
	fmt.Println(result)
}
