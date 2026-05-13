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
	_ = nextInt()
	m := nextInt()
	currentPos := nextInt()

	costToZero := 0
	costToN := 0

	for i := 0; i < m; i++ {
		tollPos := nextInt()

		if tollPos < currentPos {
			costToZero ++
		} else {
			costToN ++
		}
	}

	if costToZero < costToN {
		fmt.Println(costToZero)
	} else {
		fmt.Println(costToN)
	}
}
