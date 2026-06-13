package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
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
	n, k := nextInt(), nextInt()
	h := make([]int, n)
	for i := 0; i < n; i++ {
		h[i] = nextInt()
	}
	slices.Sort(h)
	minDiff := math.MaxInt

	for left := 0; left+k-1 < n; left++ {
		right := left + k - 1
		currentDiff := h[right] - h[left]
		minDiff = min(minDiff, currentDiff)
	}
	fmt.Println(minDiff)
}
