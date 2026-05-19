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
	h := nextInt()
	var totalCount int64 = 0
	var monsterNum int64 = 1

	for h > 0 {
		totalCount += monsterNum
		h = h / 2
		monsterNum *= 2
	}

	fmt.Println(totalCount)
}
