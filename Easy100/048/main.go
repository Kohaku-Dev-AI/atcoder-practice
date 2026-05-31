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
	h := make([]int, n)
	h[0] = nextInt()
	maxCount := 0
	correntCount := 0
	for i := 1; i < n; i++ {
		h[i] = nextInt()
		if h[i] <= h[i-1] {
			correntCount++
			if maxCount < correntCount {
				maxCount = correntCount
			}
		} else {
			correntCount = 0
		}
	}
	fmt.Println(maxCount)
}
