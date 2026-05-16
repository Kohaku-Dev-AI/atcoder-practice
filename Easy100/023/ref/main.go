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
	n, d, x := nextInt(), nextInt(), nextInt()

	totalEatemn := 0
	for i := 0; i < n; i++ {
		interval := nextInt()
		totalEatemn += (d-1)/interval + 1
	}

	fmt.Println(totalEatemn + x)
}
