package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt64() int64 {
	sc.Scan()
	i, _ := strconv.ParseInt(sc.Text(), 10, 64)
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	n, a, b := nextInt64(), nextInt64(), nextInt64()

	if (b-a)%2 == 0 {
		fmt.Println((b - a) / 2)
		return
	}

	leftRoute := a - 1 + 1 + (b-a-1)/2
	rightRoute := n - b + 1 + (b-a-1)/2

	fmt.Println(min(leftRoute, rightRoute))
}
