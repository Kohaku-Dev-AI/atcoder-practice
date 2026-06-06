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
	a, b, k := nextInt(), nextInt(), nextInt()
	if a+k-1 >= b-k+1 {
		for i := a; i <= b; i++ {
			fmt.Println(i)
		}
	} else {
		for i := a; i < a+k; i++ {
			fmt.Println(i)
		}
		for i := b - k + 1; i <= b; i++ {
			fmt.Println(i)
		}
	}
}
