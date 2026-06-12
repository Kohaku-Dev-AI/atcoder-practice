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
	X := nextInt()

	if (X % 100) <= (X/100)*5 {
		fmt.Println("1")
	} else {
		fmt.Println("0")
	}
}
