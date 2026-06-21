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
	count := 0
	for i := 0; i < n; i++ {
		a := nextInt()
		for {
			if a%2 == 0 {
				count++
				a = a / 2
			} else {
				break
			}
		}
	}
	fmt.Println(count)
}
