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
	sum := 0
	for i := 0; i < n; i++ {
		a := nextInt()
		if i >= n/2 {
			sum += a
		}
	}
	fmt.Println(sum)
}
