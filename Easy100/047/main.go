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
	if n == 1 {
		fmt.Println(1)
		return
	}
	l0 := 2
	l1 := 1
	for i := 2; i <= n; i++ {
		l0, l1 = l1, l0+l1
	}
	fmt.Println(l1)
}
