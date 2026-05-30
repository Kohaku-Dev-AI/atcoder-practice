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
	a, b := nextInt(), nextInt()
	count := 0
	for i := a; i <= b; i++ {
		s := strconv.Itoa(i)
		if s[0] == s[4] && s[1] == s[3] {
			count++
		}
	}
	fmt.Println(count)
}
