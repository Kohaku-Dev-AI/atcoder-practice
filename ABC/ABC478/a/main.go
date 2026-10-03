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
	n, m := nextInt(), nextInt()
	base_count := m / n
	extra_count := m % n
	for i := 0; i < n; i++ {
		if i+1 <= extra_count {
			fmt.Println(base_count + 1)
		} else {
			fmt.Println(base_count)
		}
	}
}
