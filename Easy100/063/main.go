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

	target := 1

	for i := 0; i < n; i++ {
		a := nextInt()

		if a == target {
			target++
		}
	}

	if target == 1 {
		fmt.Println(-1)
	} else {
		fmt.Println(n - (target - 1))
	}
}
