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

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)
	h, _ := nextInt(), nextInt()
	for i := 0; i < h; i++ {
		c := nextString()
		fmt.Println(c)
		fmt.Println(c)
	}
}
