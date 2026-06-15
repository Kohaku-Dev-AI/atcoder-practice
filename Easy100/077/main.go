package main

import (
	"bufio"
	"cmp"
	"fmt"
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
	n := nextInt()
	participants := make([]int, 3*n)
	for i := 0; i < 3*n; i++ {
		participants[i] = nextInt()
	}
	slices.SortFunc(participants, func(a, b int) int {
		return cmp.Compare(b, a)
	})
	var result int64 = 0
	count := 0
	for i := 1; count < n; i += 2 {
		result += int64(participants[i])
		count++
	}

	fmt.Println(result)
}
