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

	yOfX := make([]int, n+1)

	for i := 0; i < n; i++ {
		x := nextInt()
		y := nextInt()
		yOfX[x] = y
	}

	ans := 0
	minY := n + 1

	for x := 1; x <= n; x++ {
		currentY := yOfX[x]

		if currentY < minY {
			ans++
		}

		if currentY < minY {
			minY = currentY
		}
	}

	fmt.Println(ans)
}
