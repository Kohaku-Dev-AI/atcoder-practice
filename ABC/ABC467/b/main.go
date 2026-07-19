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
	lossMoney := 0
	for i := 0; i < n; i++ {
		a := nextInt()
		b := nextInt()
		s := nextString()
		if s == "keep" {
			lossMoney += b - a
		}
	}
	fmt.Println(lossMoney)
}
