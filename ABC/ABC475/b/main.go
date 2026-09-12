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
	n := nextInt()
	countCoin1 := 0
	countCoin10 := 0
	countCoin100 := 0
	for i := 0; i < n; i++ {
		a := nextInt()
		change := (1000 - a%1000) % 1000
		countCoin100 += change / 100
		change = change % 100
		countCoin10 += change / 10
		change = change % 10
		countCoin1 += change
	}
	fmt.Printf("%d %d %d\n", countCoin1, countCoin10, countCoin100)
}
