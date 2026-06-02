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
	prev := nextInt()
	if n == 2 {
		println(prev * 2)
		return
	}
	maxValue := prev
	for i := 0; i < n-2; i++ {
		countValue := nextInt()
		if countValue > prev {
			maxValue += prev
		} else {
			maxValue += countValue
		}
		if i == n-3 {
			maxValue += countValue
		}
		prev = countValue
	}
	fmt.Println(maxValue)
}
