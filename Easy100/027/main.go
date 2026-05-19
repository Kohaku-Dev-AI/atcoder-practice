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
	n, a, b := nextInt(), nextInt(), nextInt()
	redBall := 0
	redBall += (n/(a+b))*a
	n = n - (n/(a+b))*(a+b)
	if n >= a {
		redBall += a
	} else {
		redBall += n
	}
	fmt.Println(redBall)
}