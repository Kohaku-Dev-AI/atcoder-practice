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
	count := 1
	minValue := nextInt()
	for i := 1; i < n; i++ {
		p := nextInt()
		if p < minValue {
			minValue = p
			count++
		}
	}
	fmt.Println(count)
}
