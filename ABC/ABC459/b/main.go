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
	s := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
	}
	c := make([]string, n)
	for i := 0; i < n; i++ {
		switch s[i][0] {
		case 'a', 'b', 'c':
			c[i] = "2"
		case 'd', 'e', 'f':
			c[i] = "3"
		case 'g', 'h', 'i':
			c[i] = "4"
		case 'j', 'k', 'l':
			c[i] = "5"
		case 'm', 'n', 'o':
			c[i] = "6"
		case 'p', 'q', 'r', 's':
			c[i] = "7"
		case 't', 'u', 'v':
			c[i] = "8"
		case 'w', 'x', 'y', 'z':
			c[i] = "9"
		}
	}
	var totalC string
	for i := 0; i < n; i++ {
		totalC += c[i]
	}
	fmt.Println(totalC)
}
