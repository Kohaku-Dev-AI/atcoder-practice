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
	n := nextInt()
	s := nextString()
	t := nextString()
	for i := 0; i < n; i++ {
		if t[i] == '*' {
			continue
		} else if s[i] == t[i] {
			continue
		} else {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
