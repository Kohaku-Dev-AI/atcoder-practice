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
	a, b, c, x, y := nextInt(), nextInt(), nextInt(), nextInt(), nextInt()
	minMoney := 0
	smallestCount := min(x, y)
	if a+b >= 2*c {
		minMoney += c * 2 * smallestCount
		if x > smallestCount {
			if a > c*2 {
				minMoney += c * 2 * (x - smallestCount)
			} else {
				minMoney += a * (x - smallestCount)
			}
		} else if y > smallestCount {
			if b > c*2 {
				minMoney += c * 2 * (y - smallestCount)
			} else {
				minMoney += b * (y - smallestCount)
			}
		}
	} else {
		minMoney = a*x + b*y
	}
	fmt.Println(minMoney)
}
