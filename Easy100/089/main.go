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
	x, y := nextInt(), nextInt()
	count := 0
	for i := x; i <= y; i = i * 2 {
		count++
	}
	fmt.Println(count)
}
