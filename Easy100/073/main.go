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
	n := nextString()

	sumN := 0
	for i := 0; i < len(n); i++ {
		sumN += int(n[i] - '0')
	}

	firstDigit := int(n[0] - '0')
	sumNine := (firstDigit - 1) + 9*(len(n)-1)

	ans := sumN
	if sumNine > ans {
		ans = sumNine
	}

	fmt.Println(ans)
}
