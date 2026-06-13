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

func ipow(base, exp int) int {
	result := 1
	for exp > 0 {
		result *= base
		exp--
	}
	return result
}

func main() {
	sc.Split(bufio.ScanWords)
	x := nextInt()
	if x <= 3 {
		fmt.Println(1)
		return
	}
	maxValue := 0
	for i := 2; i*i < x; i++ {
		for j := 2; ; j++ {
			result := ipow(i, j)
			if result > x {
				break
			}
			if result > maxValue {
				maxValue = result
			}
		}
	}
	fmt.Println(maxValue)
}
