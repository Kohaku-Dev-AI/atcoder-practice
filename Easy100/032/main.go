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
	var s string
	fmt.Scan(&s)
	x := 0
	maxCount := 0
	for i := 0; i < n; i++ {
		if s[i] == 'I' {
			x++
			if x > maxCount {
				maxCount = x
			}
		} else if s[i] == 'D' {
			x--
		}
	}
	fmt.Println(maxCount)
}
