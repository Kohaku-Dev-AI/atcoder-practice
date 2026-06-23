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
	blueWord := make(map[string]int)
	for i := 0; i < n; i++ {
		s := nextString()
		blueWord[s]++
	}
	m := nextInt()
	for i := 0; i < m; i++ {
		t := nextString()
		blueWord[t]--
	}
	maxVal := 0
	for _, val := range blueWord {
		if val > maxVal {
			maxVal = val
		}
	}
	fmt.Println(maxVal)
}
