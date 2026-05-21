package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	s := nextString()

	minDif := 753

	for i := 0; i <= len(s)-3; i++ {
		subStr := s[i : i+3]
		x, _ := strconv.Atoi(subStr)
		currentDif := abs(x - 753)

		if currentDif < minDif {
			minDif = currentDif
		}
	}
	fmt.Println(minDif)
}
