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
	c := nextString()
	if c == "B" {
		fmt.Println("Y")
	} else if c == "Y" {
		fmt.Println("R")
	} else if c == "R" {
		fmt.Println("B")
	}
}
