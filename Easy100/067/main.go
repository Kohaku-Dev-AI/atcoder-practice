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

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()

	counts := map[string]int{}
	for i := 0; i < n; i++ {
		s := nextString()
		counts[s]++
	}

	maxCount := 0
	for _, v := range counts {
		if v > maxCount {
			maxCount = v
		}
	}

	result := []string{}
	for k, v := range counts {
		if v == maxCount {
			result = append(result, k)
		}
	}

	sort.Strings(result)

	for _, s := range result {
		fmt.Println(s)
	}
}
