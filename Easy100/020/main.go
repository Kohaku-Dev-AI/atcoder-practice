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
	s := nextInt()

	seen := make(map[int]bool)

	seen[s] = true
	currentValue := s

	for m := 2; ; m++ {
		if currentValue%2 == 0 {
			currentValue /= 2
		} else {
			currentValue = 3*currentValue + 1
		}

		if seen[currentValue] {
			fmt.Println(m)
			return
		}

		seen[currentValue] = true
	}
}
