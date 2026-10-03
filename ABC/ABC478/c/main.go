package main

import (
	"bufio"
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
	n, k := nextInt(), nextInt()

	a := make([]int, n)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		val := nextInt()
		a[i] = val
		s[i] = val
	}

	slices.Sort(s)

	first := -1
	last := -1

	for i := 0; i < n; i++ {
		if a[i] != s[i] {
			if first == -1 {
				first = i
			}
			last = i
		}
	}

	if first == -1 {
		fmt.Println("Yes")
		return
	}

	diffLen := last - first + 1
	if diffLen <= k {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
