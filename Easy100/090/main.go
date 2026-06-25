package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

const N = 1_000_000_000 + 7

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, 1024*1024)
	sc.Split(bufio.ScanWords)

	n := nextInt()
	result := 1
	for i := 1; i <= n; i++ {
		result = (result * i) % N
	}
	fmt.Println(result)
}
