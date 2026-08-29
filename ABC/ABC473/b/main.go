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
	appears := make(map[int]bool)

	for i := 0; i < n; i++ {
		temp := nextInt()
		if appears[temp] {
			delete(appears, temp)
		} else {
			appears[temp] = true
		}
	}
	totalCount := 0
	for k := range appears {
		totalCount += k
	}
	fmt.Println(totalCount)
}
