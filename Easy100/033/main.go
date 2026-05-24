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
	a, _ := nextInt(), nextInt()
	s := nextString()
	for i := 0; i < len(s); i++ {
		if i == a {
			if s[i] != '-' {
				fmt.Println("No")
				return
			}
		} else {
			if s[i] == '-' {
				fmt.Println("No")
				return
			}
		}
	}
	fmt.Println("Yes")
}
