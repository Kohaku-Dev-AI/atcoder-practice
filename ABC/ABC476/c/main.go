package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}

	top3 := make([]int, 3)
	copy(top3, a[:3])
	sort.Slice(top3, func(i, j int) bool {
		return top3[i] > top3[j]
	})

	if n == 3 {
		fmt.Println(top3[2])
		return
	} else {
		fmt.Println(top3[2])
	}

	for i := 3; i < n; i++ {
		if a[i] > top3[2] {
			top3[2] = a[i]
			sort.Slice(top3, func(i, j int) bool {
				return top3[i] > top3[j]
			})
			fmt.Println(top3[2])
		} else {
			fmt.Println(top3[2])
		}
	}
}
