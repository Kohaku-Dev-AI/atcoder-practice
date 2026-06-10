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
	nCount := 0
	wCount := 0
	sCount := 0
	eCount := 0
	for _, r := range s {
		if string(r) == "N" {
			nCount++
		} else if string(r) == "W" {
			wCount++
		} else if string(r) == "S" {
			sCount++
		} else if string(r) == "E" {
			eCount++
		}
	}
	if (nCount == 0 && sCount != 0) || (nCount != 0 && sCount == 0) || (wCount == 0 && eCount != 0) || (wCount != 0 && eCount == 0) {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
