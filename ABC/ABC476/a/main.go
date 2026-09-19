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
	s := nextString()
	lastString := s[len(s)-1]
	T := ""
	if lastString == 'e' {
		T = s + "r"
	} else {
		T = s + "er"
	}
	fmt.Println(T)
}
