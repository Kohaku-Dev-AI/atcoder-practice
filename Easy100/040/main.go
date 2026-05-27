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
	a := nextInt()
	b := nextInt()
	if a*b <= 0 {
		fmt.Println("Zero")
	} else if a < 0 && b < 0 {
		if (b-a+1)%2 != 0 {
			fmt.Println("Negative")
		} else {
			fmt.Println("Positive")
		}
	} else {
		fmt.Println("Positive")
	}
}
