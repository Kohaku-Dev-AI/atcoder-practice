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
	m := nextInt()

	maxL := 1
	minR := n
	for i := 0; i < m; i++ {
		l := nextInt()
		r := nextInt()

		if l > maxL {
			maxL = l
		}
		if r < minR {
			minR = r
		}
	}

	ans := minR - maxL + 1
	if ans < 0 {
		ans = 0
	}

	fmt.Println(ans)
}
