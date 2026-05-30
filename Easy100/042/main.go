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
	n, _ := nextInt(), nextInt()
	likeFood := make(map[int]int)
	for i := 0; i < n; i++ {
		k := nextInt()
		for j := 0; j < k; j++ {
			a := nextInt()
			likeFood[a]++
		}
	}
	count := 0
	for _, value := range likeFood {
		if value == n {
			count++
		}
	}
	fmt.Println(count)
}
