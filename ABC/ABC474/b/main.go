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
	for i := 1; i <= n; i++ {
		p := nextInt()
		expectedGroup := (i - 1) / 10
		actualGroup := (p - 1) / 10
		if expectedGroup != actualGroup {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
