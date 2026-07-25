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

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)
	m, d := nextInt(), nextInt()
	s := nextString()
	guard := make([]bool, m)
	for i, r := range s {
		if r == 'G' {
			for j := 0; j <= d; j++ {
				if i+j <= m-1 {
					guard[i+j] = true
				}
				if i-j >= 0 {
					guard[i-j] = true
				}
			}
		}
	}
	count := 0
	for i := 0; i < m; i++ {
		if guard[i] == false {
			count++
		}
	}

	fmt.Println(count)
}
