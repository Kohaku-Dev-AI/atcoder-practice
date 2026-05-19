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
	setLen := a + b
	blueBall := (n / setLen) * a
	rem := n % setLen
	if rem >= a {
		blueBall += a
	} else {
		blueBall += rem
	}
	fmt.Println(blueBall)
}
