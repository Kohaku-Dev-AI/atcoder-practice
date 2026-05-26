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

	totalTime := 0
	maxLoss := 0

	for i := 0; i < 5; i++ {
		time := nextInt()

		roundTime := (time + 9) / 10 * 10
		totalTime += roundTime

		loss := roundTime - time
		if loss > maxLoss {
			maxLoss = loss
		}
	}
	fmt.Println(totalTime - maxLoss)
}
